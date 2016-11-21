package main
 
import (
    "github.com/juanmougan/notepad/api/models"
    "github.com/juanmougan/notepad/api/mappings"
)

func main() {
    // TODO eventually use a real DB
    models.MockData()
    mappings.CreateUrlMappings()
    // Listen and server on 0.0.0.0:8080
    mappings.Router.Run(":8080")
}
