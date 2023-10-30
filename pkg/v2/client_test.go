package v2

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testAPIURL = "http://api.test.bonnie.com"
)

var (
	testHTTPClient = http.DefaultClient
	testClient     = &Client{testHTTPClient, make(http.Header), testAPIURL}
)

func TestProcessRequest_FailedRequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	ctx := context.Background()
	testMethod := http.MethodGet
	for _, responder := range []httpmock.Responder{
		httpmock.NewStringResponder(http.StatusBadRequest, mockErrorResponse("")),
		httpmock.NewStringResponder(http.StatusInternalServerError, mockErrorResponse("")),
		httpmock.NewBytesResponder(http.StatusInternalServerError, []byte{}),
		httpmock.NewStringResponder(http.StatusBadRequest, `{"unexpected_field": "value"}`),
	} {
		httpmock.RegisterResponder(testMethod, testAPIURL, responder)

		request, _ := http.NewRequestWithContext(ctx, testMethod, testAPIURL, nil)
		res, err := processRequest[Zone](testHTTPClient, request, nil)

		assert.Nil(t, res)
		assert.NotNil(t, err)
	}
}

func TestWithHeaders(t *testing.T) {
	// Check initial headers before method call
	assert.Equal(t, http.Header{}, testClient.defaultHeaders)

	newHeaders := make(http.Header)
	newHeaders.Add("field", "value")

	tmpClient, isOk := testClient.WithHeaders(newHeaders).(*Client)
	require.True(t, isOk, "WithHeaders must be convert to *Client")

	assert.Equal(t, http.Header{}, testClient.defaultHeaders)
	assert.Equal(t, newHeaders, tmpClient.defaultHeaders)
}

func TestParamsWithCommas(t *testing.T) {
	ctx := context.Background()

	type testCase struct {
		input  string
		length int
	}
	cases := []testCase{
		{"", 1},
		{"val", 1},
		{"val1,val2", 2},
		{"val1, val2", 2},
	}
	for _, test := range cases {
		params := &map[string]string{"field": test.input}
		r, _ := testClient.prepareRequest(ctx, http.MethodGet, "localhost", nil, params, nil)
		assert.Equal(t, test.length, len(r.URL.Query()["field"]))
	}
}

func mockErrorResponse(msg string) string {
	if msg == "" {
		msg = "random-test-error"
	}

	return fmt.Sprintf(`{"error": "%s"}`, msg)
}
