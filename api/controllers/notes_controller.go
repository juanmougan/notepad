package main
 
import (
    "gopkg.in/gin-gonic/gin.v1"
    "net/http"
)

// TODO move this to a Model file
type Note struct {
    Title   string `form:"title" json:"title" binding:"required"`
    Body    string `form:"body" json:"body" binding:"required"`
}

var notes []Note

func allNotesEndpoint(c *gin.Context) {
    c.JSON(http.StatusOK, notes)
}

func mockData() {
    notes = append(notes, Note{Title: "My first note", Body: "Brazil beated Argentina very badly :("})
    notes = append(notes, Note{Title: "Some more notes", Body: "I hope we can defeat Colombia"})
}

func main() {
    // TODO eventually use a real DB
    mockData()
    router := gin.Default()

    // v1 of the API
    v1 := router.Group("/v1")
    {
        v1.GET("/notes", allNotesEndpoint)
        //v1.GET("/notes/:id", findNoteByIdEndpoint)
        //v1.POST("/notes", newNoteEndpoint)
        //v1.PUT("/notes", editNoteEndpoint)
        //v1.DELETE("/notes/:id", deleteNoteEndpoint)
    }

    // Listen and server on 0.0.0.0:8080
    router.Run(":8080")
}
