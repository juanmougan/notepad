package models

type Note struct {
    Title   string `form:"title" json:"title" binding:"required"`
    Body    string `form:"body" json:"body" binding:"required"`
}
