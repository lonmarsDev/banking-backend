package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/lonmarsDev/banking-backend/internals/handler"
	"github.com/lonmarsDev/banking-backend/internals/services"
)

func main() {
	bankService := services.NewBankService()
	handler := handler.NewHandler(bankService)
	http.HandleFunc("POST /create-account", middleware(handler.HandleCreateAccount))
	http.HandleFunc("POST /deposit", middleware(handler.HandleDeposit))
	http.HandleFunc("GET /get-soa", middleware(handler.HandleGetAccount))
	http.HandleFunc("POST /withdraw", middleware(handler.HandleWithdraw))
	http.HandleFunc("GET /view-balance", middleware(handler.HandleViewBalance))

	slog.Info("Starting server on :8383")
	log.Fatal(http.ListenAndServe(":8383", nil))
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		slog.Info("Incoming request", "method", r.Method, "url", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
