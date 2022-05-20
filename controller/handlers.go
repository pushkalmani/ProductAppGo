package controller

import (
	errors "awesomeProject/error"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (server *Server) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	product := models.Product{}
	products, err := product.GetAllProducts(server.DB)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ErrorMsg{Message: err.Error()})
		return

	}
	json.NewEncoder(w).Encode(products)

}

func (server *Server) GetProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	product := models.Product{}
	result_product, err := product.GetProductById(server.DB, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ErrorMsg{Message: err.Error()})
		return

	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result_product)
}

func (server *Server) AddProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	var new_products []models.Product
	err := json.NewDecoder(r.Body).Decode(&new_products)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.ErrorMsg{Message: err.Error()})
		return

	}
	product := models.Product{}
	result_products, err := product.AddProducts(server.DB, new_products)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ErrorMsg{Message: err.Error()})
		return

	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode("Products Added")
	json.NewEncoder(w).Encode(result_products)

}

func (server *Server) BuyProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var orders models.Orders
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&orders)

	result_product, err := product.GetProductById(server.DB, orders.ProductID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ErrorMsg{Message: err.Error()})
		return

	}

	if result_product.Quantity >= orders.Order_qty {
		result_product.Quantity -= orders.Order_qty

		update_product, err := product.UpdateProduct(server.DB, result_product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errors.ErrorMsg{Message: err.Error()})
			return

		}
		_, err = orders.CreateOrder(server.DB, orders)

		json.NewEncoder(w).Encode("Success ,Inventory Updated")
		json.NewEncoder(w).Encode(update_product)
		return
	} else {
		json.NewEncoder(w).Encode("Required Quantity is not available")
		return
	}

}

func (server *Server) RecommendProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var orders models.Orders

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	recommended_orders, err := orders.RecommendOrders(server.DB, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ErrorMsg{Message: err.Error()})
		return

	}

	json.NewEncoder(w).Encode("The recommended Products are")
	//for i := 0; i < len(recommended_orders); i++ {
	//	var product models.Product
	//	recommended_products, err := product.GetProductById(server.DB, recommended_orders.ProductID)
	//	if err != nil {
	//		w.WriteHeader(http.StatusInternalServerError)
	//		json.NewEncoder(w).Encode(errors.ErrorMsg{Message: err.Error()})
	//		return
	//
	//	}
	//
	//
	//}
	json.NewEncoder(w).Encode(recommended_orders)

}
