package internal

type Project struct {
	Key string `json:"key"`
}

type Issuetype struct {
	ID string `json:"id"`
}

type Fields struct {
	Summary     string    `json:"summary"`
	Issuetype   Issuetype `json:"issuetype"`
	Description string    `json:"description"`
	Project     Project   `json:"project"`
	Text2       string    `json:"customfield_14404"`
}

type JiraPayload struct {
	Fields Fields `json:"fields"`
}
