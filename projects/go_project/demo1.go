package go_project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyByte, _ := ioutil.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyByte, &products)
	fmt.Println(products)
}

func AddProduct() {
	product := Product{5,"Keyboard", 1, 59.99}
	productJson, err := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(productJson))
	
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var productResp Product
	json.Unmarshal(bodyBytes, &productResp)
	
	fmt.Println(productResp, "Saved to db!")
}
