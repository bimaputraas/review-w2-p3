package server

import (
	pb "preview-w2/pb"
	"preview-w2/product_service/repository"
)

type ProductServer struct {
	pb.UnimplementedProductServiceServer
	Repository repository.Repository
}

func NewProductsServer(r repository.Repository) ProductServer {
	return ProductServer{Repository: r}
}