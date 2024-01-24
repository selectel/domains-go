# domains-go: Go SDK for Selectel Domains API
[![Go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/selectel/domains-go/)
[![Go Report Card](https://goreportcard.com/badge/github.com/selectel/domains-go)](https://goreportcard.com/report/github.com/selectel/domains-go)
[![Build Status](https://travis-ci.com/selectel/domains-go.svg?branch=master)](https://travis-ci.com/selectel/domains-go)
[![Coverage Status](https://coveralls.io/repos/github/selectel/domains-go/badge.svg?branch=master)](https://coveralls.io/github/selectel/domains-go?branch=master)

Package domains-go provides Go SDK to work with the Selectel Domains API.

## Contents

* [Documentation](#documentation)
* [Installation](#installation)
* [Authentication](#authentication)
* [Usage example](#usage-example)
* [Current version vs Legacy version](#current-version-vs-legacy-version)
* [Usage legacy example](#usage-legacy-example)

## Documentation

The Go library documentation is available at [go.dev](https://pkg.go.dev/github.com/selectel/domains-go/).

## What this library is capable of

You can use this library to work with the following objects of the Selectel Domains API:

* [zone](https://pkg.go.dev/github.com/selectel/domains-go/pkg/v2/#Zone)
* [rrset](https://pkg.go.dev/github.com/selectel/domains-go/pkg/v2/#RRSet)

## Getting started

### Installation

You can install needed `domains-go` packages via `go get` command:

```bash
go get github.com/selectel/domains-go
```

### Authentication

To work with the Selectel Domains API you first need to:

* Create a Selectel account: [registration page](https://my.selectel.ru/registration).
* For **current** version create an [Keystone Project Token](https://developers.selectel.com/docs/control-panel/authorization/#project-token)
* For **legacy** version create an [Selectel Token](https://developers.selectel.com/docs/control-panel/authorization/#selectel-token-api-key)

❗️IMPORTANT❗️  
`Selectel Token` and `Keystone Project Token` are **different** tokens!  
Above we mentioned how to get keystone project token, how to obtain selectel token read [here](https://developers.selectel.com/docs/control-panel/authorization)

### Usage example

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	v2 "github.com/selectel/domains-go/pkg/v2"
)

func main() {
	// Keystone project token. Read more in authorization.
	token := "gAAAAABeVNzu-..."

	// Domains API V2 endpoint to work with
	endpoint := "https://api.selectel.ru/domains/v2"

	httpClient := &http.Client{}
	userAgent := "domains-go-v2"
	hdrs := http.Header{}
	hdrs.Add("X-Auth-Token", token)
	hdrs.Add("User-Agent", userAgent)
	// Initialize the Domains API V2 client
	client := v2.NewClient(endpoint, httpClient, hdrs)

	createZoneOpts := &v2.Zone{
		Name: "domains-go-v2.ru.",
	}

	// Create zone
	selectelCreatedZone, err := client.CreateZone(context.Background(), createZoneOpts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created zone: %+v\n", selectelCreatedZone)

	listZonesOpts := &map[string]string{}
	// List zones
	selectelZones, err := client.ListZones(context.Background(), listZonesOpts)
	if err != nil {
		log.Fatal(err)
	}

	for _, zone := range selectelZones.GetItems() {
		fmt.Printf("%+v\n", zone)
	}

	createRrsetOpts := &v2.RRSet{
		Name: "txt.domains-go-v2.ru.",
		Type: v2.TXT,
		TTL:  60,
		Records: []v2.RecordItem{
			// Only for TXT Rrset escaping quotes
			{Content: "\"Hello world!\""},
		},
	}

	// Create rrset type TXT
	selectelCreatedRrset, err := client.CreateRRSet(context.Background(), selectelCreatedZone.ID, createRrsetOpts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created rrset: %+v\n", selectelCreatedRrset)
}
```

## Current version vs Legacy version

Current version is `github.com/selectel/domains-go/pkg/v2`  
Legacy version is `github.com/selectel/domains-go/pkg/v1`  

They are not compatible. They utilize different API and created zones live on different authoritative servers.
Zone created in v2 API with current version is entirely new zone, and not available via v1 api and vice versa.  

If you are going to create new zone, we strongly recommend to use `github.com/selectel/domains-go/pkg/v2`.  
If you have zones in v1, you still can manage them with `github.com/selectel/domains-go/pkg/v1`.

Legacy version following objects of the Selectel Domains API v1:

* [domain](https://pkg.go.dev/github.com/selectel/domains-go/pkg/v1/domain)
* [record](https://pkg.go.dev/github.com/selectel/domains-go/pkg/v1/record)

Current version following objects of the Selectel Domains API v2:

* [zone](https://pkg.go.dev/github.com/selectel/domains-go/pkg/v2/#Zone)
* [rrset](https://pkg.go.dev/github.com/selectel/domains-go/pkg/v2/#RRSet)

### Usage legacy example

❗️IMPORTANT❗️
We are not recommending using this example.

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
