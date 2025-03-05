package rpc

import (
	"context"
	"log"

	"github.com/MakiJOJO/douyin-mall-echo/app/product/internal/logic"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	page := req.GetPage()
	pageSize := req.GetPageSize()
	categoryName := req.GetCategoryName()
	products, err := logic.ListProducts(int(page), int(pageSize), categoryName)
	if err != nil {
		log.Println("Error listing products:", err)
	}
	results := make([]*product.Product, 0, len(products))
	for _, retProduct := range products {
		results = append(results, &product.Product{
			Id:          uint32(retProduct.ID),
			Name:        retProduct.Name,
			Description: retProduct.Description,
			Picture:     retProduct.Picture,
			Price:       retProduct.Price,
		})
	}

	resp = &product.ListProductsResp{
		Products: results,
	}
	return resp, nil
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (*product.GetProductResp, error) {
	productID := req.GetId() // 从请求中获取产品ID
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}
	retProduct, err := logic.GetProductByID(productID) // 获取产品信息
	if err != nil {
		log.Println("Error fetching product:", err)
		return nil, err
	}

	resp := &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(retProduct.ID),
			Name:        retProduct.Name,
			Description: retProduct.Description,
			Picture:     retProduct.Picture,
			Price:       retProduct.Price,
		},
	}
	return resp, nil
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	query := req.GetQuery() // 从请求中获取搜索关键字
	if query == "" {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "query is required")
	}
	retProducts, err := logic.SearchProducts(query) // 搜索产品信息
	if err != nil {
		log.Println("Error searching product:", err)
		return nil, err
	}

	results := make([]*product.Product, 0, len(retProducts))
	for _, retProduct := range retProducts {
		results = append(results, &product.Product{
			Id:          uint32(retProduct.ID),
			Name:        retProduct.Name,
			Description: retProduct.Description,
			Picture:     retProduct.Picture,
			Price:       retProduct.Price,
		})
	}

	resp = &product.SearchProductsResp{
		Results: results,
	}
	return resp, nil
}
