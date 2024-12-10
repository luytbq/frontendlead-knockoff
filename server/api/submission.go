package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/luytbq/frontendlead-knockoff/util"
)

var testEngineDir = "../test-engine"
var testCaseDir = testEngineDir + "/spec/test-cases"

func handleSubmission(w http.ResponseWriter, r *http.Request) {
	log.Println("handle submission")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var req RunRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil || req.Submission == "" || req.Slug == "" {
		http.Error(w, "Invalid submission", http.StatusBadRequest)
		return
	}

	question, err := util.GetQuestionBySlug(req.Slug)
	if err != nil {
		log.Printf("Error reading question: %s", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	if err = cleanPreviousTestScripts(); err != nil {
		log.Printf("Error cleaning previous test: %s", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	testCasePaths := make([]string, len(question.TestCases))

	for i := range question.TestCases {
		tc := question.TestCases[i]
		path, err := writeTestCase(tc.ID, tc.Test, req.Submission)
		if err != nil {
			log.Printf("Error writing test cases to disk: %s", err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
		testCasePaths[i] = path
	}

	res := &RunResponse{State: "success"}
	w.Header().Add("Content-Type", "application/json")

	for i := range question.TestCases {
		cmd := exec.Command("npx", "jasmine", testCasePaths[i])
		cmd.Dir = testEngineDir

		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()
		cmdOut := stdout.String()
		cmdErr := stderr.String()
		log.Printf("error: %v", err)
		log.Printf("stdout: %v", cmdOut)
		log.Printf("stderr: %v", cmdErr)
		if err != nil {
			res.State = "fail"
			res.FailAtID = i
			if cmdErr != "" {
				res.Message = cmdErr
			} else {
				res.Message = minifyJasmineOutput(cmdOut)
			}
			_ = json.NewEncoder(w).Encode(res)
			return
		}
	}

	res.Message = "All test cases passed"
	_ = json.NewEncoder(w).Encode(res)
}

/*
Edit the given testScript to the syntactically correct Jasmine code.
Appending the testScript with the submission code and write to a file on disk.
Returning the absolute file path and an error if any.
*/
func writeTestCase(id int, testScript string, submission string) (string, error) {
	testScript = editScriptSyntax(testScript)

	filePath := fmt.Sprintf("%s/test-case-%d.js", testCaseDir, id)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Printf("Failed to open file: %s", err.Error())
		return "", err
	}
	_, err = file.WriteString(appendTestScript(testScript, submission))
	if err != nil {
		log.Printf("Failed to write user code: %s", err.Error())
		return "", err
	}
	absPath, err := filepath.Abs(filePath)
	file.Close()
	return absPath, err
}

// Transform the given testScript to the syntactically correct Jasmine code.
func editScriptSyntax(testScript string) string {
	result := strings.ReplaceAll(testScript, ".to.deep.equal(", ".toEqual(")
	result = strings.ReplaceAll(result, ".to.equal(", ".toEqual(")
	result = strings.ReplaceAll(result, ".to.be.true", ".toEqual(true)")

	return result
}

func appendTestScript(testScript, submission string) string {
	result := ""
	if strings.Contains(testScript, "fetchMock.") {
		// by the test case syntax I deduced that the original app uses fetch-mock version 11
		log.Println("test case using fetchMock detected")
		result += "import fetchMock from 'fetch-mock';\n"
	}

	result += testScript + "\n" + submission

	return result
}

// Remove unnessessary text in jasmine stdout
func minifyJasmineOutput(cmdOut string) string {
	re := regexp.MustCompile(`(?s)Message:\s*(.*?)\s*Stack:`)
	match := re.FindStringSubmatch(cmdOut)
	if len(match) > 1 {
		return match[1]
	} else {
		return cmdOut
	}
}

// Clean temp test script files
func cleanPreviousTestScripts() error {
	files, err := filepath.Glob(testCaseDir + "/*.js")
	if err != nil {
		return err
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			return err
		}
	}

	return nil
}
