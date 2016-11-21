package controllers
 
import (
    "gopkg.in/gin-gonic/gin.v1"
    "net/http"
    "github.com/juanmougan/notepad/api/models"
)

func AllNotesEndpoint(c *gin.Context) {
    // TODO use this method instead
    //c.JSON(http.StatusOK, models.GetAllNotes)
    c.JSON(http.StatusOK, models.Notes)
}
