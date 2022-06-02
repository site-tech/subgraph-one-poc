package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/site-tech/subgraph-one-poc/graph/generated"
	"github.com/site-tech/subgraph-one-poc/graph/model"
)

func (r *entityResolver) FindManufacturerByID(ctx context.Context, id string) (*model.Manufacturer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *entityResolver) FindProductByManufacturerIDAndID(ctx context.Context, manufacturerID string, id string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *entityResolver) FindProductByUpc(ctx context.Context, upc string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
