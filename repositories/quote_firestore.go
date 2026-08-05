package repositories

import (
	"context"
	"log"
	"quotes-api/models"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type firestoreQuoteRepository struct {
	client     *firestore.Client
	collection string
}

func NewFirestoreQuoteRepository(client *firestore.Client, collection string) QuoteService {
	if collection == "" {
		collection = "quotes"
	}

	return &firestoreQuoteRepository{
		client:     client,
		collection: collection,
	}
}

func (r *firestoreQuoteRepository) Save(quote models.Quote) models.Quote {
	quote.ID = uuid.NewString()
	ctx := context.Background()
	_, err := r.client.Collection(r.collection).Doc(quote.ID).Create(ctx, models.Quote{
		ID:     quote.ID,
		Text:   quote.Text,
		Author: quote.Author,
	})
	if err != nil {
		log.Printf("firestore Save failed: %v", err)
		return models.Quote{}
	}

	return quote
}

func (r *firestoreQuoteRepository) SaveMany(quotes []models.Quote) []models.Quote {
	savedQuotes := make([]models.Quote, 0, len(quotes))
	for _, quote := range quotes {
		savedQuote := r.Save(quote)
		savedQuotes = append(savedQuotes, savedQuote)
	}

	return savedQuotes
}

func (r *firestoreQuoteRepository) FindAll() []models.Quote {
	ctx := context.Background()
	iter := r.client.Collection(r.collection).Documents(ctx)
	defer iter.Stop()

	docs, err := iter.GetAll()
	if err != nil {
		log.Printf("firestore FindAll failed: %v", err)
		return []models.Quote{}
	}

	quotes := make([]models.Quote, 0, len(docs))
	for _, doc := range docs {
		var quote models.Quote
		if err := doc.DataTo(&quote); err != nil {
			log.Printf("firestore FindAll decode failed: %v", err)
			return []models.Quote{}
		}
		quote.ID = doc.Ref.ID
		quotes = append(quotes, quote)
	}

	return quotes
}

func (r *firestoreQuoteRepository) FindByID(id string) (models.Quote, bool) {
	ctx := context.Background()
	doc, err := r.client.Collection(r.collection).Doc(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return models.Quote{}, false
		}
		log.Printf("firestore FindByID failed: %v", err)
		return models.Quote{}, false
	}

	var quote models.Quote
	if err := doc.DataTo(&quote); err != nil {
		log.Printf("firestore FindByID decode failed: %v", err)
		return models.Quote{}, false
	}

	quote.ID = doc.Ref.ID
	return quote, true
}

func (r *firestoreQuoteRepository) Update(id string, updatedQuote models.Quote) (models.Quote, bool) {
	ctx := context.Background()
	_, found := r.FindByID(id)
	if !found {
		return models.Quote{}, false
	}

	updatedQuote.ID = id
	_, err := r.client.Collection(r.collection).Doc(id).Set(ctx, updatedQuote)
	if err != nil {
		log.Printf("firestore Update failed: %v", err)
		return models.Quote{}, false
	}

	return updatedQuote, true
}

func (r *firestoreQuoteRepository) Delete(id string) bool {
	ctx := context.Background()
	_, found := r.FindByID(id)
	if !found {
		return false
	}

	_, err := r.client.Collection(r.collection).Doc(id).Delete(ctx)
	if err != nil {
		log.Printf("firestore Delete failed: %v", err)
		return false
	}

	return true
}
