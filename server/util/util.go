package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/luytbq/frontendlead-knockoff/types"
)

func GetQuestionBySlug(slug string) (*types.QuestionDetail, error) {
	filename := fmt.Sprintf("./static/questions/%s.json", slug)
	log.Printf("Reading question from file: %s", filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("error opening file: %s", err.Error())
		return nil, err
	}
	defer file.Close()

	questionDetail := &types.QuestionDetail{}

	err = json.NewDecoder(file).Decode(questionDetail)
	if err != nil {
		log.Printf("error decoding file: %s", err.Error())
		return nil, err
	}

	return questionDetail, nil
}
