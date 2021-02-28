package rpc

import (
	"context"

	"github.com/pawzio/arda/service/rivendell/genproto"
)

// GetProducts gets a list of products
func (s Server) GetProducts(context.Context, *genproto.Empty) (*genproto.GetProductsResponse, error) {
	return &genproto.GetProductsResponse{Products: make([]*genproto.Product, 0)}, nil
}
