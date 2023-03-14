package announcements

import "time"

// Announcement represents response content returned by mercado livre's API
type Announcement struct {
	ID                string    `json:"id"`
	SiteID            string    `json:"site_id"`
	Title             string    `json:"title"`
	Subtitle          string    `json:"subtitle"`
	SellerID          string    `json:"seller_id"`
	CategoryID        string    `json:"category_id"`
	Price             float64   `json:"price"`
	BasePrice         float64   `json:"base_price"`
	OriginalPrice     float64   `json:"original_price"`
	CurrencyID        string    `json:"currency_id"`
	InitialQuantity   int       `json:"initial_quantity"`
	AvailableQuantity int       `json:"available_quantity"`
	SoldQuantity      int       `json:"sold_quantity"`
	BuyingMode        string    `json:"buying_mode"`
	ListingTypeID     string    `json:"listing_type_id"`
	StartTime         time.Time `json:"start_time"`
	StopTime          time.Time `json:"stop_time"`
	EndTime           time.Time `json:"end_time"`
	ExpirationTime    time.Time `json:"expiration_time"`
	Condition         string    `json:"condition"`
	Thumbnail         string    `json:"thumbnail"`
	SecureThumbnail   string    `json:"secure_thumbnail"`
	VideoID           string    `json:"video_id"`
	Warranty          string    `json:"warranty"`
	DateCreated       time.Time `json:"date_created"`
	LastUpdated       time.Time `json:"last_updated"`
	Permalink         string    `json:"permalink"`
}

type AttributeValue struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AttributeCombination struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ValueID   string `json:"value_id"`
	ValueName string `json:"value_name"`
}

type Attribute struct {
	ID           string           `json:"id"`
	Name         string           `json:"name"`
	Hierarchy    string           `json:"hierarchy"`
	ValueType    string           `json:"value_type"`
	ValueName    string           `json:"value_name"`
	ValueID      string           `json:"value_id"`
	Tags         []string         `json:"tags"`
	Values       []AttributeValue `json:"values"`
	AllowedUnits []string         `json:"allowed_units"`
}

type Variation struct {
	ID                    string                 `json:"id"`
	AttributeCombinations []AttributeCombination `json:"attribute_combinations"`
	Price                 float64                `json:"price"`
	AvailableQuantity     int                    `json:"available_quantity"`
	Attributes            []Attribute            `json:"attributes"`
	SoldQuantity          int                    `json:"sold_quantity"`
	PictureIDS            []string               `json:"picture_ids"`
}

type Location struct {
	AddressLine string `json:"address_line"`
	ZipCode     string `json:"zip_code"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}

type SellerAddress struct {
	ID          string `json:"id"`
	Comment     string `json:"comment"`
	AddressLine string `json:"address_line"`
	ZipCode     string `json:"zip_code"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}

type Shipping struct {
	Mode         string   `json:"mode"`
	LocalPickUp  bool     `json:"local_pick_up"`
	FreeShipping bool     `json:"free_shipping"`
	Dimensions   string   `json:"dimensions"`
	Tags         []string `json:"tags"`
}

type Picture struct {
	Source string `json:"source"`
}

type AnnouncementRequest struct {
	ID                string        `json:"id"`
	Title             string        `json:"title"`
	CategoryID        string        `json:"category_id"`
	Price             float64       `json:"price"`
	CurrencyID        string        `json:"currency_id"`
	AvailableQuantity int           `json:"available_quantity"`
	BuyingMode        string        `json:"buying_mode"`
	ListingTypeID     string        `json:"listing_type_id"`
	Condition         string        `json:"condition"`
	Description       string        `json:"description"`
	VideoID           string        `json:"video_id"`
	Tags              []string      `json:"tags"`
	Warranty          string        `json:"warranty"`
	Status            status        `json:"status"`
	Shipping          Shipping      `json:"shipping"`
	SellerAddress     SellerAddress `json:"seller_address"`
	Location          Location      `json:"location"`
	Pictures          []Picture     `json:"pictures"`
	Attributes        []Attribute   `json:"attributes"`
	Variations        []Variation   `json:"variations"`
}

func (i *AnnouncementRequest) SetID(id string) *AnnouncementRequest {
	i.ID = id

	return i
}

func (i *AnnouncementRequest) SetTitle(title string) *AnnouncementRequest {
	i.Title = title

	return i
}

func (i *AnnouncementRequest) SetCategoryID(id string) *AnnouncementRequest {
	i.CategoryID = id

	return i
}

func (i *AnnouncementRequest) SetPrice(price float64) *AnnouncementRequest {
	i.Price = price

	return i
}

func (i *AnnouncementRequest) SetCurrencyID(id string) *AnnouncementRequest {
	i.CurrencyID = id

	return i
}

func (i *AnnouncementRequest) SetAvailableQuantity(quantity int) *AnnouncementRequest {
	i.AvailableQuantity = quantity

	return i
}

func (i *AnnouncementRequest) SetBuyingMode(mode string) *AnnouncementRequest {
	i.BuyingMode = mode

	return i
}

func (i *AnnouncementRequest) SetListingTypeID(id string) *AnnouncementRequest {
	i.ListingTypeID = id

	return i
}

func (i *AnnouncementRequest) SetCondition(condition string) *AnnouncementRequest {
	i.Condition = condition

	return i
}

func (i *AnnouncementRequest) SetDescription(description string) *AnnouncementRequest {
	i.Description = description

	return i
}

func (i *AnnouncementRequest) SetVideoID(id string) *AnnouncementRequest {
	i.VideoID = id

	return i
}

func (i *AnnouncementRequest) SetTags(tags []string) *AnnouncementRequest {
	i.Tags = tags

	return i
}

func (i *AnnouncementRequest) SetWarranty(warranty string) *AnnouncementRequest {
	i.Warranty = warranty

	return i
}

func (i *AnnouncementRequest) SetStatus(s status) *AnnouncementRequest {
	i.Status = s

	return i
}

func (i *AnnouncementRequest) SetShipping(shipping Shipping) *AnnouncementRequest {
	i.Shipping = shipping

	return i
}

func (i *AnnouncementRequest) SetSellerAddress(address SellerAddress) *AnnouncementRequest {
	i.SellerAddress = address

	return i
}

func (i *AnnouncementRequest) SetLocation(location Location) *AnnouncementRequest {
	i.Location = location

	return i
}

func (i *AnnouncementRequest) SetPictures(pictures []Picture) *AnnouncementRequest {
	i.Pictures = pictures

	return i
}

func (i *AnnouncementRequest) AddPicture(picture Picture) *AnnouncementRequest {
	i.Pictures = append(i.Pictures, picture)

	return i
}

func (i *AnnouncementRequest) SetAttributes(attributes []Attribute) *AnnouncementRequest {
	i.Attributes = attributes

	return i
}

func (i *AnnouncementRequest) AddAttribute(attribute Attribute) *AnnouncementRequest {
	i.Attributes = append(i.Attributes, attribute)

	return i
}

func (i *AnnouncementRequest) SetVariations(variations []Variation) *AnnouncementRequest {
	i.Variations = variations

	return i
}

func (i *AnnouncementRequest) AddVariation(variation Variation) *AnnouncementRequest {
	i.Variations = append(i.Variations, variation)

	return i
}

// NewAnnouncement return a new Announcement with empty attributes
func NewAnnouncementRequest() *AnnouncementRequest {
	return &AnnouncementRequest{}
}
