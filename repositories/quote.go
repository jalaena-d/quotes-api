package repositories

import (
	"context"
	"log"
	"os"
	"quotes-api/models"
	"strings"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/compute/metadata"
	"github.com/google/uuid"
	"golang.org/x/oauth2/google"
)

const datastoreScope = "https://www.googleapis.com/auth/datastore"

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

	if !isFirestoreEnabled() {
		log.Println("USE_FIRESTORE is not enabled; using in-memory repository")
		return defaultRepository
	}

	ctx := context.Background()
	projectID, projectIDSource := resolveProjectID(ctx)
	if projectID == "" {
		log.Println("USE_FIRESTORE is enabled but no project ID was found (checked FIRESTORE_PROJECT_ID, GOOGLE_CLOUD_PROJECT, GCP_PROJECT, GCLOUD_PROJECT, metadata, ADC); falling back to in-memory repository")
		return defaultRepository
	}

	databaseID := os.Getenv("FIRESTORE_DATABASE")
	collection := os.Getenv("FIRESTORE_COLLECTION")

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
		log.Printf("failed to create Firestore client (project=%q database=%q collection=%q); falling back to in-memory repository: %v", projectID, databaseID, collection, err)
		return defaultRepository
	}

	log.Printf("using Firestore repository (project=%q source=%s database=%q collection=%q)", projectID, projectIDSource, databaseID, collection)

	return NewFirestoreQuoteRepository(client, collection)
}

func isFirestoreEnabled() bool {
	value := strings.ToLower(strings.TrimSpace(os.Getenv("USE_FIRESTORE")))
	return value == "true" || value == "1" || value == "yes"
}

func resolveProjectID(ctx context.Context) (string, string) {
	envVars := []string{"FIRESTORE_PROJECT_ID", "GOOGLE_CLOUD_PROJECT", "GCP_PROJECT", "GCLOUD_PROJECT"}
	for _, envVar := range envVars {
		value := strings.TrimSpace(os.Getenv(envVar))
		if value != "" {
			return value, "env:" + envVar
		}
	}

	if metadata.OnGCE() {
		projectID, err := metadata.ProjectID()
		if err == nil && projectID != "" {
			return projectID, "metadata"
		}
	}

	credentials, err := google.FindDefaultCredentials(ctx, datastoreScope)
	if err == nil && credentials.ProjectID != "" {
		return credentials.ProjectID, "adc"
	}

	return "", ""
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
