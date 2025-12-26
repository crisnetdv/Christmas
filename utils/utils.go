package utils

import (
	"os"
	"encoding/json"
	"github.com/Cristiano67867/christmas/models"
)

const filename string = "./comments.json"

func GetAllComments() ([]models.Comment, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []models.Comment{}, err
	}
	defer file.Close()

	var comments []models.Comment

	json.NewDecoder(file).Decode(&comments)

	return comments, nil
}

func SetComment(comment models.Comment) error {
	comments, err := GetAllComments()
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	json.NewDecoder(file).Decode(&comments)

	comments = append(comments, comment)
	
	json.NewEncoder(file).Encode(comments)

	return nil
}

// func main() {

// 	err := SetComment(Comment{
// 		Country: "Angola",
// 		Comment: "Great!",
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// }
