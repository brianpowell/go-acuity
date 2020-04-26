/* acuity.go
// Acuity API interface
*/

package acuity

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/adityarama1210/go-url-builder/src/builder"
)

type (
	// Acuity main interface
	Acuity struct {
		config *Config
		// Bring in Routes
	}

	// Config structure
	Config struct {
		userID  string
		apiKey  string
		baseURL string
	}

	// Request for submission
	Request struct {
		Method    string
		URL       string
		Params    map[string]string
		Query     map[string]interface{}
		Body      interface{}
		BodyBytes io.Reader
	}
)

// Setup Acuity and return an Acuity Interface
func Setup(userID, apiKey string) *Acuity {
	conf := &Config{
		userID:  userID,
		apiKey:  apiKey,
		baseURL: "https://acuityscheduling.com/api/v1/",
	}

	return &Acuity{
		config: conf,
	}
}

func (a *Acuity) request(r Request, target interface{}) error {

	var err error
	// Do the Body for POST and PUT
	if r.Body != nil {
		body, _ := json.Marshal(r.Body)
		r.BodyBytes = bytes.NewBuffer(body)
	}

	// Build the URL and do the URL params + query
	url := a.urlParms(a.config.baseURL+r.URL, r.Params)
	url, _ = builder.CreateURLWithQuery(url, r.Query)

	// Send the Request
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, url, r.BodyBytes)
	req.Header.Add("Authorization", "Basic "+a.basicAuth(a.config.userID, a.config.apiKey))
	// Check the error
	if err != nil {
		return err
	}
	// Send the Request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Cast the output back to the interface
	if target != nil {
		if err := json.NewDecoder(resp.Body).Decode(&target); err != nil {
			return err
		}
	}

	return nil
}

func (a *Acuity) basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (a *Acuity) urlParms(url string, params map[string]string) string {
	// Do we have Params to replace
	if len(params) == 0 {
		return url
	}
	// replace them
	for i, p := range params {
		url = strings.Replace(url, ":"+i, p, 1)
	}
	return url
}
