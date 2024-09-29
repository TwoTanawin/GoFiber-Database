package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5433         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection
	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	db = sdb

	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	err = createProduct(&Product{Name: "Go Product", Price: 222})
	if err != nil {
		log.Fatal(err)
	}

	print("Create Successful!\n")

	product, err := getProduct(1)
	fmt.Println("Get Successful!", product)

	product, err = updateProduct(3, &Product{Name: "Springboot", Price: 300})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Update Successful!", product)

	err = deleteProduct(6)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Delete Successful!")

	products, err := getProducts()

	err = deleteProduct(6)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(products)
}
