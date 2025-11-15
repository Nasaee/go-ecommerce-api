package orders

import (
	"context"

	repo "github.com/Nasaee/go-ecommerce/internal/adapters/postgresql/sqlc"
)

type orderItem struct {
	ProductID int64 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}

/*
example payload

	{
	    "customerId": 42,
	    "items": [
	        {"productId": 1, "quantity": 1}
	    ]
	}
*/
type createOrderParams struct {
	CustomerID int64       `json:"customerId"`
	Items      []orderItem `json:"items"`
}

type Service interface {
	PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error)
}
