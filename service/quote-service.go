package service

import (
	"quotes-api/entity"
	"strconv"
)

type QuoteService interface {
	Save(quote entity.Quote) entity.Quote
	SaveMany(quotes []entity.Quote) []entity.Quote
	FindAll() []entity.Quote
	FindByID(id string) (entity.Quote, bool)
	Update(id string, quote entity.Quote) (entity.Quote, bool)
	Delete(id string) bool
}

type quoteService struct {
	quotes []entity.Quote
}

func New() QuoteService {
	return &quoteService{}
}

func (service *quoteService) Save(quote entity.Quote) entity.Quote {
	quote.ID = strconv.Itoa(len(service.quotes) + 1)
	service.quotes = append(service.quotes, quote)
	return quote
}

func (service *quoteService) SaveMany(quotes []entity.Quote) []entity.Quote {
	savedQuotes := []entity.Quote{}

	for _, quote := range quotes {
		savedQuote := service.Save(quote)
		savedQuotes = append(savedQuotes, savedQuote)
	}

	return savedQuotes
}

func (service *quoteService) FindAll() []entity.Quote {
	return service.quotes
}

func (service *quoteService) FindByID(id string) (entity.Quote, bool) {
	for _, quote := range service.quotes {
		if quote.ID == id {
			return quote, true
		}
	}

	return entity.Quote{}, false
}

func (service *quoteService) Update(id string, updatedQuote entity.Quote) (entity.Quote, bool) {
	for index, quote := range service.quotes {
		if quote.ID == id {
			updatedQuote.ID = id
			service.quotes[index] = updatedQuote
			return updatedQuote, true
		}
	}

	return entity.Quote{}, false
}

func (service *quoteService) Delete(id string) bool {
	for index, quote := range service.quotes {
		if quote.ID == id {
			service.quotes = append(service.quotes[:index], service.quotes[index+1:]...)
			return true
		}
	}

	return false
}
