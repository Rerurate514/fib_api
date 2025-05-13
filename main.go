package main

import (
	fibonacci_rest_api "example.com/mod/lib/features/fibonacci_rest_api/rest"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/fib", fibonacci_rest_api.GetNumberAt)

	r.Run()
}
