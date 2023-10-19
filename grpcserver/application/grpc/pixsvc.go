package grpc

import (
	"context"

	"github.com/lbaroni/imersao-desafio-01/grpcserver/application/grpc/pb"
	"github.com/lbaroni/imersao-desafio-01/grpcserver/application/usecase"
)

type ProductGrpcService struct {
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func (p ProductGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product, err := p.ProductUseCase.AddProduct(in.Name, in.Description, float64(in.Price))
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          int32(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price:       float32(product.Price),
		},
	}, nil
}

func (p ProductGrpcService) FindProducts(ctx context.Context, in *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	products, err := p.ProductUseCase.FindAllProducts()
	if err != nil {
		return &pb.FindProductsResponse{}, err
	}

	productSlice :=make([]*pb.Product,0,len(products))


	for i := 0; i < len(products); i++{ //aqui temos que transformar *model.Product em *pb.Product para passar no retorno
		productSlice = append(productSlice,	&pb.Product{
									Id: int32(products[i].ID),
									Name: products[i].Name,
									Description: products[i].Description,
									Price: float32(products[i].Price),
								 })
	}

	return &pb.FindProductsResponse {
				Products: productSlice,
			},nil
}
func NewProductGrpcService(usecase usecase.ProductUseCase) *ProductGrpcService{
	return &ProductGrpcService{
		ProductUseCase: usecase,
	}
}
