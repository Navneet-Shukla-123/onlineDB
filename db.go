package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func ConnectToDB() {
	host := os.Getenv("host") // e.g., db.your-supabase-instance.supabase.co
	port := os.Getenv("port")
	database := os.Getenv("dbname")
	user := os.Getenv("user")         // Found in the Supabase project settings
	password := os.Getenv("password") // The database password you set in Supabase

	// Form the connection string

	// log.Println("Host is ",host)
	// log.Println("Port is ",port)
	// log.Println("User is ",user)
	// log.Println("Db is ",database)
	// log.Println("Password is ",password)
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)

	// Connect to the database
	var err error
	conn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	//defer conn.Close(context.Background())

	fmt.Println("Connected to Supabase PostgreSQL database successfully")
}

func InsertToDB(data reponse) error {

	// Insert data into the 'users' table
	insertQuery := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := conn.QueryRow(context.Background(), insertQuery, data.Name, data.Email, data.Password).Scan(&id)
	if err != nil {
		log.Println("Failed to insert user: %v\n", err)
		return err
	}

	fmt.Printf("Inserted user with ID: %d\n", id)
	return nil
}
