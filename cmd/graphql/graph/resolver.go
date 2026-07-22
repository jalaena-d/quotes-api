package graph

import (
	"quotes-api/repositories"
	"quotes-api/services"
)

type Resolver struct {
	QuoteController services.QuoteController
}

func NewResolver() *Resolver {
	quoteService := repositories.NewQuoteRepository()
	quoteController := services.NewQuoteService(quoteService)

	return &Resolver{
		QuoteController: quoteController,
	}
}
