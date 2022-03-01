package announcements

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

type Item struct {
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

func (i *Item) SetID(id string) *Item {
	i.ID = id

	return i
}

func (i *Item) SetTitle(title string) *Item {
	i.Title = title

	return i
}

func (i *Item) SetCategoryID(id string) *Item {
	i.CategoryID = id

	return i
}

func (i *Item) SetPrice(price float64) *Item {
	i.Price = price

	return i
}

func (i *Item) SetCurrencyID(id string) *Item {
	i.CurrencyID = id

	return i
}

func (i *Item) SetAvailableQuantity(quantity int) *Item {
	i.AvailableQuantity = quantity

	return i
}

func (i *Item) SetBuyingMode(mode string) *Item {
	i.BuyingMode = mode

	return i
}

func (i *Item) SetListingTypeID(id string) *Item {
	i.ListingTypeID = id

	return i
}

func (i *Item) SetCondition(condition string) *Item {
	i.Condition = condition

	return i
}

func (i *Item) SetDescription(description string) *Item {
	i.Description = description

	return i
}

func (i *Item) SetVideoID(id string) *Item {
	i.VideoID = id

	return i
}

func (i *Item) SetTags(tags []string) *Item {
	i.Tags = tags

	return i
}

func (i *Item) SetWarranty(warranty string) *Item {
	i.Warranty = warranty

	return i
}

func (i *Item) SetStatus(s status) *Item {
	i.Status = s

	return i
}

func (i *Item) SetShipping(shipping Shipping) *Item {
	i.Shipping = shipping

	return i
}

func (i *Item) SetSellerAddress(address SellerAddress) *Item {
	i.SellerAddress = address

	return i
}

func (i *Item) SetLocation(location Location) *Item {
	i.Location = location

	return i
}

func (i *Item) SetPictures(pictures []Picture) *Item {
	i.Pictures = pictures

	return i
}

func (i *Item) AddPicture(picture Picture) *Item {
	i.Pictures = append(i.Pictures, picture)

	return i
}

func (i *Item) SetAttributes(attributes []Attribute) *Item {
	i.Attributes = attributes

	return i
}

func (i *Item) AddAttribute(attribute Attribute) *Item {
	i.Attributes = append(i.Attributes, attribute)

	return i
}

func (i *Item) SetVariations(variations []Variation) *Item {
	i.Variations = variations

	return i
}

func (i *Item) AddVariation(variation Variation) *Item {
	i.Variations = append(i.Variations, variation)

	return i
}

func NewItem() *Item {
	return &Item{}
}
