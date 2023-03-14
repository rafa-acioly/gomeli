package announcements

import (
	"encoding/json"
	"fmt"
	"meli"
)

// AnnouncementManager is the contract to announcements
// announcement means some sort of offer made on the marketplace
// it can be a product or an auction
type AnnouncementManager interface {
	Create(a AnnouncementRequest) (Announcement, error)
	Update(code string, data map[string]string) (Announcement, error)
	Delete(code string) error
	ChangeStatus(code string, s status) error
	SendDescription(code, desc string) error
	ChangeDescription(code, desc string) error
	AddVariation(code string, v Variation) error
	ChangeVariation(code string, vs []Variation) error
	DeleteVariation(code, variationCode string) error
	UseClient(cli meli.HttpClientWrapper)
}

type announcementManager struct {
	m   meli.Meli
	cli meli.HttpClientWrapper
}

// Create send a POST request to create a new product
// [more information]: http://developers.mercadolibre.com/list-products/#ListAnItem
func (manager announcementManager) Create(a AnnouncementRequest) (Announcement, error) {
	response, err := manager.cli.Post("/items", a)
	if err != nil {
		return Announcement{}, err
	}

	var item Announcement
	if err := json.Unmarshal(response.Body(), &item); err != nil {
		return Announcement{}, err
	}

	return item, nil
}

// Update send a PUT request to update a existing product
// [more information]: http://developers.mercadolibre.com/products-sync-listings/#Update-your-item
func (manager announcementManager) Update(code string, data map[string]string) (Announcement, error) {
	response, err := manager.cli.Put("/items"+code, data)
	if err != nil {
		return Announcement{}, err
	}

	var item Announcement
	if err := json.Unmarshal(response.Body(), &item); err != nil {
		return Announcement{}, err
	}

	return item, nil
}

// Delete will mark a product as deleted
// [more information]: http://developers.mercadolibre.com/products-sync-listings/#Delete-listing
func (manager announcementManager) Delete(code string) error {
	requestContent := map[string]string{"deleted": "true"}
	_, err := manager.cli.Delete("/item"+code, requestContent)
	if err != nil {
		return err
	}

	return nil
}

// ChangeStatus wil change the current status of the product
// be aware that the status follow a status machine, once the status changes
// you cannot change it to the previous status
// [more information]: http://developers.mercadolibre.com/products-sync-listings/#Changing-listing-status
func (manager announcementManager) ChangeStatus(code string, s status) error {
	requestContent := map[string]string{"status": string(s)}
	_, err := manager.cli.Put("/items"+code, requestContent)
	if err != nil {
		return err
	}

	return nil
}

// SendDescription will set the description of the product.
// Since september 1st 2021 the product description should be defined on a separate request
// and no longer with the announcement creation
// [more information]: https://developers.mercadolivre.com.br/pt_br/publicacao-de-produtos#Description
func (manager announcementManager) SendDescription(code, desc string) error {
	return manager.ChangeDescription(code, desc)
}

// ChangeDescription will change a product description, some characters are not allowed
// by mercado's livre such as HTML tags.
// [more information]: https://developers.mercadolivre.com.br/pt_br/descricao-de-produtos
func (manager announcementManager) ChangeDescription(code string, desc string) error {
	requestContent := map[string]string{"plain_text": desc}
	_, err := manager.cli.Put("/items"+code, requestContent)
	if err != nil {
		return err
	}

	return nil
}

// AddVariation will add a new variation on the product
// [more information]: https://developers.mercadolivre.com.br/pt_br/publicacao-de-produtos#Variacoes
func (manager announcementManager) AddVariation(code string, v Variation) error {
	_, err := manager.cli.Put("/items"+code, v)
	if err != nil {
		return err
	}

	return nil
}

// ChangeVariation will change all product variations, this has the same behavior as updating
// [more information]: https://developers.mercadolibre.com/pt_br/variacoes#Modificar-varia%C3%A7%C3%B5es
func (manager announcementManager) ChangeVariation(code string, vs []Variation) error {
	_, err := manager.cli.Put("/items"+code, vs)
	if err != nil {
		return err
	}

	return nil
}

// DeleteVariation will remove completely a variation of a product
// [more information]: https://developers.mercadolibre.com/pt_br/variacoes#Remover-varia%C3%A7%C3%B5es
func (manager announcementManager) DeleteVariation(code string, variationCode string) error {
	_, err := manager.cli.Delete(fmt.Sprintf("/items/%s/variations/%s/", code, variationCode), nil)
	if err != nil {
		return err
	}

	return nil
}

// UseClient will redefine an HTTP Client used to make API calls
func (manager *announcementManager) UseClient(cli meli.HttpClientWrapper) {
	manager.cli = cli
}

// NewAnnouncement return an announcement manager to handle announcements on mercado's livre API
func NewAnnouncementManager(m meli.Meli) AnnouncementManager {
	return &announcementManager{
		m:   m,
		cli: meli.NewHttpClient(m),
	}
}
