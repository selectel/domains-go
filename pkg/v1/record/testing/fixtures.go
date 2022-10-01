package testing

import (
	"github.com/selectel/domains-go/pkg/testutils"
	"github.com/selectel/domains-go/pkg/v1/record"
)

const (
	testDomainID   = 123
	testDomainName = "testdomain.xyz"
	testRecordID   = 321
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

// testListResponseRaw represents a raw response from the List request by ID.
const testListResponseRaw = `
[
   {
      "ttl" : 86400,
      "content" : "ns4.selectel.org",
      "type" : "NS",
      "id" : 4419121,
      "name" : "testdomain.xyz"
   },
   {
      "type" : "NS",
      "id" : 4419120,
      "name" : "testdomain.xyz",
      "ttl" : 86400,
      "content" : "ns3.selectel.org"
   },
   {
      "type" : "NS",
      "id" : 4419119,
      "name" : "testdomain.xyz",
      "ttl" : 86400,
      "content" : "ns2.selectel.org"
   },
   {
      "name" : "testdomain.xyz",
      "id" : 4419118,
      "type" : "NS",
      "content" : "ns1.selectel.org",
      "ttl" : 86400
   },
   {
      "type" : "CNAME",
      "id" : 4879135,
      "name" : "share.testdomain.xyz",
      "ttl" : 60,
      "content" : "origin.com"
   },
   {
      "ttl" : 60,
      "content" : "hello, world!",
      "type" : "TXT",
      "name" : "hello.testdomain.xyz",
      "id" : 4879136
   },
   {
      "content" : "ns1.selectel.org. support.selectel.ru. 2020050440 10800 3600 604800 300",
      "refresh" : 10800,
      "id" : 4419117,
      "name" : "testdomain.xyz",
      "type" : "SOA",
      "ns" : "ns1.selectel.org",
      "minimum" : 300,
      "email" : "support@selectel.ru",
      "ttl" : 60,
      "retry" : 3600,
      "change_date" : 1585642005,
      "expire" : 604800
   },
   {
      "id" : 4892420,
      "type" : "MX",
      "content" : "mail.example.com",
      "priority" : 20,
      "name" : "mx.testdomain.xyz",
      "ttl" : 86400
   },
   {
      "name" : "srv.testdomain.xyz",
      "ttl" : 86400,
      "target" : "xmpp.example.com",
      "weight" : 0,
      "id" : 4892452,
      "port" : 5222,
      "type" : "SRV",
      "priority" : 20
   },
   {
      "name" : "caa.testdomain.xyz",
      "ttl" : 86400,
      "id" : 4892462,
      "type" : "CAA",
      "tag" : "issue",
      "flag" : 1,
      "value" : "letsencrypt.com"
   },
   {
      "name" : "sshfp.testdomain.xyz",
      "ttl" : 86400,
      "id" : 4892472,
      "type" : "SSHFP",
      "algorithm" : 1,
      "fingerprint_type" : 1,
      "fingerprint" : "001a2B3CFF"
   },
   {
      "name" : "testdomain.zyx",
      "ttl" : 86400,
      "id" : 4892482,
      "type" : "ALIAS",
      "content" : "testdomain.xyz"
   }
]
`

// expectedListResponse represents an unmarshalled testListResponseRaw.
var expectedListResponse = []*record.View{
	{
		ID:      4419121,
		Name:    "testdomain.xyz",
		Type:    record.TypeNS,
		Content: "ns4.selectel.org",
		TTL:     86400,
	},
	{
		ID:      4419120,
		Name:    "testdomain.xyz",
		Type:    record.TypeNS,
		Content: "ns3.selectel.org",
		TTL:     86400,
	},
	{
		ID:      4419119,
		Name:    "testdomain.xyz",
		Type:    record.TypeNS,
		Content: "ns2.selectel.org",
		TTL:     86400,
	},
	{
		ID:      4419118,
		Name:    "testdomain.xyz",
		Type:    record.TypeNS,
		Content: "ns1.selectel.org",
		TTL:     86400,
	},
	{
		ID:      4879135,
		Type:    record.TypeCNAME,
		TTL:     60,
		Name:    "share.testdomain.xyz",
		Content: "origin.com",
	},
	{
		ID:      4879136,
		TTL:     60,
		Content: "hello, world!",
		Type:    record.TypeTXT,
		Name:    "hello.testdomain.xyz",
	},
	{
		ID:         4419117,
		Name:       "testdomain.xyz",
		Type:       record.TypeSOA,
		Content:    "ns1.selectel.org. support.selectel.ru. 2020050440 10800 3600 604800 300",
		Email:      "support@selectel.ru",
		TTL:        60,
		ChangeDate: testutils.IntPtr(1585642005),
	},
	{
		ID:       4892420,
		Type:     record.TypeMX,
		Content:  "mail.example.com",
		Priority: testutils.IntPtr(20),
		Name:     "mx.testdomain.xyz",
		TTL:      86400,
	},
	{
		ID:       4892452,
		Name:     "srv.testdomain.xyz",
		Type:     record.TypeSRV,
		Priority: testutils.IntPtr(20),
		Weight:   testutils.IntPtr(0),
		Port:     testutils.IntPtr(5222),
		Target:   "xmpp.example.com",
		TTL:      86400,
	},
   {
		ID:       4892462,
		Name:     "caa.testdomain.xyz",
		Type:     record.TypeCAA,
		Tag:      "issue",
		Flag:     testutils.IntPtr(1),
		Value:    "letsencrypt.com",
		TTL:      86400,
	},
   {
		ID:                  4892472,
		Name:                "sshfp.testdomain.xyz",
		Type:                record.TypeSSHFP,
		Algorithm:           testutils.IntPtr(1),
		FingerprintType:     testutils.IntPtr(1),
		Fingerprint:         "001a2B3CFF",
		TTL:                 86400,
	},
   {
		ID:       4892482,
		Name:     "testdomain.zyx",
		Type:     record.TypeALIAS,
		Content:  "testdomain.xyz",
		TTL:      86400,
	},
}

// testCreateRecordOptsRaw represents a raw request options for Create request.
const testCreateRecordOptsRaw = `
{
   "name": "example.testdomain.xyz",
   "type": "CNAME",
   "ttl": 60,
   "content": "origin.example.com"
}
`

// testCreateRecordOpts represents an unmarshalled testCreateRecordOptsRaw.
var testCreateRecordOpts = &record.CreateOpts{
	Name:    "example.testdomain.xyz",
	Type:    record.TypeCNAME,
	TTL:     60,
	Content: "origin.example.com",
}

// testCreateRecordResponseRaw represents a raw response for Create request.
const testCreateRecordResponseRaw = `
{
   "content" : "origin.example.com",
   "type" : "CNAME",
   "id" : 4894583,
   "ttl" : 60,
   "name" : "example.testdomain.xyz"
}
`

var expectedCreateResponse = &record.View{
	ID:      4894583,
	Name:    "example.testdomain.xyz",
	Type:    record.TypeCNAME,
	TTL:     60,
	Content: "origin.example.com",
}

// testUpdateRecordOptsRaw represents a raw request options for Update request.
const testUpdateRecordOptsRaw = `
{
   "name": "example.testdomain.xyz",
   "type": "CNAME",
   "ttl": 100,
   "content": "origin2.example.com"
}
`

// testUpdateRecordOpts represents an unmarshalled testUpdateRecordOptsRaw.
var testUpdateRecordOpts = &record.UpdateOpts{
	Name:    "example.testdomain.xyz",
	Type:    record.TypeCNAME,
	TTL:     100,
	Content: "origin2.example.com",
}

// testUpdateRecordResponseRaw represents a raw response for Update request.
const testUpdateRecordResponseRaw = `
{
   "content" : "origin2.example.com",
   "type" : "CNAME",
   "id" : 4894583,
   "ttl" : 100,
   "name" : "example.testdomain.xyz"
}
`

var expectedUpdateResponse = &record.View{
	ID:      4894583,
	Name:    "example.testdomain.xyz",
	Type:    record.TypeCNAME,
	TTL:     100,
	Content: "origin2.example.com",
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

// testListRecordsInvalidResponseRaw represents a raw invalid response with a list of records.
const testListRecordsInvalidResponseRaw = `
[{
   "" : "CNAME"
   "id" : 4879135,
   "ttl
   "name" : "share.testdomain.xyz",
}]
`
