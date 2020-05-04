package record

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	v1 "github.com/selectel/domains-go/pkg/v1"
)

// Get returns a single record by its id.
func Get(ctx context.Context, client *v1.ServiceClient, domainID, recordID int) (*View, *v1.ResponseResult, error) {
	url := strings.Join([]string{
		client.Endpoint,
		strconv.Itoa(domainID),
		v1.RecordsEndpoint,
		strconv.Itoa(recordID)}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a single record from the response body.
	record := &View{}
	err = responseResult.ExtractResult(record)
	if err != nil {
		return nil, responseResult, err
	}

	return record, responseResult, nil
}
