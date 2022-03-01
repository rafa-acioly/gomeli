package meli

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"meli/announcements"
	"meli/environments"
)

type AnnouncementTestSuite struct {
	m       Meli
	manager AnnouncementManager
	suite.Suite
}

func (test AnnouncementTestSuite) SetupTest() {
	testEnv := environments.NewSandboxEnv(environments.BRASIL, environments.NewConfiguration(nil))
	test.m = NewClient("client-id", "client-secret", testEnv)
	test.manager = NewAnnouncement(test.m)
}

func (test *AnnouncementTestSuite) TestCanCreateNewAnnouncement() {
	item := announcements.Item{ID: "id"}

	_, err := test.manager.Create(item)

	assert.NoError(test.T(), err)
}
