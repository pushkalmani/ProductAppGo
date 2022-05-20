package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          int    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `gorm:"size:255;not null" json:"name"`
	Description string `gorm:"not null"json:"description"`
	Price       int    `gorm:"not null"json:"price"`
	Quantity    int    `gorm:"not null"json:"quantity"`
}

type Orders struct {
	UserId    int       `json:"user_id"`
	Order_qty int       `gorm:"not null"json:"order_qty"`
	ProductID int       `gorm:"not null"json:"product_id"`
	Product   Product   `gorm:"foreignkey:ProductID"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Product) GetAllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	return &products, nil
}
func (p *Product) GetProductById(db *gorm.DB, productID int) (*Product, error) {
	var err error
	product_by_id := Product{}
	err = db.Debug().Model(Product{}).Where("id = ?", productID).Take(&product_by_id).Error
	if err != nil {
		return nil, err
	}
	return &product_by_id, nil
}

func (p *Product) AddProducts(db *gorm.DB, listproduct []Product) (*[]Product, error) {
	var err error
	products := []Product{}

	for i := 0; i < len(listproduct); i++ {
		var add_product Product
		err = db.Where("name = ?", listproduct[i].Name).First(&add_product).Error
		if err != nil {
			err = db.Create(&listproduct[i]).Error
			if err != nil {
				return nil, err
			}
			var p1 Product
			db.Where("name = ?", listproduct[i].Name).First(&p1)
			products = append(products, p1)
		} else {
			add_product.Quantity = add_product.Quantity + listproduct[i].Quantity
			err = db.Save(&add_product).Error
			if err != nil {
				return nil, err
			}
			products = append(products, add_product)
		}
	}
	if err != nil {
		return nil, err
	}
	return &products, nil
}
func (p *Product) UpdateProduct(db *gorm.DB, product *Product) (*Product, error) {

	err := db.Save(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (o *Orders) CreateOrder(db *gorm.DB, order Orders) (*Orders, error) {
	err := db.Create(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}
func (o *Orders) RecommendOrders(db *gorm.DB, user_id int) (*[]Orders, error) {
	var order []Orders
	err := db.Where("user_id=?", user_id).Limit(5).Order("order_qty desc").Find(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}
