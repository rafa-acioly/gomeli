package responses

import "time"

// Item represents the content of a response when the announcement is created
type Item struct {
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
