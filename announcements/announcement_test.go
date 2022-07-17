package announcements

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"meli"
	"meli/environments"
	"testing"
)

type AnnouncementTestSuite struct {
	m       meli.Meli
	manager AnnouncementManager
	suite.Suite
}

func (test *AnnouncementTestSuite) SetupTest() {
	testEnv := environments.NewSandboxEnv(environments.BRASIL, environments.NewConfiguration(nil))
	test.m = meli.NewCustomClient("client-id", "client-secret", testEnv)
	test.manager = NewAnnouncement(test.m)
}

func (test *AnnouncementTestSuite) TestCanCreateNewAnnouncement() {
	item := Item{ID: "id"}

	_, err := test.manager.Create(item)

	assert.NoError(test.T(), err)
}

func TestAnnouncementManager(t *testing.T) {
	suite.Run(t, new(AnnouncementTestSuite))
}
