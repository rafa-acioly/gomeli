package meli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"meli/announcements"
	"meli/responses"
	"net/http"
	"time"
)

// AnnouncementManager is the contract to announcements
// announcement means some sort of offer made on the marketplace
// it can be a product or an auction
type AnnouncementManager interface {
	Create(a announcements.Item) (responses.Item, error)
	Update(code string, data map[string]string) (responses.Item, error)
	Delete(code string) error
	ChangeStatus(code, status string) error // status should be a constant of available statuses
	SendDescription(code, desc string) error
	ChangeDescription(code, desc string) error
	AddVariation(code string, v announcements.Variation) error
	ChangeVariation(code string, vs []announcements.Variation) error
	DeleteVariation(code, variationCode string) error
	UseClient(cli http.Client)
}

type announcementManager struct {
	m   Meli
	cli http.Client
}

// Create send a POST request to create a new product
// more information: http://developers.mercadolibre.com/list-products/#ListAnItem
func (manager announcementManager) Create(a announcements.Item) (responses.Item, error) {
	content, _ := json.Marshal(a)
	response, err := manager.cli.Post(
		manager.m.GetEnvironment().GetWsURL("/items"),
		"application/json",
		bytes.NewBuffer(content),
	)
	if err != nil {
		return responses.Item{}, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return responses.Item{}, err
	}

	var item responses.Item
	err = json.Unmarshal(body, &item)
	if err != nil {
		return responses.Item{}, err
	}

	return item, nil
}

// Update send a PUT request to update a existing product
// more information: http://developers.mercadolibre.com/products-sync-listings/#Update-your-item
func (manager announcementManager) Update(code string, data map[string]string) (responses.Item, error) {
	content, _ := json.Marshal(data)
	request, _ := http.NewRequest(
		http.MethodPut,
		manager.m.GetEnvironment().GetWsURL("/items/"+code),
		bytes.NewBuffer(content),
	)
	response, err := manager.cli.Do(request)
	if err != nil {
		return responses.Item{}, err
	}

	if response.StatusCode != http.StatusOK {
		return responses.Item{}, fmt.Errorf("error on update product: %s", response.Status)
	}

	return responses.Item{}, nil
}

// Delete will mark a product as deleted
// more information: http://developers.mercadolibre.com/products-sync-listings/#Delete-listing
func (manager announcementManager) Delete(code string) error {
	content, _ := json.Marshal(map[string]string{"deleted": "true"})
	request, _ := http.NewRequest(
		http.MethodPut,
		manager.m.GetEnvironment().GetWsURL("/items/"+code),
		bytes.NewBuffer(content),
	)
	response, err := manager.cli.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error on delete product: %s", response.Status)
	}

	return nil
}

// ChangeStatus wil change the current status of the product
// be aware that the status follow a status machine, once the status changes
// you cannot change it to the previous status
// more information: http://developers.mercadolibre.com/products-sync-listings/#Changing-listing-status
func (manager announcementManager) ChangeStatus(code string, status string) error {
	content, _ := json.Marshal(map[string]string{"status": status})
	request, _ := http.NewRequest(
		http.MethodPut,
		manager.m.GetEnvironment().GetWsURL("/items/"+code),
		bytes.NewBuffer(content),
	)
	response, err := manager.cli.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error on change status: %s", response.Status)
	}

	return nil
}

// SendDescription will set the description of the product.
// Since september 1st 2021 the product description should be defined on a separate request
// and no longer with the announcement creation
// more information: https://developers.mercadolivre.com.br/pt_br/publicacao-de-produtos#Description
func (manager announcementManager) SendDescription(code, desc string) error {
	return manager.ChangeDescription(code, desc)
}

// ChangeDescription will change a product description, some characters are not allowed
// by mercado's livre such as HTML tags.
// more information: https://developers.mercadolivre.com.br/pt_br/descricao-de-produtos
func (manager announcementManager) ChangeDescription(code string, desc string) error {
	content, _ := json.Marshal(map[string]string{"plain_text": desc})
	request, _ := http.NewRequest(
		http.MethodPut,
		manager.m.GetEnvironment().GetWsURL("/items/"+code+"/description"),
		bytes.NewBuffer(content),
	)
	response, err := manager.cli.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error on change description: %s", response.Status)
	}

	return nil
}

// AddVariation will add a new variation on the product
// more information: https://developers.mercadolivre.com.br/pt_br/publicacao-de-produtos#Variacoes
func (manager announcementManager) AddVariation(code string, v announcements.Variation) error {
	content, _ := json.Marshal(v)
	request, _ := http.NewRequest(
		http.MethodPut,
		manager.m.GetEnvironment().GetWsURL("/items"+code+"/variations"),
		bytes.NewBuffer(content),
	)

	response, err := manager.cli.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error on adding variation: %s", response.Status)
	}

	return nil
}

// ChangeVariation will change all product variations, this has the same behaviour as updating
// more information: https://developers.mercadolibre.com/pt_br/variacoes#Modificar-varia%C3%A7%C3%B5es
func (manager announcementManager) ChangeVariation(code string, vs []announcements.Variation) error {
	content, _ := json.Marshal(vs)
	request, _ := http.NewRequest(
		http.MethodPut,
		manager.m.GetEnvironment().GetWsURL("/items"+code),
		bytes.NewBuffer(content),
	)

	response, err := manager.cli.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error on changing variations: %s", response.StatusCode)
	}

	return nil
}

// DeleteVariation will remove completely a variation of a product
// more information: https://developers.mercadolibre.com/pt_br/variacoes#Remover-varia%C3%A7%C3%B5es
func (manager announcementManager) DeleteVariation(code string, variationCode string) error {
	request, _ := http.NewRequest(
		http.MethodDelete,
		manager.m.GetEnvironment().GetWsURL("/items"+code+"/variations/"+variationCode),
		nil,
	)

	response, err := manager.cli.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error deleting product: %s", response.StatusCode)
	}

	return nil
}

// UseClient will redefine an HTTP Client used to make API calls
func (manager announcementManager) UseClient(cli http.Client) {
	manager.cli = cli
}

// NewAnnouncement return an announcement manager to handle announcements on mercado's livre API
func NewAnnouncement(m Meli) AnnouncementManager {
	return &announcementManager{
		m: m,
		cli: http.Client{
			Timeout: time.Second * 15,
		},
	}
}
