package main
 
import (
    "gopkg.in/gin-gonic/gin.v1"
    "net/http"
)

func createUrlMappings() {
    // v1 of the API
    v1 := router.Group("/v1")
    {
        v1.GET("/notes", allNotesEndpoint)
        //v1.GET("/notes/:id", findNoteByIdEndpoint)
        //v1.POST("/notes", newNoteEndpoint)
        //v1.PUT("/notes", editNoteEndpoint)
        //v1.DELETE("/notes/:id", deleteNoteEndpoint)
    }
}
