package main

import (
 "encoding/json"
 "fmt"
 "net/http"
 "log"
)
// Struktur data untuk menerima JSON dari Croot JS
type TokenRequest struct {
 Token string `json:"token"`
}
// Middleware untuk mengatasi CORS
func enableCORS(w *http.ResponseWriter) {
 (*w).Header().Set("Access-Control-Allow-Origin", "*") // Di tahap produksi, ganti "*" dengan URL GitHub Pages
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
 /* 
    DI SINI: Lakukan verifikasi token menggunakan package resmi Google:
    "google.golang.org/api/idtoken"
    Contoh: payload, err := idtoken.Validate(context.Background(), req.Token, "CLIENT_ID_KAMU")
 */
 // Simulasi respons sukses
 w.Header().Set("Content-Type", "application/json")
 w.WriteHeader(http.StatusOK)
 w.Write([]byte(`{"status": "success", "message": "Login Google berhasil diverifikasi oleh Golang!"}`))
}
func main() {
 http.HandleFunc("/api/verify", verifyGoogleToken)
 
 port := ":8080"
 fmt.Println("Server Golang berjalan di port", port)
 log.Fatal(http.ListenAndServe(port, nil))}