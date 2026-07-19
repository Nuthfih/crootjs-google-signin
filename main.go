package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/idtoken"
)

// Struktur data untuk menerima JSON dari Croot JS
type TokenRequest struct {
	Token string `json:"token"`
}

func init() {
	// Memuat konfigurasi dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Tidak bisa membaca file .env, pastikan file tersebut ada di root folder.")
	}
}

// Middleware untuk mengatasi CORS
func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "https://nuthfih.github.io")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
}

func verifyGoogleToken(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	// Handle preflight request dari browser
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req TokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Format JSON tidak valid", http.StatusBadRequest)
		return
	}

	// 1. Ambil Client ID dari Environment Variable
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	if clientID == "" {
		http.Error(w, "Internal Server Error: Client ID belum diset", http.StatusInternalServerError)
		return
	}

	// 2. Verifikasi token menggunakan package resmi Google
	// Context background dan token dari body JSON akan divalidasi terhadap Client ID kita
	payload, err := idtoken.Validate(context.Background(), req.Token, clientID)
	if err != nil {
		http.Error(w, "Token Google tidak valid atau sudah kadaluarsa", http.StatusUnauthorized)
		return
	}

	// Opsional: Kamu bisa mengekstrak data dari token yang sudah berhasil divalidasi
	// Contoh mengambil email: email := payload.Claims["email"]

	// Simulasi respons sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Memberikan response balikan dengan subject (ID Unik) pengguna
	responseMessage := fmt.Sprintf(`{"status": "success", "message": "Login berhasil diverifikasi!", "user_id": "%s"}`, payload.Subject)
	w.Write([]byte(responseMessage))
}

func main() {
	http.HandleFunc("/api/verify", verifyGoogleToken)

	// Ambil port dari server Domcloud
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback jika sedang di-run secara lokal
	}

	fmt.Println("Server Golang berjalan di port", port)
	
	// Tambahkan titik dua (":") di depan variabel port
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
