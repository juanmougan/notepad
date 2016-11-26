package mappings
 
import (
    "gopkg.in/gin-gonic/gin.v1"
    "github.com/juanmougan/notepad/api/controllers"
)

var Router *gin.Engine

func CreateUrlMappings() {
    Router = gin.Default()
    // v1 of the API
    v1 := Router.Group("/v1")
    {
        v1.GET("/notes", controllers.AllNotesEndpoint)
        //v1.GET("/notes/:id", findNoteByIdEndpoint)
        //v1.POST("/notes", newNoteEndpoint)
        //v1.PUT("/notes", editNoteEndpoint)
        //v1.DELETE("/notes/:id", deleteNoteEndpoint)
    }
}
