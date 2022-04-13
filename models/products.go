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
	producsSelect, err := db.Query("select * from products order by id asc")
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

	//Set products on database
	insertData, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, price, quantity)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.DatabaseConect()

	//Delete products from database
	delete, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.DatabaseConect()

	//Read products on database
	product, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	baseProduct := Product{}

	for product.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = product.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		baseProduct.Id = id
		baseProduct.Name = name
		baseProduct.Description = description
		baseProduct.Price = price
		baseProduct.Quantity = quantity
	}

	defer db.Close()
	return baseProduct
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.DatabaseConect()

	//Update products on database
	UpdateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	UpdateProduct.Exec(name, description, price, quantity, id)

	defer db.Close()
}
