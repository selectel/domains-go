# domains-go: Go SDK for Selectel Domains API
[![Go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/selectel/domains-go/)
[![Go Report Card](https://goreportcard.com/badge/github.com/selectel/domains-go)](https://goreportcard.com/report/github.com/selectel/domains-go)
[![Build Status](https://travis-ci.com/selectel/domains-go.svg?branch=master)](https://travis-ci.com/selectel/domains-go)
[![Coverage Status](https://coveralls.io/repos/github/selectel/domains-go/badge.svg?branch=master)](https://coveralls.io/github/selectel/domains-go?branch=master)

Package domains-go provides Go SDK to work with the Selectel Domains API.

## Documentation

The Go library documentation is available at [go.dev](https://pkg.go.dev/github.com/selectel/domains-go/).

## What this library is capable of

You can use this library to work with the following objects of the Selectel Domains API:

* [domain](https://pkg.go.dev/github.com/selectel/domains-go/pkg/v1/domain)
* [record](https://pkg.go.dev/github.com/selectel/domains-go/pkg/v1/record)

## Getting started

### Installation

You can install needed `domains-go` packages via `go get` command:

```bash
go get github.com/selectel/domains-go/pkg/v1/domain github.com/selectel/domains-go/pkg/v1/record
```

### Authentication

To work with the Selectel Domains API you first need to:

* Create a Selectel account: [registration page](https://my.selectel.ru/registration).
* Create an API token: https://my.selectel.ru/profile/apikeys

### Usage example

```go
package main

import (
	"context"
	"fmt"
	"log"

	v1 "github.com/selectel/domains-go/pkg/v1"
	"github.com/selectel/domains-go/pkg/v1/domain"
	"github.com/selectel/domains-go/pkg/v1/record"
)

func main() {
	// Token to work with Selectel Cloud project
	token := "gAAAAABeVNzu-..."

	// Domains API V1 endpoint to work with
	endpoint := "https://api.selectel.ru/domains/v1"

	// Initialize the Domains API V1 client
	client := v1.NewDomainsClientV1(token, endpoint)

	createDomainOpts := &domain.CreateOpts{
		Name: "testdomain.xyz",
	}

	// Create domain
	selectelDomain, _, err := domain.Create(context.Background(), client, createDomainOpts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created domain: %+v\n", selectelDomain)

	// List domains
	selectelDomains, _, err := domain.List(context.Background(), client)
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range selectelDomains {
		fmt.Printf("%+v\n", d)
	}

	createRecordOpts := &record.CreateOpts{
		Name:    "share.testdomain.xyz",
		Type:    record.TypeCNAME,
		TTL:     60,
		Content: "origin.example.com",
	}

	// Create domain record
	domainRecord, _, err := record.Create(context.Background(), client, selectelDomain.ID, createRecordOpts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created record: %+v\n", domainRecord)
}
```
