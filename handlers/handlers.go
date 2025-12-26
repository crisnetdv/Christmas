package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/Cristiano67867/christmas/models"
	"github.com/Cristiano67867/christmas/utils"
)

func HandlerCommentRouter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var body models.BodyData
	json.NewDecoder(r.Body).Decode(&body)

	utils.SetComment(models.Comment{
		Comment: body.Comment,
		Country: body.Country,
	})

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Comment stored",
	})
}


func HandlerGetAllCommentRouter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	comments, err := utils.GetAllComments()
	if err != nil {
		http.Error(w, "Error to get all comments", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(comments)
}