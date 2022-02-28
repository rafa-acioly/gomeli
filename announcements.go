package meli

import (
	"net/http"
	"time"
)

// Announcement is the contract that represents a product
type Announcement interface {
	GetID() string
}

// AnnouncementManager is the contract to announcements
type AnnouncementManager interface {
	Create(a Announcement)
	Update(code string, data []string)
	Delete(code string)
	ChangeStatus(code, status string) // status should be a constant of available statuses
	SendDescription(code, desc string)
	ChangeDescription(code, desc string)
	AddVariation(code, v string)              // v should be a Variation type
	ChangeVariation(code string, vs []string) // vs should be a slice of Variation type
	DeleteVariation(code, variationCode string)
}

type announcementManager struct {
	m   Meli
	cli http.Client
}

// Create send a POST request to create a new product
func (am announcementManager) Create(a Announcement) {
	panic("implement me")
}

// Update send a PUT request to update a existing product
func (am announcementManager) Update(code string, data []string) {
	panic("implement me")
}

// Delete will remove completely a product
func (am announcementManager) Delete(code string) {
	panic("implement me")
}

// ChangeStatus wil change the current status of the product
// be aware that the status follow a status machine, once the status changes
// you cannot change it to the previous status
func (am announcementManager) ChangeStatus(code string, status string) {
	panic("implement me")
}

// SendDescription will set the description of the product.
// Since september 1st 2021 the product description should be defined on a separate request
// and no longer with the announcement creation
// more information: https://developers.mercadolivre.com.br/pt_br/publicacao-de-produtos#Description
func (am announcementManager) SendDescription(code, desc string) {
	panic("implement me")
}

// ChangeDescription will change a product description, some characters are not allowed
// by mercado's livre such as HTML tags.
// more information: https://developers.mercadolivre.com.br/pt_br/descricao-de-produtos
func (am announcementManager) ChangeDescription(code string, desc string) {
	panic("implement me")
}

// AddVariation will add a new variation on the product
// more information: https://developers.mercadolivre.com.br/pt_br/publicacao-de-produtos#Variacoes
func (am announcementManager) AddVariation(code string, v string) {
	panic("implement me")
}

// ChangeVariation will update an existing variation of the product
func (am announcementManager) ChangeVariation(code string, vs []string) {
	panic("implement me")
}

// DeleteVariation will remove completely a variation of a product
func (am announcementManager) DeleteVariation(code string, variationCode string) {
	panic("implement me")
}

// SetClient will redefine an HTTP Client used to make API calls
func (am announcementManager) SetClient(cli http.Client) {
	am.cli = cli
}

// NewAnnouncement return an announcement manager to handle announcements on mercado's livre API
func NewAnnouncement(m Meli) AnnouncementManager {
	return &announcementManager{
		m: m,
		cli: http.Client{
			Timeout: time.Minute,
		},
	}
}
