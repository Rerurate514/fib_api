package fibonacci_rest_api

import (
	"net/http"
	"strconv"

	fibonacci_rest_api "example.com/mod/lib/features/fibonacci_rest_api/entities"
	"github.com/gin-gonic/gin"
)

func GetNumberAt(c *gin.Context) {
	indexStr := c.Param("index")
	index, err := strconv.Atoi(indexStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なインデックスです"})
		return
	}

	fib := fibonacci_rest_api.NewFibonacci()

	value, err := fib.GetNumberAt(index)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"value": value})
}
