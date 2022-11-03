package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// see products.json for sample JSON data.

// Product struct
type Product struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Response struct
// Response because it is the full JSON response
// Field is "products" which is a slice of Product structs.
type Response struct {
	Products []Product `json:"products"`
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

/*
  By default, object keys which don't have a corresponding struct field (or annotation) are ignored
*/

type OrderItem struct {
	Id       int     `json:"id"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}

type Order struct {
	Id    int         `json:"id"`
	Items []OrderItem `json:"items"`
}

type OrderResponse struct {
	Orders []Order `json:"orders"`
}

func main() {
	// Reading JSON example
	str := `{"name": "product", "id": 1}`
	product := Product{}
	json.Unmarshal([]byte(str), &product)
	fmt.Println(product)

	// Writing JSON example
	aBoolean, _ := json.Marshal(true)
	aString, _ := json.Marshal("a string")
	person := Person{
		Id:   1,
		Name: "Jim",
	}
	aPerson, _ := json.Marshal(&person)
	fmt.Println(string(aBoolean)) // true
	fmt.Println(string(aString))  // a string
	fmt.Println(string(aPerson))  // { "id": 1, "name": "Jim" }

	// Assignment to read data
	file, _ := ioutil.ReadFile("response.json")
	data := OrderResponse{}
	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Println("Something went wrong!", err)
	} else {
		for _, order := range data.Orders {
			orderTotal := 0.0
			fmt.Println("#########")
			fmt.Println("# Order ID:", order.Id)
			for _, item := range order.Items {
				fmt.Println("# Item ID:", item.Id)
				fmt.Println("# - Quant:", item.Quantity)
				fmt.Println("# - Total:", item.Total)
				orderTotal += item.Total
			}
			fmt.Println("# Order Total: $", orderTotal)
		}
		fmt.Println("#########")
	}
}
