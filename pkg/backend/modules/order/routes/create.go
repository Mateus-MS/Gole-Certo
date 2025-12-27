package order_routes

import (
	order_model "alves.com/backend/modules/order/model"
	order_service "alves.com/backend/modules/order/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func OrderCreate(stockService order_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request order_model.OrderRequest

		// Parse JSON body
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Convert UserID to ObjectID
		userID, err := primitive.ObjectIDFromHex(request.UserID)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid UserID"})
			return
		}

		// Convert products:
		//      map[string]int -> map[primitive.ObjectID]int
		products := make(map[primitive.ObjectID]int)
		for prodID, amount := range request.Products {
			id, err := primitive.ObjectIDFromHex(prodID)
			if err != nil {
				c.JSON(400, gin.H{"error": "invalid product ID: " + prodID})
				return
			}
			products[id] = amount
		}

		stockService.Create(c.Request.Context(), order_model.OrderEntity{
			ID:       primitive.NewObjectID(),
			UserID:   userID,
			Products: products,
		})
	}
}
