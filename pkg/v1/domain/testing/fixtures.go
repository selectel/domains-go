package testing

import "github.com/selectel/domains-go/pkg/v1/domain"

const (
	testDomainID   = 123
	testDomainName = "testdomain.xyz"
)

// testGetByIDResponseRaw represents a raw response from the Get request by ID.
const testGetByIDResponseRaw = `
{
   "tags" : ["test-tag1", "test-tag2"],
   "user_id" : 12345,
   "name" : "testdomain.xyz",
   "create_date" : 1585639110,
   "id" : 225474,
   "change_date" : 1585643272
}
`

// expectedGetByIDResponse represents an unmarshalled testGetDomainByIDResponseRaw.
var expectedGetByIDResponse = &domain.View{
	ID:         225474,
	UserID:     12345,
	Name:       "testdomain.xyz",
	CreateDate: 1585639110,
	ChangeDate: 1585643272,
	Tags:       []string{"test-tag1", "test-tag2"},
}

// testListResponseRaw represents a raw response from the List request by ID.
const testListResponseRaw = `
[
   {
      "id" : 225474,
      "change_date" : 1585643272,
      "user_id" : 12345,
      "name" : "testdomain.xyz",
      "create_date" : 1585639110,
      "tags" : ["test-tag1", "test-tag2"]
   },
   {
      "id" : 225475,
      "change_date" : 1585643272,
      "user_id" : 12345,
      "name" : "testdomain2.xyz",
      "create_date" : 1585639110,
      "tags" : ["test-tag3"]
   }
]
`

// expectedListResponse represents an unmarshalled testListResponseRaw.
var expectedListResponse = []*domain.View{
	{
		ID:         225474,
		UserID:     12345,
		Name:       "testdomain.xyz",
		CreateDate: 1585639110,
		ChangeDate: 1585643272,
		Tags:       []string{"test-tag1", "test-tag2"},
	},
	{
		ID:         225475,
		UserID:     12345,
		Name:       "testdomain2.xyz",
		CreateDate: 1585639110,
		ChangeDate: 1585643272,
		Tags:       []string{"test-tag3"},
	},
}

// testCreateDomainResponseRaw represents a raw response from the Create request.
const testCreateDomainResponseRaw = `
{
   "create_date" : 1585643272,
   "user_id" : 12345,
   "name" : "testdomain.xyz",
   "tags" : [],
   "id" : 225474,
   "change_date" : 1585639110
}
`

// expectedListResponse represents an unmarshalled testListResponseRaw.
var expectedCreateDomainResponse = &domain.View{
	ID:         225474,
	UserID:     12345,
	Name:       "testdomain.xyz",
	CreateDate: 1585643272,
	ChangeDate: 1585639110,
	Tags:       []string{},
}

// testCreateDomainOptsRaw represents marshalled options for the Create request.
const testCreateDomainOptsRaw = `
{
	"name": "testdomain.xyz"
}
`

// testCreateDomainOpts represents options for the Create request.
var testCreateDomainOpts = &domain.CreateOpts{
	Name: "testdomain.xyz",
}

// testCreateDomainWithFileZoneResponseRaw represents a raw response from the Create request.
const testCreateDomainWithBindZoneResponseRaw = `
{
   "domain" : {
      "create_date" : 1585643272,
      "user_id" : 12345,
      "name" : "testdomain.xyz",
      "tags" : [],
      "id" : 225474,
      "change_date" : 1585639110
   },
   "records" : []
}
`

// expectedListResponse represents an unmarshalled testListResponseRaw.
var expectedCreateDomainWithBindZoneResponse = &domain.View{
	ID:         225474,
	UserID:     12345,
	Name:       "testdomain.xyz",
	CreateDate: 1585643272,
	ChangeDate: 1585639110,
	Tags:       []string{},
}

// testCreateDomainOptsRaw represents marshalled options for the Create request.
const testCreateDomainWithBindZoneOptsRaw = `
{
     "name": "testdomain.xyz",
     "bind_zone": "@ IN SOA ns.test.org. support.selectel.ru. (2020050349 10800 3600 604800 300)"
}
`

// testCreateDomainOpts represents options for the Create request.
var testCreateDomainWithBindZoneOpts = &domain.CreateOpts{
	Name:     "testdomain.xyz",
	BindZone: "@ IN SOA ns.test.org. support.selectel.ru. (2020050349 10800 3600 604800 300)",
}

// testErrGenericResponseRaw represents a raw response with an error in the generic format.
const testErrGenericResponseRaw = `{"error":"bad gateway"}`

// testSingleDomainInvalidResponseRaw represents a raw invalid response with a single domain.
const testSingleDomainInvalidResponseRaw = `
{
   "tags" : "["test-tag1", "test-tag2"]",
   "user_id" :
   "name" : .xyz",
   "create_date" : 1585639110,
   "id" : "",
   "change_date" : -1
}
`

// testListInvalidResponseRaw represents a raw invalid response with a list of domains.
const testListInvalidResponseRaw = `
[
	{
	   "tags" : "["test-tag1", "test-tag2"]",
	   "user_id" :
	   "name" : .xyz",
	   "create_date" : 1585639110,
	   "id" : "",
	   "change_date" : -1
	},
]
`
