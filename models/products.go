package models

import "github.com/Johnman67112/web_go/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetProducts() []Product {
	db := db.DatabaseConect()

	//Database Query
	producsSelect, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	//Get rows from select and instance on p, following appends on products
	for producsSelect.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = producsSelect.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func SetProduct(name, description string, price float64, quantity int) {
	db := db.DatabaseConect()

	insertData, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, price, quantity)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.DatabaseConect()

	delete, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}
