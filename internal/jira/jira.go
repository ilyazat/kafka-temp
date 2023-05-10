package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kafka_template/internal"
	"log"
	"net/http"
)

type Config struct {
	ProjectKey string
	URL        string
	UserKey    string
	IssueType  string
}

type Jira struct {
	client http.Client
	config Config
}

func New(c http.Client, cfg Config) *Jira {
	return &Jira{
		client: c,
		config: cfg,
	}
}

func (j *Jira) CreateIssue(cr CreateIssueRequest) error {
	jiraPayload := &internal.JiraPayload{
		Fields: internal.Fields{
			Summary: cr.Name,
			Issuetype: internal.Issuetype{
				ID: j.config.IssueType,
			},
			Description: cr.Description,
			Project: internal.Project{
				Key: j.config.ProjectKey,
			},
		},
	}

	return j.sendRequest(jiraPayload)
}

func (j *Jira) sendRequest(payload interface{}) error {
	// Convert struct to JSON
	jiraJSON, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	log.Println(string(jiraJSON))

	// Create request
	req, err := http.NewRequest(
		http.MethodPost,
		j.config.URL,
		bytes.NewBuffer(jiraJSON),
	)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	j.useBasic(req)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := j.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	fmt.Println(string(body))

	// Read response
	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices {
		fmt.Println("Task created successfully")
		return nil
	}

	return fmt.Errorf("error creating task: %s", resp.Status)
}

func (j *Jira) useToken(req *http.Request) {
	// Set headers
	log.Println("Bearer " + j.config.UserKey)
	req.Header.Set("Authorization", "Bearer "+j.config.UserKey)
}

func (j *Jira) useBasic(req *http.Request) {
	// Set headers
	req.Header.Set("Authorization", "")
}
