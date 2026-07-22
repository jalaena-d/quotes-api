package repositories

import (
	"quotes-api/models"

	"github.com/google/uuid"
)

type QuoteService interface {
	Save(quote models.Quote) models.Quote
	SaveMany(quotes []models.Quote) []models.Quote
	FindAll() []models.Quote
	FindByID(id string) (models.Quote, bool)
	Update(id string, quote models.Quote) (models.Quote, bool)
	Delete(id string) bool
}

type quoteService struct {
	quotes []models.Quote
}

func NewQuoteRepository() QuoteService {
	return &quoteService{}
}

func (service *quoteService) Save(quote models.Quote) models.Quote {
	quote.ID = uuid.NewString()
	service.quotes = append(service.quotes, quote)
	return quote
}

func (service *quoteService) SaveMany(quotes []models.Quote) []models.Quote {
	savedQuotes := []models.Quote{}

	for _, quote := range quotes {
		savedQuote := service.Save(quote)
		savedQuotes = append(savedQuotes, savedQuote)
	}

	return savedQuotes
}

func (service *quoteService) FindAll() []models.Quote {
	return service.quotes
}

func (service *quoteService) FindByID(id string) (models.Quote, bool) {
	for _, quote := range service.quotes {
		if quote.ID == id {
			return quote, true
		}
	}

	return models.Quote{}, false
}

func (service *quoteService) Update(id string, updatedQuote models.Quote) (models.Quote, bool) {
	for index, quote := range service.quotes {
		if quote.ID == id {
			updatedQuote.ID = id
			service.quotes[index] = updatedQuote
			return updatedQuote, true
		}
	}

	return models.Quote{}, false
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
