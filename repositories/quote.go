package repositories

import (
	"context"
	"log"
	"os"
	"quotes-api/models"

	"cloud.google.com/go/firestore"
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
	defaultRepository := &quoteService{}

	if os.Getenv("USE_FIRESTORE") != "true" {
		return defaultRepository
	}

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		log.Println("USE_FIRESTORE=true but GOOGLE_CLOUD_PROJECT is empty; falling back to in-memory repository")
		return defaultRepository
	}

	ctx := context.Background()
	databaseID := os.Getenv("FIRESTORE_DATABASE")

	var (
		client *firestore.Client
		err    error
	)

	if databaseID != "" {
		client, err = firestore.NewClientWithDatabase(ctx, projectID, databaseID)
	} else {
		client, err = firestore.NewClient(ctx, projectID)
	}

	if err != nil {
		log.Printf("failed to create Firestore client; falling back to in-memory repository: %v", err)
		return defaultRepository
	}

	return NewFirestoreQuoteRepository(client, os.Getenv("FIRESTORE_COLLECTION"))
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
