package main
 
import (
    "gopkg.in/gin-gonic/gin.v1"
    "net/http"
)

Router := gin.Default()

func mockData() {
    notes = append(notes, Note{Title: "My first note", Body: "Brazil beated Argentina very badly :("})
    notes = append(notes, Note{Title: "Some more notes", Body: "I hope we can defeat Colombia"})
}

func main() {
    // TODO eventually use a real DB
    mockData()

    // Listen and server on 0.0.0.0:8080
    router.Run(":8080")
}
