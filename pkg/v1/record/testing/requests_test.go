package testing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/domains-go/pkg/testutils"
	v1 "github.com/selectel/domains-go/pkg/v1"
	"github.com/selectel/domains-go/pkg/v1/record"
)

func TestGet(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/v1/%d/records/%d", testDomainID, testRecordID),
		RawResponse: testGetResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		Token:      testutils.Token,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := record.Get(ctx, testClient, testDomainID, testRecordID)

	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Get method")
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusOK, httpResponse.StatusCode)
	}
	if !reflect.DeepEqual(expectedGetResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedGetResponse, actual)
	}
}

func TestGetHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/v1/%d/records/%d", testDomainID, testRecordID),
		RawResponse: testErrGenericResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		Token:      testutils.Token,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := record.Get(ctx, testClient, testDomainID, testRecordID)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no record from the Get method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestGetTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		Token:      testutils.Token,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := record.Get(ctx, testClient, testDomainID, testRecordID)

	if actual != nil {
		t.Fatal("expected no record from the Get method")
	}
	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestGetUnmarshallError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/v1/%d/records/%d", testDomainID, testRecordID),
		RawResponse: testSingleRecordInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		Token:      testutils.Token,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := record.Get(ctx, testClient, testDomainID, testRecordID)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no record from the Get method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestListByDomainID(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/v1/%d/records/", testDomainID),
		RawResponse: testListResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		Token:      testutils.Token,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := record.ListByDomainID(ctx, testClient, testDomainID)

	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the ListByDomainID method")
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusOK, httpResponse.StatusCode)
	}

	if !reflect.DeepEqual(expectedListResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedListResponse, actual)
	}
}

func TestListByDomainIDHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/v1/%d/records/", testDomainID),
		RawResponse: testErrGenericResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		Token:      testutils.Token,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := record.ListByDomainID(ctx, testClient, testDomainID)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no records from the ListByDomainID method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the ListByDomainID method")
	}
	if err == nil {
		t.Fatal("expected error from the ListByDomainID method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestListByDomainIDTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		Token:      testutils.Token,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := record.ListByDomainID(ctx, testClient, testDomainID)

	if actual != nil {
		t.Fatal("expected no records from the ListByDomainID method")
	}
	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the ListByDomainID method")
	}
	if err == nil {
		t.Fatal("expected error from the ListByDomainID method")
	}
}

func TestListByDomainIDUnmarshallError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/v1/%d/records/", testDomainID),
		RawResponse: testListRecordsInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		Token:      testutils.Token,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := record.ListByDomainID(ctx, testClient, testDomainID)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no records from the ListByDomainID method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the ListByDomainID method")
	}
	if err == nil {
		t.Fatal("expected error from the ListByDomainID method")
	}
}

func TestListByDomainName(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/v1/%s/records/", testDomainName),
		RawResponse: testListResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		Token:      testutils.Token,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := record.ListByDomainName(ctx, testClient, testDomainName)

	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the ListByDomainName method")
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusOK, httpResponse.StatusCode)
	}

	if !reflect.DeepEqual(expectedListResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedListResponse, actual)
	}
}

func TestListByDomainNameHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/v1/%s/records/", testDomainName),
		RawResponse: testErrGenericResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		Token:      testutils.Token,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := record.ListByDomainName(ctx, testClient, testDomainName)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no records from the ListByDomainName method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the ListByDomainName method")
	}
	if err == nil {
		t.Fatal("expected error from the ListByDomainName method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestListByDomainNameTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		Token:      testutils.Token,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := record.ListByDomainName(ctx, testClient, testDomainName)

	if actual != nil {
		t.Fatal("expected no records from the ListByDomainName method")
	}
	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the ListByDomainName method")
	}
	if err == nil {
		t.Fatal("expected error from the ListByDomainName method")
	}
}

func TestListByDomainNameUnmarshallError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/v1/%s/records/", testDomainName),
		RawResponse: testListRecordsInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		Token:      testutils.Token,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := record.ListByDomainName(ctx, testClient, testDomainName)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no records from the ListByDomainName method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the ListByDomainName method")
	}
	if err == nil {
		t.Fatal("expected error from the ListByDomainName method")
	}
}

func TestCreateRecord(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/v1/%d/records/", testDomainID),
		RawResponse: testCreateRecordResponseRaw,
		RawRequest:  testCreateRecordOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		Token:      testutils.Token,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := record.Create(ctx, testClient, testDomainID, testCreateRecordOpts)

	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Create method")
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusOK, httpResponse.StatusCode)
	}
	if !reflect.DeepEqual(expectedCreateResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedCreateResponse, actual)
	}
}

func TestCreateRecordHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/v1/%d/records/", testDomainID),
		RawResponse: testErrGenericResponseRaw,
		RawRequest:  testCreateRecordOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		Token:      testutils.Token,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := record.Create(ctx, testClient, testDomainID, testCreateRecordOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no record from the Create method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestCreateRecordTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		Token:      testutils.Token,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := record.Create(ctx, testClient, testDomainID, testCreateRecordOpts)

	if actual != nil {
		t.Fatal("expected no record from the Create method")
	}
	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestCreateRecordUnmarshallError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/v1/%d/records/", testDomainID),
		RawResponse: testSingleRecordInvalidResponseRaw,
		RawRequest:  testCreateRecordOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		Token:      testutils.Token,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := record.Create(ctx, testClient, testDomainID, testCreateRecordOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no record from the Create method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}
