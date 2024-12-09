package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/luytbq/frontendlead-knockoff/types"
)

func main() {
	questions, err := queryQuestionsList()
	if err != nil {
		log.Fatal(err)
		return
	}

	queryQuestions(questions)

	log.Println("Done.")
}

func queryQuestionsList() ([]types.Question, error) {
	graphQLEndpoint := "https://api.frontendlead.com/graphql"
	graphQLStr := `{"query":"\n  {\n    problems(first: 250) {\n      nodes {\n        title\n\t\tcontent\n        slug\n        date\n        featuredImage {\n          node {\n            sourceUrl\n          }\n        }\n        info {\n            company\n            category\n            type\n            }\n        common {\n            difficulty\n            videoLink {\n                url\n            }\n            comingSoon\n            subtitle\n            premiumQuestion\n\t\t\tstudyPlan\n        }\n      }\n    }\n  }\n  "} `

	log.Printf("Crawling data at %s", graphQLEndpoint)
	body := strings.NewReader(graphQLStr)

	res, err := http.Post(graphQLEndpoint, "application/json", body)
	if err != nil {
		return nil, fmt.Errorf("Error making POST request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	crawlResponse := &types.CrawlListQuestionResponse{}
	err = json.NewDecoder(res.Body).Decode(crawlResponse)
	if err != nil {
		return nil, fmt.Errorf("Error decoding response: %v", err)
	}

	questions := crawlResponse.Data.Problems.Nodes
	log.Printf("Crawled %d questions", len(questions))
	sortQuestions(questions)

	filename := "./static/crawled-questions.json"
	log.Printf("Create/overwrite %s", filename)

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("Open file error: %s", err.Error())
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(questions); err != nil {
		return nil, fmt.Errorf("Failed to encode JSON to file: %v", err)
	}

	log.Printf("%d questions found\n\n", len(questions))
	log.Printf("Question list successfully written to %s", filename)

	return questions, nil
}

func queryQuestions(questions []types.Question) {
	cSlug := make(chan string)
	wg := &sync.WaitGroup{}
	numWorkers := 8

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for slug := range cSlug {
				if err := querySingleQuestion(slug); err != nil {
					log.Printf("error: %s", err.Error())
				} else {
					log.Println("success")
				}
			}
		}()
	}

	for i := range questions {
		cSlug <- questions[i].Slug
	}
	close(cSlug)

	wg.Wait()
}

func querySingleQuestion(slug string) error {
	if slug == "" {
		return fmt.Errorf("empty slug")
	}

	questionDetailEndpoint := fmt.Sprintf("https://frontendlead.com/_next/data/y6tsfnZoSCXbXMavCyns1/coding-questions/%s.json", slug)

	log.Printf("Crawling data at %s", questionDetailEndpoint)

	res, err := http.Get(questionDetailEndpoint)
	if err != nil {
		return fmt.Errorf("error making GET request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	crawlResponse := types.CrawQuestionDetailResponse{}
	err = json.NewDecoder(res.Body).Decode(&crawlResponse)
	if err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}
	questionDetail := crawlResponse.PageProps.Problem
	filename := "./static/questions/" + slug + ".json"
	log.Printf("Create/overwrite %s", filename)

	// write to file
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("open file error: %s", err.Error())
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(questionDetail); err != nil {
		return fmt.Errorf("failed to encode JSON to file: %v", err)
	}

	log.Printf("Data successfully written to %s", filename)

	return nil
}

func sortQuestions(questions []types.Question) {
	sort.Slice(questions, func(i, j int) bool {
		prefix := regexp.MustCompile(`^[\d]*\. `)
		title1 := prefix.ReplaceAllString(questions[i].Title, "")
		title2 := prefix.ReplaceAllString(questions[j].Title, "")

		return title1 < title2
	})
}
