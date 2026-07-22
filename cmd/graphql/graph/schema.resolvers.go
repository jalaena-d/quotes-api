package graph

import (
	"context"
	"fmt"
	"quotes-api/cmd/graphql/graph/model"
	"quotes-api/models"
)

func (r *mutationResolver) CreateQuote(ctx context.Context, input model.NewQuote) (*model.Quote, error) {
	quote := r.QuoteController.Save(models.Quote{
		Text:   input.Text,
		Author: input.Author,
	})

	return toGraphQLQuote(quote), nil
}

func (r *mutationResolver) UpdateQuote(ctx context.Context, id string, input model.UpdateQuote) (*model.Quote, error) {
	quote, found := r.QuoteController.Update(id, models.Quote{
		Text:   input.Text,
		Author: input.Author,
	})

	if !found {
		return nil, fmt.Errorf("quote not found")
	}

	return toGraphQLQuote(quote), nil
}

func (r *mutationResolver) DeleteQuote(ctx context.Context, id string) (bool, error) {
	deleted := r.QuoteController.Delete(id)

	if !deleted {
		return false, fmt.Errorf("quote not found")
	}

	return true, nil
}

func (r *queryResolver) Quotes(ctx context.Context) ([]*model.Quote, error) {
	quotes := r.QuoteController.FindAll()

	result := []*model.Quote{}
	for _, quote := range quotes {
		result = append(result, toGraphQLQuote(quote))
	}

	return result, nil
}

func (r *queryResolver) Quote(ctx context.Context, id string) (*model.Quote, error) {
	quote, found := r.QuoteController.FindByID(id)

	if !found {
		return nil, nil
	}

	return toGraphQLQuote(quote), nil
}

func toGraphQLQuote(quote models.Quote) *model.Quote {
	return &model.Quote{
		ID:     quote.ID,
		Text:   quote.Text,
		Author: quote.Author,
	}
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
