package handler

import (
	"encoding/json"
	"go-delivery-system/internal/model"
	"net/http"
)

var orders []model.Order

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order model.Order

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	order.ID = len(orders) + 1
	order.Status = "pending"
	orders = append(orders, order)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func ListOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
