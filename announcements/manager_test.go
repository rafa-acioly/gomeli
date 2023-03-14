package announcements

import (
	"meli"
	"meli/environments"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AnnouncementTestSuite struct {
	m       meli.Meli
	manager AnnouncementManager
	suite.Suite
}

func (test *AnnouncementTestSuite) SetupTest() {
	testEnv := environments.NewSandboxEnv(environments.BRASIL, environments.NewConfiguration(nil))
	test.m = meli.NewCustomClient("client-id", "client-secret", "random-tenant", testEnv)
	test.manager = NewAnnouncementManager(test.m)
}

func (test *AnnouncementTestSuite) TestCanCreateNewAnnouncement() {
	item := AnnouncementRequest{ID: "id"}

	_, err := test.manager.Create(item)

	assert.NoError(test.T(), err)
}

func TestAnnouncementManager(t *testing.T) {
	suite.Run(t, new(AnnouncementTestSuite))
}
