package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/selectel/domains-go/pkg/testutils"
)

func TestNewDomainsClientV1(t *testing.T) {
	token := testutils.Token
	endpoint := "http://example.org"
	expected := &ServiceClient{
		Token:     token,
		Endpoint:  endpoint,
		UserAgent: userAgent,
	}

	actual := NewDomainsClientV1(token, endpoint)

	if expected.Token != actual.Token {
		t.Errorf("expected Endpoint %s, but got %s", expected.Endpoint, actual.Endpoint)
	}
	if expected.Endpoint != actual.Endpoint {
		t.Errorf("expected Token %s, but got %s", expected.Token, actual.Token)
	}
	if expected.UserAgent != actual.UserAgent {
		t.Errorf("expected UserAgent %s, but got %s", expected.UserAgent, actual.UserAgent)
	}
	if actual.HTTPClient == nil {
		t.Errorf("expected initialized HTTPClient but it's nil")
	}
}

func TestNewDomainsClientV1WithCustomHTTP(t *testing.T) {
	token := testutils.Token
	endpoint := "http://example.org"
	expected := &ServiceClient{
		Token:     token,
		Endpoint:  endpoint,
		UserAgent: userAgent,
	}

	customHTTPClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	actual := NewDomainsClientV1WithCustomHTTP(customHTTPClient, token, endpoint)

	if expected.Token != actual.Token {
		t.Errorf("expected Endpoint %s, but got %s", expected.Endpoint, actual.Endpoint)
	}
	if expected.Endpoint != actual.Endpoint {
		t.Errorf("expected Token %s, but got %s", expected.Token, actual.Token)
	}
	if expected.UserAgent != actual.UserAgent {
		t.Errorf("expected UserAgent %s, but got %s", expected.UserAgent, actual.UserAgent)
	}
	if actual.HTTPClient == nil {
		t.Errorf("expected initialized HTTPClient but it's nil")
	}
}

func TestNewDomainsClientV1WithDefaultEndpoint(t *testing.T) {
	token := testutils.Token
	expected := &ServiceClient{
		Token:     token,
		Endpoint:  defaultEndpoint,
		UserAgent: userAgent,
	}

	actual := NewDomainsClientV1WithDefaultEndpoint(token)

	if expected.Token != actual.Token {
		t.Errorf("expected Endpoint %s, but got %s", expected.Endpoint, actual.Endpoint)
	}
	if expected.Endpoint != actual.Endpoint {
		t.Errorf("expected Token %s, but got %s", expected.Token, actual.Token)
	}
	if expected.UserAgent != actual.UserAgent {
		t.Errorf("expected UserAgent %s, but got %s", expected.UserAgent, actual.UserAgent)
	}
	if actual.HTTPClient == nil {
		t.Errorf("expected initialized HTTPClient but it's nil")
	}
}

func TestDoGetRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, "response")

		if r.Method != http.MethodGet {
			t.Errorf("got %s method, expected GET", r.Method)
		}
	})

	token := testutils.Token
	userAgent := testutils.UserAgent
	endpoint := testEnv.Server.URL + "/"
	client := &ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   endpoint,
		Token:      token,
		UserAgent:  userAgent,
	}

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("got %d response status, expected 200", response.StatusCode)
	}
}

func TestDoPostRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, "response")

		if r.Method != http.MethodPost {
			t.Errorf("got %s method, expected POST", r.Method)
		}

		_, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("unable to read the request body: %v", err)
		}
	})

	token := testutils.Token
	userAgent := testutils.UserAgent
	endpoint := testEnv.Server.URL + "/"
	client := &ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   endpoint,
		Token:      token,
		UserAgent:  userAgent,
	}

	requestBody, err := json.Marshal(&struct {
		ID string `json:"id"`
	}{
		ID: "uuid",
	})
	if err != nil {
		t.Fatalf("can't marshal JSON: %v", err)
	}

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodPost, endpoint, bytes.NewReader(requestBody))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("got %d response status, expected 200", response.StatusCode)
	}
}

func TestDoErrNotFoundRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(w, `{"error": "domain_not_found"}`)

		if r.Method != http.MethodGet {
			t.Errorf("got %s method, expected GET", r.Method)
		}
	})

	token := testutils.Token
	userAgent := testutils.UserAgent
	endpoint := testEnv.Server.URL + "/"
	client := &ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   endpoint,
		Token:      token,
		UserAgent:  userAgent,
	}

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusNotFound {
		t.Fatalf("got %d response status, expected 404", response.StatusCode)
	}
	if response.ErrNotFound.Error != "domain_not_found" {
		t.Fatalf("got %s error message, expected 'domain_not_found'", response.ErrNotFound.Error)
	}
}

func TestDoErrGenericRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `{"error": "generic error"}`)

		if r.Method != http.MethodGet {
			t.Errorf("got %s method, want GET", r.Method)
		}
	})

	token := testutils.Token
	userAgent := testutils.UserAgent
	endpoint := testEnv.Server.URL + "/"
	client := &ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   endpoint,
		Token:      token,
		UserAgent:  userAgent,
	}

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("got %d response status, want 400", response.StatusCode)
	}

	if response.ErrGeneric.Error != "generic error" {
		t.Fatalf("got %s error message, want 'generic error'", response.ErrGeneric.Error)
	}
}

func TestDoErrNoContentRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		_, _ = fmt.Fprint(w, "") // write no content in the response body.

		if r.Method != http.MethodGet {
			t.Errorf("got %s method, want GET", r.Method)
		}
	})

	token := testutils.Token
	userAgent := testutils.UserAgent
	endpoint := testEnv.Server.URL + "/"
	client := &ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   endpoint,
		Token:      token,
		UserAgent:  userAgent,
	}

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusBadGateway {
		t.Fatalf("got %d response status, want 502", response.StatusCode)
	}

	if response.Err.Error() != "domains-go: got the 502 status code from the server" {
		t.Fatalf("got %s error message, want 'domains-go: got the 502 status code from the server'",
			response.Err.Error())
	}
}

func TestDoRequestInvalidResponseFromServer(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = fmt.Fprint(w, "<") // might be as a beginning of HTTP body

		if r.Method != http.MethodGet {
			t.Errorf("got %s method, want GET", r.Method)
		}
	})

	token := testutils.Token
	userAgent := testutils.UserAgent
	endpoint := testEnv.Server.URL + "/"
	client := &ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   endpoint,
		Token:      token,
		UserAgent:  userAgent,
	}

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("got %d response status, want 503", response.StatusCode)
	}

	if response.Err.Error() != "domains-go: got invalid response from the server, status code 503" {
		t.Fatalf("got %s error message, want 'domains-go: got invalid response from the server, status code 503'",
			response.Err.Error())
	}
}

func TestClientWithOSToken(t *testing.T) {
	token := testutils.Token
	client := NewDomainsClientV1(token, "http://example.org")
	OSClient := client.WithOSToken()

	if client == OSClient {
		t.Fatal(".WithOSToken() should create copy and point to different instance")
	}
	if client.isOpenstackToken != false {
		t.Fatal("initial client should have value of .isOpenstackToken = false")
	}
	if OSClient.isOpenstackToken != true {
		t.Fatal("OSClient should have value of .isOpenstackToken = true")
	}
}
