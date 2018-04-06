package search_and_availability

import (
	"azri_hub/dcache"
)

type RequestPayload struct {
	Agent                     dcache.Agent
	UserAgent                 string
	IP                        string
	SoftCutoffTime            int
	KillCutoffTime            int
	MoreResults               bool
	SupplierCityCode          string
	SupplierDestinationCode   string
	HotelInfo                 map[string]string
	Format                    string
	ItemCode                  string
	HotelCodes                []string
	HotelIDs                  []string
	SearchID                  string
	Markup                    map[string]string
	Supplier                  string
	Bundle                    string
	RequestID                 string
	SearchType                string
	ItemType                  string
	DestinationType           string
	LocationName              string
	AvailableOnly             bool
	ExactDestination          string
	Destination               string
	Country                   string
	ImmediateConfirmationOnly string
	IncludeRecommendedOnly    string
	IataAirportCode           string
	ItemGroupcode             string
	FacilityCodes             []string
	OrderBy                   string
	NoOfResults               int
	PointOfInterest           string
	Language                  string
	Shiftdays                 string
	MinPrice                  float64
	MaxPrice                  float64
	MaxRatesPerRoom           int
	PaymentType               string
	Suppliers                 []string
	DestinationCode           string `json:"destination_code"`
	Checkin                   string `json:"checkin"`
	Checkout                  string `json:"checkout"`
	ClientNationality         string `json:"client_nationality"`
	HotelName                 string `json:"hotel_name"`
	HotelCategory             string `json:"hotel_category"`
	Rates                     string `json:"rates"`
	Response                  string `json:"response"`
	Currency                  string `json:"currency"`
	Rooms                     []Room `json:"rooms"`
}

type Room struct {
	Adults        int   `json:"adults"`
	Children      []int `json:"children"`
	NoOfExtraBeds int   `json:"no_of_extra_beds"`
	NoOfCots      int   `json:"no_of_cots"`
}
