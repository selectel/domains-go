package domain

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	v1 "github.com/selectel/domains-go/pkg/v1"
)

// GetByID returns a single domain by its id.
func GetByID(ctx context.Context, client *v1.ServiceClient, domainID int) (*View, *v1.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, strconv.Itoa(domainID)}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a domain from the response body.
	domain := &View{}
	err = responseResult.ExtractResult(domain)
	if err != nil {
		return nil, responseResult, err
	}

	return domain, responseResult, nil
}

// GetByName returns a single domain by its domain name.
func GetByName(ctx context.Context, client *v1.ServiceClient, domainName string) (*View, *v1.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, domainName}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a domain from the response body.
	domain := &View{}
	err = responseResult.ExtractResult(domain)
	if err != nil {
		return nil, responseResult, err
	}

	return domain, responseResult, nil
}
