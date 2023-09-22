package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"publisher/delivery"
	kafkaRepo "publisher/repository/kafka"
	postgresRepo "publisher/repository/postgres"
	"publisher/usecase"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/segmentio/kafka-go"
)

var kafkaConn *kafka.Writer
var postgresConn *sql.DB

func init() {
	fmt.Println("Init...")
	err := godotenv.Load(".env")
	kafkaHost := os.Getenv("KAFKA_HOST")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Kafka Host: " + kafkaHost)

	kafkaConn = &kafka.Writer{
		Addr:     kafka.TCP(kafkaHost),
		Balancer: &kafka.LeastBytes{},
		Topic:    os.Getenv("KAFKA_TOPIC"),
	}

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_DATABASE"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_SSLMODE"))
	postgresConn, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	port := os.Getenv("PUBLISHER_PORT")
	fmt.Println("Publisher Running on Port " + port + "...")
	kafkaRepo := kafkaRepo.NewProductRepository(kafkaConn)
	postgresRepo := postgresRepo.NewProductRepository(postgresConn)

	uc := usecase.NewProductUsecase(postgresRepo, kafkaRepo)
	h := delivery.NewHandler(uc)
	http.HandleFunc("/store", h.Store)
	http.HandleFunc("/update", h.Update)
	http.HandleFunc("/delete", h.Delete)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
