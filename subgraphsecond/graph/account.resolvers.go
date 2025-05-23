package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.73

import (
	"context"
	"subgraphsecond/graph/data"
	"subgraphsecond/graph/model"
)

// GetAccount is the resolver for the getAccount field.
func (r *queryResolver) GetAccount(ctx context.Context, accountReferenceID string) (*model.Account, error) {
	for _, account := range data.Accounts() {
		if account.AccountReferenceID == accountReferenceID {
			return account, nil
		}
	}
	return nil, nil
}

// GetAccounts is the resolver for the getAccounts field.
func (r *queryResolver) GetAccounts(ctx context.Context, customerReferenceID string) ([]*model.Account, error) {
	return data.Accounts(), nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
