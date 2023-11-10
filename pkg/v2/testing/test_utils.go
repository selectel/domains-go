package testing

import (
	"context"
	"fmt"
	"net/http"

	v2 "github.com/selectel/domains-go/pkg/v2"
)

const (
	testDomainName = "bonnie-test.com"
	testUUID       = "a1b1a1e3-6cc2-4578-8ec7-5c4b2fcba3f7"
	testAPIURL     = "http://api.test.bonnie.com"
	testIPv4       = "10.20.30.40"
	testTTL        = 60
)

const (
	rootPath          = "/zones/"
	zonePath          = "/zones/%v"
	rrsetPath         = "/zones/%v/rrset"
	singleRRSetPath   = "/zones/%v/rrset/%v"
	rrsetManageByPath = "/zones/%v/rrset/%v/managed_by"
)

var (
	testHTTPClient = http.DefaultClient
	testCtx        = context.Background()
	testClient     = v2.NewClient(testAPIURL, testHTTPClient, make(http.Header))
)

func mockGetZoneResponse() string {
	return fmt.Sprintf(
		`
			{
				"uuid": "%v",
				"project_id": "%v",
				"name": "%v",
				"created_at": "2023-03-09T18:47:25Z",
				"updated_at": "2023-03-09T18:47:25Z",
				"last_check_status": false,
				"delegation_checked_at": null,
				"last_delegated_at": null
			}
			`,
		testUUID,
		testUUID,
		testDomainName,
	)
}

func mockListZonesResponse(count int) string {
	return fmt.Sprintf(
		`{
					"count": %v,
					"next_offset": null,
					"result": [%s, %s]
				}`,
		count,
		mockGetZoneResponse(),
		mockGetZoneResponse(),
	)
}

func mockGetRRSetResponse() string {
	return fmt.Sprintf(
		`{
					"uuid": "%v",
					"zone": "%v",
					"name": "go-test-record.%v",
					"type": "%v",
					"ttl": %v,
					"comment": "I am mock response from bonnie-go",
					"managed_by": null,
					"records": [{"content":"%v", "disabled": true},{"content":"%v", "disabled":false}]
				}`,
		testUUID,
		testUUID,
		testDomainName,
		v2.A,
		testTTL,
		testIPv4,
		testIPv4,
	)
}

func mockListRRSetResponse(count int) string {
	return fmt.Sprintf(
		`{
					"count": %v,
					"next_offset": null,
					"result": [%s, %s]
				}`,
		count,
		mockGetRRSetResponse(),
		mockGetRRSetResponse(),
	)
}

func mockCreateRRSetResponse() string {
	return mockGetRRSetResponse()
}
