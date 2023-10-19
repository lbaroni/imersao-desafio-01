package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/lbaroni/imersao-desafio-01/grpcserver/application/grpc/pb"
	"github.com/lbaroni/imersao-desafio-01/grpcserver/application/usecase"
	"github.com/lbaroni/imersao-desafio-01/grpcserver/infrastructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func StartGrpcServer(database *gorm.DB, port int){
	grpcServer := grpc.NewServer();
	reflection.Register(grpcServer) //para debug via grpc client

	productRepository := repository.ProductRepositoryDB{Db: database}
	productUseCase := usecase.ProductUseCase{ProductRepository: productRepository}
	productGrpcService := NewProductGrpcService(productUseCase)

	pb.RegisterProductServiceServer(grpcServer, productGrpcService)


	address := fmt.Sprintf("0.0.0.0:%d",port)
	listener, err := net.Listen("tcp", address)
	if err != nil{
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d",port)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}