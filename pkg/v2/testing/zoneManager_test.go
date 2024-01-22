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
	ZoneManageSuite struct {
		suite.Suite
	}
)

//nolint:paralleltest
func TestZoneManager(t *testing.T) {
	suite.Run(t, new(ZoneManageSuite))
}

func (s *ZoneManageSuite) SetupTest() {
	httpmock.Activate()
}

func (s *ZoneManageSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (s *ZoneManageSuite) TestCreateZone_ok() {
	httpmock.RegisterResponder(
		http.MethodPost,
		fmt.Sprintf("%s%s", testAPIURL, rootPath),
		httpmock.NewStringResponder(http.StatusOK, mockGetZoneResponse()),
	)

	//nolint: exhaustruct
	newZone := &v2.Zone{
		Name: testDomainName,
	}
	zone, err := testClient.CreateZone(testCtx, newZone)

	s.Nil(err)
	//nolint: exhaustruct
	s.IsType(&v2.Zone{}, zone)
	s.Equal(testDomainName, zone.Name)
	s.Equal(testUUID, zone.UUID)
	s.Equal(testUUID, zone.ProjectID)
}

func (s *ZoneManageSuite) TestGetZone_ok() {
	path := fmt.Sprintf(zonePath, testUUID)
	httpmock.RegisterResponder(
		http.MethodGet,
		fmt.Sprintf("%s%s", testAPIURL, path),
		httpmock.NewStringResponder(http.StatusOK, mockGetZoneResponse()),
	)

	zone, err := testClient.GetZone(testCtx, testUUID, nil)

	s.Nil(err)
	//nolint: exhaustruct
	s.IsType(&v2.Zone{}, zone)
	s.Equal(testDomainName, zone.Name)
	s.Equal(testUUID, zone.UUID)
	s.Equal(testUUID, zone.ProjectID)
}

func (s *ZoneManageSuite) TestListZones_ok() {
	testCount := 10
	httpmock.RegisterResponder(
		http.MethodGet,
		fmt.Sprintf("%s%s", testAPIURL, rootPath),
		httpmock.NewStringResponder(http.StatusOK, mockListZonesResponse(testCount)),
	)

	zoneList, err := testClient.ListZones(testCtx, nil)

	s.Nil(err)
	//nolint: exhaustruct
	s.IsType(&v2.List[v2.Zone]{}, zoneList)
	s.Len(zoneList.GetItems(), 2)
	s.Equal(testCount, zoneList.GetCount())
	s.Equal(0, zoneList.GetNextOffset())
}

func (s *ZoneManageSuite) TestDeleteZone_ok() {
	path := fmt.Sprintf(zonePath, testUUID)
	httpmock.RegisterResponder(
		http.MethodDelete,
		fmt.Sprintf("%s%s", testAPIURL, path),
		httpmock.NewBytesResponder(http.StatusNoContent, []byte{}),
	)

	err := testClient.DeleteZone(testCtx, testUUID)

	s.Nil(err)
}

func (s *ZoneManageSuite) TestCreateZone_err_conflict() {
	httpmock.RegisterResponder(
		http.MethodPost,
		fmt.Sprintf("%s%s", testAPIURL, rootPath),
		httpmock.NewStringResponder(http.StatusConflict, mockCreateZoneConflictResponse()),
	)

	//nolint: exhaustruct
	newZone := &v2.Zone{
		Name: testDomainName,
	}
	zone, err := testClient.CreateZone(testCtx, newZone)

	s.Empty(zone)
	expectedError := "error response: bad_request. Description: Conflict."
	s.EqualValues(err.Error(), expectedError)
}

func (s *ZoneManageSuite) TestCreateZone_err_bad_request_with_description_and_location() {
	httpmock.RegisterResponder(
		http.MethodPost,
		fmt.Sprintf("%s%s", testAPIURL, rootPath),
		httpmock.NewStringResponder(http.StatusBadRequest, mockCreateZoneFieldRequiredResponse()),
	)

	//nolint: exhaustruct
	newZone := &v2.Zone{
		Name: testDomainName,
	}
	zone, err := testClient.CreateZone(testCtx, newZone)

	s.Empty(zone)
	expectedError := "error response: bad_request. Description: field required. Location: body.name."
	s.EqualValues(err.Error(), expectedError)
}
