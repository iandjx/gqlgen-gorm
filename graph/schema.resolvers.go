package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/iandjx/go-order-graphql-api/graph/generated"
	"github.com/iandjx/go-order-graphql-api/graph/model"
	"github.com/iandjx/go-order-graphql-api/pkg/dbmodel"
)

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.OrderInput) (*model.Order, error) {
	var dbItems []dbmodel.Item

	for _, item := range input.Items {
		dbItems = append(dbItems, dbmodel.Item{ProductCode: item.ProductCode, ProductName: item.ProductName, Quantity: item.Quantity})
	}
	dbOrder := dbmodel.Order{
		CustomerName: input.CustomerName,
		OrderAmount:  input.OrderAmount,
		Items:        dbItems,
	}
	err := r.DB.Create(&dbOrder).Error
	if err != nil {
		return nil, err
	}

	var items []*model.Item
	for _, item := range input.Items {
		items = append(items, &model.Item{ProductCode: item.ProductCode, ProductName: item.ProductName, Quantity: item.Quantity})
	}
	order := model.Order{
		CustomerName: input.CustomerName,
		OrderAmount:  input.OrderAmount,
		Items:        items,
	}

	return &order, nil
}

func (r *mutationResolver) UpdateOrder(ctx context.Context, orderID int, input model.OrderInput) (*model.Order, error) {

	var items []*model.Item
	for _, item := range input.Items {
		items = append(items, &model.Item{ProductCode: item.ProductCode, ProductName: item.ProductName, Quantity: item.Quantity})
	}
	updatedOrder := model.Order{
		ID:           orderID,
		CustomerName: input.CustomerName,
		OrderAmount:  input.OrderAmount,
		Items:        items,
	}

	r.DB.Save(&updatedOrder)
	return &updatedOrder, nil
}

func (r *mutationResolver) DeleteOrder(ctx context.Context, orderID int) (bool, error) {
	// r.DB.Where("order_id = ?", orderID).Delete(&model.Order{})
	return true, nil
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	var orders []*model.Order
	var dborders []dbmodel.Order
	r.DB.Preload("Items").Find(&dborders)

	for _, order := range dborders {
		var items []*model.Item
		for _, item := range order.Items {
			items = append(items, &model.Item{ProductCode: item.ProductCode, ProductName: item.ProductName, Quantity: item.Quantity})
		}

		orders = append(orders, &model.Order{CustomerName: order.CustomerName, OrderAmount: order.OrderAmount, Items: items})
	}

	return orders, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
