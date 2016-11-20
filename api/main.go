package main
 
import (
    "gopkg.in/gin-gonic/gin.v1"
    "models"
    "mappings"
)

var Router Router

func main() {
    Router := gin.Default()
    // TODO eventually use a real DB
    models.MockData()
    CreateUrlMappings()
    // Listen and server on 0.0.0.0:8080
    Router.Run(":8080")
}
