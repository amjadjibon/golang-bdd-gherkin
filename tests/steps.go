package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"

	"github.com/amjadjibon/golang-bdd-gherkin/models"
	"github.com/cucumber/godog"
)

type apiFeature struct {
	resp *httptest.ResponseRecorder
}

func (a *apiFeature) resetResponse(*godog.Scenario) {
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) httpRequest(method, path string, payload []byte) error {
	req, err := http.NewRequest(method, path, bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	a.resp = httptest.NewRecorder()
	a.resp.Body = bytes.NewBuffer(body)
	a.resp.Code = resp.StatusCode

	return nil
}

func (a *apiFeature) iSendRequestToWithPayload(method, path string, payload *godog.DocString) error {
	var err error
	path = "http://localhost:8080" + path

	if method == "GET" || method == "DELETE" {
		err = a.httpRequest(method, path, nil)
	} else {
		err = a.httpRequest(method, path, []byte(payload.Content))
	}

	if err != nil {
		return err
	}

	return nil
}

func (a *apiFeature) theResponseCodeShouldBe(code int) error {
	if code != a.resp.Code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
	}
	return nil
}

func (a *apiFeature) theResponseShouldMatchJSON(body *godog.DocString) (err error) {
	var expected, actual models.Book

	// re-encode actual response too
	if err = json.Unmarshal(a.resp.Body.Bytes(), &actual); err != nil {
		return
	}

	// define the pattern to match the expected ID
	pattern := `\{id\}`

	// replace the expected ID with the actual ID
	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Replace the matched text with the ID
	// Replace {id} with "1"
	expectedBody := re.ReplaceAllString(body.Content, fmt.Sprintf("%d", actual.ID))

	// re-encode expected response
	if err = json.Unmarshal([]byte(expectedBody), &expected); err != nil {
		return
	}

	// the matching may be adapted per different requirements.
	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected JSON does not match actual, %v vs. %v", expected, actual)
	}

	return nil
}

func (a *apiFeature) theFollowingBooksExist(table *godog.Table) error {
	for _, row := range table.Rows {
		book := models.BookBase{
			Title:  row.Cells[0].Value,
			Author: row.Cells[1].Value,
		}

		payload, err := json.Marshal(book)
		if err != nil {
			return err
		}

		err = a.httpRequest("POST", "http://localhost:8080/v1/books", payload)
		if err != nil {
			return err
		}

		if a.resp.Code != http.StatusCreated {
			return fmt.Errorf("expected response code to be: %d, but actual is: %d", http.StatusCreated, a.resp.Code)
		}
	}
	return nil
}

func (a *apiFeature) theResponseShouldBeAJSONArrayWithTheBooks() error {
	var books []models.Book
	if err := json.Unmarshal(a.resp.Body.Bytes(), &books); err != nil {
		return err
	}

	if len(books) == 0 {
		return fmt.Errorf("expected books to be returned, but got none")
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	api := &apiFeature{}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		api.resetResponse(sc)
		return ctx, nil
	})

	ctx.Step(`^I send "([^"]*)" request to "([^"]*)" with payload:$`, api.iSendRequestToWithPayload)
	ctx.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	ctx.Step(`^the response payload should match json:$`, api.theResponseShouldMatchJSON)

	// read all books steps
	ctx.Step(`^the following books exist:$`, api.theFollowingBooksExist)
	ctx.Step(`^the response should be a JSON array with the books$`, api.theResponseShouldBeAJSONArrayWithTheBooks)
}
