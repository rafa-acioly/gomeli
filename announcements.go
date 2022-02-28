package meli

import "net/http"

type Announcement interface {
	GetID() string
}

type AnnouncementManager interface {
	Create(a Announcement)
	Update(code string, data []string)
	Delete(code string)
	ChangeStatus(code string, status string) // status should be a constant of available statuses
	ChangeDescription(code string, desc string)
	AddVariation(code string, v string)       // v should be a Variation type
	ChangeVariation(code string, vs []string) // vs should be a slice of Variation type
	DeleteVariation(code string, variationCode string)
}

type announcementManager struct {
	m   Meli
	cli http.Client
}

func (am announcementManager) Create(a Announcement) {
	panic("implement me")
}

func (am announcementManager) Update(code string, data []string) {
	panic("implement me")
}

func (am announcementManager) Delete(code string) {
	panic("implement me")
}

func (am announcementManager) ChangeStatus(code string, status string) {
	panic("implement me")
}

func (am announcementManager) ChangeDescription(code string, desc string) {
	panic("implement me")
}

func (am announcementManager) AddVariation(code string, v string) {
	panic("implement me")
}

func (am announcementManager) ChangeVariation(code string, vs []string) {
	panic("implement me")
}

func (am announcementManager) DeleteVariation(code string, variationCode string) {
	panic("implement me")
}

func (am announcementManager) SetClient(cli http.Client) {
	am.cli = cli
}

func NewAnnouncement(m Meli) AnnouncementManager {
	return &announcementManager{
		m: m,
	}
}
