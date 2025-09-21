package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func IsAdmin(admCollection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithError(http.StatusUnauthorized, errors.New("authorization required"))
			return
		}

		filter := bson.M{"name": username, "password": password}
		count, err := admCollection.CountDocuments(c, filter)
		if err != nil || count == 0 {
			c.AbortWithError(http.StatusForbidden, errors.New("must be ADM"))
			return
		}

		c.Next()
	}
}
