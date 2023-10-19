package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/lbaroni/imersao-desafio-01/application/grpc"
	"github.com/lbaroni/imersao-desafio-01/infrastructure/db"
)

var database *gorm.DB

func main(){
	database = db.ConnectDB(os.Getenv("env"))
	stats := database.DB().Stats()
	log.Printf("Db stats: %v", stats)
	grpc.StartGrpcServer(database, 50051)
	
}