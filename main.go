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

type TokenRequest struct {
	Token string `json:"token"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Tidak bisa membaca file .env, pastikan file tersebut ada di root folder.")
	}
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token")
}

func verifyGoogleToken(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

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

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	if clientID == "" {
		http.Error(w, "Internal Server Error: Client ID belum diset", http.StatusInternalServerError)
		return
	}

	payload, err := idtoken.Validate(context.Background(), req.Token, clientID)
	if err != nil {
		http.Error(w, "Token Google tidak valid atau sudah kadaluarsa", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	responseMessage := fmt.Sprintf(`{"status": "success", "message": "Login berhasil diverifikasi!", "user_id": "%s"}`, payload.Subject)
	w.Write([]byte(responseMessage))
}

func main() {
	http.HandleFunc("/api/verify", verifyGoogleToken)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}

	fmt.Println("Server Golang berjalan di port", port)
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
