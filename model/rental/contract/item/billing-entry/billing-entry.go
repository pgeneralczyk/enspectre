package billingentry

type BillingEntry struct {
	BillingEntryID     int `json:"billingEntryId"`
	RentalContractItem struct {
		RentalContractItemID int `json:"rentalContractItemId"`
		Status               int `json:"status"`
		LineNumber           int `json:"lineNumber"`
		RentalContract       struct {
			RentalContractID   int    `json:"rentalContractId"`
			RentalContractLink string `json:"rentalContractLink"`
		} `json:"rentalContract"`
	} `json:"rentalContractItem"`
	Status string `json:"status"`
	Period struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	} `json:"period"`
	PlannedInvoiceDate string `json:"plannedInvoiceDate"`
	Quantity           int    `json:"quantity"`
	Product            struct {
		ProductID   int    `json:"productId"`
		Name        string `json:"name"`
		Type        string `json:"type"`
		Description string `json:"description"`
		ProductLink string `json:"productLink"`
	} `json:"product"`
	RecurringRate struct {
		RecurringRateID   int    `json:"recurringRateId"`
		Status            string `json:"status"`
		RateInterval      int    `json:"rateInterval"`
		Type              int    `json:"type"`
		RecurringRateLink string `json:"recurringRateLink"`
	} `json:"recurringRate"`
	CurrentRateStep struct {
		Currency string  `json:"currency"`
		Rate     float64 `json:"rate"`
	} `json:"currentRateStep"`
	TotalDays         int     `json:"totalDays"`
	TotalOffRentDays  int     `json:"totalOffRentDays"`
	AverageRatePerDay float64 `json:"averageRatePerDay"`
	RentalCalendar    struct {
		RentalCalendarID   int         `json:"rentalCalendarId"`
		RentalCalendarLink interface{} `json:"rentalCalendarLink"`
	} `json:"rentalCalendar"`
	Totals struct {
		Currency       string      `json:"currency"`
		TotalAmount    float64     `json:"totalAmount"`
		DiscountAmount float64     `json:"discountAmount"`
		TotalNetAmount float64     `json:"totalNetAmount"`
		ProformaAmount interface{} `json:"proformaAmount"`
		InvoicedAmount float64     `json:"invoicedAmount"`
	} `json:"totals"`
	TotalsConverted struct {
		Currency       string  `json:"currency"`
		TotalAmount    float64 `json:"totalAmount"`
		DiscountAmount float64 `json:"discountAmount"`
		TotalNetAmount float64 `json:"totalNetAmount"`
		InvoicedAmount float64 `json:"invoicedAmount"`
	} `json:"totalsConverted"`
}
