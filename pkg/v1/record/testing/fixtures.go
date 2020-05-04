package testing

import (
	"github.com/selectel/domains-go/pkg/v1/record"
)

const (
	testDomainID = 123
	testRecordID = 321
)

// testGetResponseRaw represents a raw response from the Get request by ID.
const testGetResponseRaw = `
{
   "type" : "CNAME",
   "id" : 4879135,
   "ttl" : 60,
   "name" : "share.testdomain.xyz",
   "content" : "origin.com"
}
`

// expectedGetResponse represents an unmarshalled testGetResponseRaw.
var expectedGetResponse = &record.View{
	ID:      4879135,
	Type:    record.TypeCNAME,
	TTL:     60,
	Name:    "share.testdomain.xyz",
	Content: "origin.com",
}

// testErrGenericResponseRaw represents a raw response with an error in the generic format.
const testErrGenericResponseRaw = `{"error":"bad gateway"}`

// testSingleRecordInvalidResponseRaw represents a raw invalid response with a single record.
const testSingleRecordInvalidResponseRaw = `
{
   "" : "CNAME"
   "id" : 4879135,
   "ttl
   "name" : "share.testdomain.xyz",
}
`
