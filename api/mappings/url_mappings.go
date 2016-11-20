package mappings
 
import (
    "../controllers"
    "../main"
)

func CreateUrlMappings() {
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
