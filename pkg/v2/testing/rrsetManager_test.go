package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	v2 "github.com/selectel/domains-go/pkg/v2"
	"github.com/stretchr/testify/suite"
)

type (
	RRSetManagerSuite struct {
		suite.Suite
	}
)

//nolint:paralleltest
func TestRRSetManager(t *testing.T) {
	suite.Run(t, new(RRSetManagerSuite))
}

func (s *RRSetManagerSuite) SetupTest() {
	httpmock.Activate()
}

func (s *RRSetManagerSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (s *RRSetManagerSuite) TestGetRRSet_ok() {
	path := fmt.Sprintf(singleRRSetPath, testUUID, testUUID)
	httpmock.RegisterResponder(
		http.MethodGet,
		fmt.Sprintf("%s%s", testAPIURL, path),
		httpmock.NewStringResponder(http.StatusOK, mockGetRRSetResponse()),
	)

	rrset, err := testClient.GetRRSet(testCtx, testUUID, testUUID)

	s.Nil(err)
	//nolint: exhaustruct
	s.IsType(&v2.RRSet{}, rrset)
	s.NotNil(rrset.Records)
	s.Equal(testUUID, rrset.UUID)
	s.Len(rrset.Records, 2)
}

func (s *RRSetManagerSuite) TestListRRSets_ok() {
	path := fmt.Sprintf(rrsetPath, testUUID)
	testCount := 10
	httpmock.RegisterResponder(
		http.MethodGet,
		fmt.Sprintf("%s%s", testAPIURL, path),
		httpmock.NewStringResponder(http.StatusOK, mockListRRSetResponse(testCount)),
	)

	rrsetList, err := testClient.ListRRSets(testCtx, testUUID, nil)

	s.Nil(err)
	//nolint: exhaustruct
	s.IsType(&v2.List[v2.RRSet]{}, rrsetList)
	s.Len(rrsetList.GetItems(), 2)
	s.Equal(testCount, rrsetList.GetCount())
	s.Equal(0, rrsetList.GetNextOffset())
}

func (s *RRSetManagerSuite) TestSetRRSetManagedBy_ok() {
	path := fmt.Sprintf(rrsetManageByPath, testUUID, testUUID)
	httpmock.RegisterResponder(
		http.MethodPost,
		fmt.Sprintf("%s%s", testAPIURL, path),
		httpmock.NewBytesResponder(http.StatusNoContent, []byte{}),
	)

	err := testClient.SetRRSetManagedBy(testCtx, testUUID, testUUID, testUUID)
	s.Nil(err)
}

func (s *RRSetManagerSuite) TestResetRRSetManagedBy_ok() {
	path := fmt.Sprintf(rrsetManageByPath, testUUID, testUUID)
	httpmock.RegisterResponder(
		http.MethodDelete,
		fmt.Sprintf("%s%s", testAPIURL, path),
		httpmock.NewBytesResponder(http.StatusNoContent, []byte{}),
	)

	err := testClient.ResetRRSetManagedBy(testCtx, testUUID, testUUID)
	s.Nil(err)
}

func (s *RRSetManagerSuite) TestCreateRRSet_ok() {
	path := fmt.Sprintf(rrsetPath, testUUID)
	httpmock.RegisterResponder(
		http.MethodPost,
		fmt.Sprintf("%s%s", testAPIURL, path),
		httpmock.NewStringResponder(http.StatusOK, mockCreateRRSetResponse()),
	)

	//nolint: exhaustruct
	newRRSet := &v2.RRSet{
		Name: "example.com",
		Type: v2.A,
		TTL:  60,
		Records: []v2.RecordItem{
			{Content: "", Disabled: false},
		},
	}
	rrset, err := testClient.CreateRRSet(testCtx, testUUID, newRRSet)

	s.Nil(err)
	//nolint: exhaustruct
	s.IsType(&v2.RRSet{}, rrset)
	s.Equal(testUUID, rrset.UUID)
}

func (s *RRSetManagerSuite) TestDeleteRRSet_ok() {
	path := fmt.Sprintf(singleRRSetPath, testUUID, testUUID)
	httpmock.RegisterResponder(
		http.MethodDelete,
		fmt.Sprintf("%s%s", testAPIURL, path),
		httpmock.NewBytesResponder(http.StatusNoContent, []byte{}),
	)

	err := testClient.DeleteRRSet(testCtx, testUUID, testUUID)
	s.Nil(err)
}

func (s *RRSetManagerSuite) TestUpdateRRSet_ok() {
	path := fmt.Sprintf(singleRRSetPath, testUUID, testUUID)
	httpmock.RegisterResponder(
		http.MethodPatch,
		fmt.Sprintf("%s%s", testAPIURL, path),
		httpmock.NewBytesResponder(http.StatusNoContent, []byte{}),
	)

	//nolint: exhaustruct
	changeForm := &v2.RRSet{
		TTL: testTTL,
		Records: []v2.RecordItem{
			{Content: "content", Disabled: false},
		},
	}
	err := testClient.UpdateRRSet(testCtx, testUUID, testUUID, changeForm)
	s.Nil(err)
}
