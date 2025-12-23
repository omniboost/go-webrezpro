package webrezpro

type TrialBalanceResponse struct {
	Success      bool           `json:"success"`
	Date         Date           `json:"date"`
	Statistics   []Statistics   `json:"statistics"`
	TrialBalance []TrialBalance `json:"trial_balance"`
	MarketCodes  []MarketCodes  `json:"marketCodes"`
}
type AccountRef struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}
type RoomTotalDetail struct {
	AccountRef AccountRef `json:"AccountRef"`
}
type Statistics struct {
	RoomTotalDetail RoomTotalDetail `json:"RoomTotalDetail"`
}
type JournalEntryDetailAccountRef struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}
type JournalEntryDetail struct {
	JournalEntryDetailAccountRef JournalEntryDetailAccountRef `json:"AccountRef"`
	PostingType                  string                       `json:"PostingType"`
}
type TrialBalance struct {
	JournalEntryDetail JournalEntryDetail `json:"JournalEntryDetail"`
	Amount             StringFloat        `json:"Amount"`
}
type MarketCodeDetailAccountRef struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}
type MarketCodeDetail struct {
	MarketCodeDetailAccountRef MarketCodeDetailAccountRef `json:"AccountRef"`
}
type MarketCodes struct {
	MarketCodeDetail MarketCodeDetail `json:"MarketCodeDetail"`
	Amount           StringFloat      `json:"Amount"`
	Rooms            int              `json:"Rooms"`
}

type HotelDetails struct {
	PhoneNumbers   []map[string]string `json:"phone_numbers"`
	StateProv      string              `json:"state_prov"`
	EmailAddresses []string            `json:"email_addresses"`
	CheckOutTime   string              `json:"check_out_time"`
	PaymentTypes   []string            `json:"payment_types"`
	City           string              `json:"city"`
	PropertyName   string              `json:"property_name"`
	Taxes          []struct {
		ApplyToRoomCost int    `json:"apply_to_room_cost"`
		ApplyToRoomTax  int    `json:"apply_to_room_tax"`
		Name            string `json:"name"`
		Type            string `json:"type"`
		Rate            string `json:"rate"`
	} `json:"taxes"`
	Departments []struct {
		DepartmentID string `json:"department_id"`
		Name         string `json:"name"`
	} `json:"departments"`
	Country            string `json:"country"`
	Website            string `json:"website"`
	CheckInTime        string `json:"check_in_time"`
	ParkingInformation string `json:"parking_information"`
	ZIPPostalCode      string `json:"zip_postal_code"`
	Address            string `json:"address"`
}
