package controller

import (
	errors "awesomeProject/error"
	"awesomeProject/models"
	"encoding/json"
	"fmt"
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
	_ = json.NewDecoder(r.Body).Decode(&new_products)
	product := models.Product{}
	result_products, err := product.AddProducts(server.DB, new_products)

	fmt.Println("the new products is ", new_products)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ErrorMsg{Message: err.Error()})
		return

	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode("Products Added")
	json.NewEncoder(w).Encode(result_products)

}

//func (server *Server) BuyProduct(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r)
//
//	id, _ := strconv.Atoi(params["id"])
//	qty, _ := strconv.Atoi(params["qty"])
//
//	var product models.Product
//	var product2 models.Inventory
//	//if result := h.DB.First(&product, id); result.Error != nil {
//	//	fmt.Println(result.Error)
//	//}
//
//	if product.Quantity >= qty {
//		product.Quantity -= qty
//		product2.Orders += qty
//		//h.DB.Save(&product)
//		json.NewEncoder(w).Encode("Success ,Inventory Updated")
//		json.NewEncoder(w).Encode(product)
//		return
//	} else {
//		json.NewEncoder(w).Encode("Required Quantity is not available")
//		return
//	}
//
////}
//func  TopProducts(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	fmt.Println("top product is called")
//	//var products []Product
//	h.DB.Limit(5).Order("orders desc").Group("id").Where("updated_at - ? > ? ", time.Now(), time.Hour*-1).Find(&products)
//
//	json.NewEncoder(w).Encode(products)
//
//}