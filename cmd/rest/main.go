package main

import (
	"log"
	"os"
	"quotes-api/models"
	"quotes-api/repositories"
	"quotes-api/services"

	"github.com/gin-gonic/gin"
)

var (
	quoteService    repositories.QuoteService = repositories.NewQuoteRepository()
	quoteController services.QuoteController  = services.NewQuoteService(quoteService)
)

func main() {
	server := gin.Default()

	server.GET("/quotes", func(ctx *gin.Context) {
		ctx.JSON(200, quoteController.FindAll())
	})

	server.POST("/quotes/bulk", func(ctx *gin.Context) {
		var quotes []models.Quote
		ctx.BindJSON(&quotes)

		ctx.JSON(201, quoteController.SaveMany(quotes))
	})

	server.GET("/quotes/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		quote, found := quoteController.FindByID(id)

		if !found {
			ctx.JSON(404, gin.H{
				"message": "Quote not found",
			})
			return
		}

		ctx.JSON(200, quote)
	})

	server.POST("/quotes", func(ctx *gin.Context) {
		var quote models.Quote
		ctx.BindJSON(&quote)

		ctx.JSON(201, quoteController.Save(quote))
	})

	server.PUT("/quotes/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		var quote models.Quote
		ctx.BindJSON(&quote)

		updatedQuote, found := quoteController.Update(id, quote)

		if !found {
			ctx.JSON(404, gin.H{
				"message": "Quote not found",
			})
			return
		}

		ctx.JSON(200, updatedQuote)
	})

	server.DELETE("/quotes/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		deleted := quoteController.Delete(id)

		if !deleted {
			ctx.JSON(404, gin.H{
				"message": "Quote not found",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "Quote deleted",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("starting REST server on port %s", port)
	server.Run(":" + port)
}
