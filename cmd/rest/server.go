package main

import (
	"quotes-api/controller"
	"quotes-api/entity"
	"quotes-api/service"

	"github.com/gin-gonic/gin"
)

var (
	quoteService    service.QuoteService       = service.New()
	quoteController controller.QuoteController = controller.New(quoteService)
)

func main() {
	server := gin.Default()

	server.GET("/quotes", func(ctx *gin.Context) {
		ctx.JSON(200, quoteController.FindAll())
	})

	server.POST("/quotes/bulk", func(ctx *gin.Context) {
		var quotes []entity.Quote
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
		var quote entity.Quote
		ctx.BindJSON(&quote)

		ctx.JSON(201, quoteController.Save(quote))
	})

	server.PUT("/quotes/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		var quote entity.Quote
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

	server.Run(":8080")
}
