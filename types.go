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
