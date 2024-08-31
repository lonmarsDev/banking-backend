package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/lonmarsDev/banking-backend/internals/services"
	"github.com/lonmarsDev/banking-backend/models"
)

type Handler struct {
	svc services.BankService
}

func NewHandler(service services.BankService) *Handler {
	return &Handler{svc: service}
}

func (s *Handler) HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var account models.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		slog.Error("failed to decode request body", "error", err)
		errorHandler(w, http.StatusBadRequest, err)
		return
	}

	// Create the account
	account = s.svc.CreateAccount(account.Name, account.Balance)

	// Return the account
	// impossible to fail, no need to capture error value
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func (s *Handler) HandleDeposit(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var depositRequest struct {
		ID     int     `json:"id"`
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&depositRequest); err != nil {
		errorHandler(w, http.StatusBadRequest, err)
		return
	}

	// Deposit the amount
	err := s.svc.Deposit(depositRequest.ID, depositRequest.Amount)
	if err != nil {
		slog.Error("failed to deposit amount", "error", err)
		errorHandler(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Handler) HandleWithdraw(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var withdrawRequest struct {
		ID     int     `json:"id"`
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&withdrawRequest); err != nil {
		slog.Error("failed to decode request body", "error", err)
		errorHandler(w, http.StatusBadRequest, err)
		return
	}

	// Withdraw the amount
	err := s.svc.Withdraw(withdrawRequest.ID, withdrawRequest.Amount)
	if err != nil {
		slog.Error("failed to withdraw amount", "error", err)
		errorHandler(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Handler) HandleGetAccount(w http.ResponseWriter, r *http.Request) {
	accountID := r.URL.Query().Get("id")
	if accountID == "" {
		errorHandler(w, http.StatusBadRequest, errors.New("missing account id"))
		return
	}
	// account id is not an integer
	id, _ := strconv.Atoi(accountID)
	// Get the account
	account, err := s.svc.GetAccount(id)
	if err != nil {
		slog.Error("failed to get account", "error", err)
		errorHandler(w, http.StatusBadRequest, err)
		return
	}

	// Return the account
	// impossible to fail, no need to capture error value
	json.NewEncoder(w).Encode(account)
}

func errorHandler(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.Error{
		Error: err.Error(),
	})
}

func (s *Handler) HandleViewBalance(w http.ResponseWriter, r *http.Request) {
	// get on query params
	accountID := r.URL.Query().Get("id")
	if accountID == "" {
		errorHandler(w, http.StatusBadRequest, errors.New("missing account id"))
		return
	}
	// account id is not an integer
	id, _ := strconv.Atoi(accountID)
	// Get the account
	account, err := s.svc.GetAccount(id)
	if err != nil {
		slog.Error("failed to get account", "error", err)
		errorHandler(w, http.StatusBadRequest, err)
		return
	}
	// map response
	type response struct {
		Name    string  `json:"name"`
		Balance float64 `json:"balance"`
	}
	resp := response{
		Balance: account.Balance,
		Name:    account.Name,
	}
	// Return the account
	// impossible to fail, no need to capture error value
	json.NewEncoder(w).Encode(resp)

}
