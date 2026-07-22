package services

import (
	"quotes-api/models"
	"quotes-api/repositories"
)

type QuoteController interface {
	FindAll() []models.Quote
	FindByID(id string) (models.Quote, bool)
	Save(quote models.Quote) models.Quote
	SaveMany(quotes []models.Quote) []models.Quote
	Update(id string, quote models.Quote) (models.Quote, bool)
	Delete(id string) bool
}

type controller struct {
	service repositories.QuoteService
}

func NewQuoteService(service repositories.QuoteService) QuoteController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []models.Quote {
	return c.service.FindAll()
}

func (c *controller) FindByID(id string) (models.Quote, bool) {
	return c.service.FindByID(id)
}

func (c *controller) Save(quote models.Quote) models.Quote {
	return c.service.Save(quote)
}

func (c *controller) SaveMany(quotes []models.Quote) []models.Quote {
	return c.service.SaveMany(quotes)
}

func (c *controller) Update(id string, quote models.Quote) (models.Quote, bool) {
	return c.service.Update(id, quote)
}

func (c *controller) Delete(id string) bool {
	return c.service.Delete(id)
}
