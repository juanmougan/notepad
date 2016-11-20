package models

type Note struct {
    Title   string `form:"title" json:"title" binding:"required"`
    Body    string `form:"body" json:"body" binding:"required"`
}

var Notes []Note

func MockData() {
    Notes = append(Notes, Note{Title: "My first note", Body: "Brazil beated Argentina very badly :("})
    Notes = append(Notes, Note{Title: "Some more notes", Body: "I hope we can defeat Colombia"})
}

/*
func GetAllNotes() Note {
	return Notes
}
*/
