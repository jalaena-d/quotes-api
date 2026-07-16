package controller

import (
	"quotes-api/entity"
	"quotes-api/service"
)

type QuoteController interface {
	FindAll() []entity.Quote
	FindByID(id string) (entity.Quote, bool)
	Save(quote entity.Quote) entity.Quote
	SaveMany(quotes []entity.Quote) []entity.Quote
	Update(id string, quote entity.Quote) (entity.Quote, bool)
	Delete(id string) bool
}

type controller struct {
	service service.QuoteService
}

func New(service service.QuoteService) QuoteController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Quote {
	return c.service.FindAll()
}

func (c *controller) FindByID(id string) (entity.Quote, bool) {
	return c.service.FindByID(id)
}

func (c *controller) Save(quote entity.Quote) entity.Quote {
	return c.service.Save(quote)
}

func (c *controller) SaveMany(quotes []entity.Quote) []entity.Quote {
	return c.service.SaveMany(quotes)
}

func (c *controller) Update(id string, quote entity.Quote) (entity.Quote, bool) {
	return c.service.Update(id, quote)
}

func (c *controller) Delete(id string) bool {
	return c.service.Delete(id)
}
