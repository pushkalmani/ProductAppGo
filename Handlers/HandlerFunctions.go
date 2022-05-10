package Handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type Product struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Quantity    int       `json:"quantity"`
	Orders      int       `json:"orders"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var products []Product

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{db}
}

func (h Handler) TopProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("top product is called")
	var products []Product
	h.DB.Limit(5).Order("orders desc").Group("id").Where("updated_at - ? > ? ", time.Now(), time.Hour*-1).Find(&products)

	json.NewEncoder(w).Encode(products)

}
func (h Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var products []Product
	if result := h.DB.Find(&products); result.Error != nil {
		fmt.Println(result.Error)
	}
	json.NewEncoder(w).Encode(products)
}

func (h Handler) GetProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	fmt.Println("here this is called ")
	var product Product
	if result := h.DB.First(&product, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	json.NewEncoder(w).Encode(product)
}

func (h Handler) AddProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	var products []Product
	_ = json.NewDecoder(r.Body).Decode(&products)

	if result := h.DB.Create(&products); result.Error != nil {
		fmt.Println(result.Error)
	}

	//for _, product := range products {
	//	h.DB.Model(&product).Update("CreatedAt", time.Now())
	//	h.DB.Update("updated_at", time.Now())
	//}

	json.NewEncoder(w).Encode("Products Added")
	json.NewEncoder(w).Encode(products)

}

func (h Handler) BuyProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])
	qty, _ := strconv.Atoi(params["qty"])

	var product Product

	if result := h.DB.First(&product, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	if product.Quantity >= qty {
		product.Quantity -= qty
		product.Orders += qty
		h.DB.Save(&product)
		json.NewEncoder(w).Encode("Success ,Inventory Updated")
		json.NewEncoder(w).Encode(product)
		return
	} else {
		json.NewEncoder(w).Encode("Required Quantity is not available")
		return
	}

}
