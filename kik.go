package kik

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	Text      = "text"      // default message type
	Suggested = "suggested" // default keyboard type

	apiEndpoint = "https://api.kik.com/v1"
	batchLimit  = 5 // Maximum messages that can be sent at once, enforeced by Kik.
	get         = "GET"
	post        = "POST"
)

// Kik API errors
type ApiError struct {
	ErrorText string `json:"error,omitempty"`
	Message   string `json:"message,omitempty"`
}

// Implement the `Error()` interface
func (a ApiError) Error() string {
	return a.Message
}

// Set of messages recieved via webhook
type Payload struct {
	Username string    // identifies the user to which these messages were sent
	Messages []Message `json:"messages,omitempty"`
}

// End-user callback for handling webhook events
type webhookCallback func(Payload, error)

// Kik API Client
type Client struct {
	Username string
	ApiKey   string
	Verbose  bool
	Callback webhookCallback
}

// Webhook handler. Parse incoming data and send it to the `Client.Callback`
// function.
func (c *Client) Webhook(w http.ResponseWriter, r *http.Request) {
	if c.Callback == nil {
		log.Println("Kik Webhook Error: no callback set")
		return
	}

	var payload Payload

	payload.Username = r.Header.Get("X-Kik-Username")
	if c.Verbose {
		log.Printf("Kik Webhook: incoming payload for user '%s'", payload.Username)
	}

	// Read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		if c.Verbose {
			log.Printf("Kik Webhook Error (Read Body): %s", err.Error())
		}
		c.Callback(payload, err)
		return
	}

	if c.Verbose && string(body) == "" {
		log.Println("Kik Webhook: received empty body")
	} else if c.Verbose {
		log.Printf("Kik Webhook: %s", body)
	}

	// Parse JSON
	if string(body) != "" {
		if err = json.Unmarshal(body, &payload); err != nil {
			if c.Verbose {
				log.Printf("Kik Webhook Error (JSON): %s", err.Error())
			}
			c.Callback(payload, err)
			return
		}
	}

	c.Callback(payload, err)
}

// Peform an HTTP request against the Kik API
func (c *Client) apiRequest(method string, path string, reqBody *[]byte) (int, []byte, error) {
	var (
		url      string = apiEndpoint + path
		req      *http.Request
		res      *http.Response
		resBody  []byte
		err      error
		apiError *ApiError
	)

	// Log requests
	if c.Verbose {
		var body string
		if reqBody != nil {
			body = string(*reqBody)
		}
		log.Printf("-> Kik API: %s %s %s", method, path, body)
	}

	// We have to be super careful about type of var we send here or
	// `http.NewRequest` till take it the wrong way and panic()
	if reqBody == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		b := bytes.NewBuffer(*reqBody)
		req, err = http.NewRequest(method, url, b)
	}

	if err != nil {
		return http.StatusInternalServerError, []byte(""), err
	}

	// Authenticate and set the right content type, so important
	req.SetBasicAuth(c.Username, c.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	// Fire away
	client := &http.Client{}
	res, err = client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return res.StatusCode, []byte(""), err
	}

	// Read body
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, []byte(""), err
	}

	// Log responses
	if c.Verbose {
		var body string
		if res.Body != nil {
			body = string(resBody)
		}
		log.Printf("<- Kik API: %s", body)
	}

	// Make sure API errors are treated as actual errors
	if res.StatusCode != http.StatusOK {
		if err := json.Unmarshal(resBody, &apiError); err != nil {
			if c.Verbose {
				log.Printf("Kik API Error (JSON): %s", err.Error())
			}
			return res.StatusCode, []byte(""), err
		}
		if c.Verbose {
			log.Printf("Kik API Error: %s - %s", apiError.ErrorText, apiError.Message)
		}
		return res.StatusCode, resBody, apiError
	}

	return res.StatusCode, resBody, nil
}
