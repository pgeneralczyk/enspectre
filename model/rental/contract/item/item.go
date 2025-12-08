package item

import "time"

type RentalContractItemResponse struct {
	RentalContractItem RentalContractItem `json:"rentalContractItem"`
}

type RentalContractItem struct {
	RentalContractItemId             int64                     `json:"rentalContractItemId"`
	LineNumber                       int64                     `json:"lineNumber"`
	CreatedAt                        time.Time                 `json:"createdAt"`
	RentalContractVersion            RentalContractVersion     `json:"rentalContractVersion"`
	RentalContractQuote              RentalContractQuote       `json:"rentalContractQuote"`
	IsActive                         bool                      `json:"isActive"`
	Status                           string                    `json:"status"`
	Type                             int                       `json:"type"`
	Note                             string                    `json:"note"`
	InternalNote                     string                    `json:"internalNote"`
	CustomersReferenceNumber         string                    `json:"customersReferenceNumber"`
	InitialRentalContractItem        InitialRentalContractItem `json:"initialRentalContractItem"`
	VATPercent                       float64                   `json:"vATPercent"`
	RateChangeType                   interface{}               `json:"rateChangeType"`
	Customer                         Customer                  `json:"customer"`
	Product                          Product                   `json:"product"`
	ProductDescription               string                    `json:"productDescription"`
	Quantity                         int64                     `json:"quantity"`
	RemainingQuantity                int64                     `json:"remainingQuantity"`
	RequestedDeliveryDate            string                    `json:"requestedDeliveryDate"`
	ConfirmedDeliveryDate            *string                   `json:"confirmedDeliveryDate"` // nullable
	DiscountPercent                  *float64                  `json:"discountPercent"`       // nullable
	RecurringRate                    RecurringRate             `json:"recurringRate"`
	CurrentRateStep                  CurrentRateStep           `json:"currentRateStep"`
	Period                           Period                    `json:"period"`
	OnRentDays                       int64                     `json:"onRentDays"`
	AverageRatePerDay                float64                   `json:"averageRatePerDay"`
	CheckOut                         Check                     `json:"checkOut"`
	CheckIn                          Check                     `json:"checkIn"`
	ShippingAddress                  Address                   `json:"shippingAddress"`
	ShippingAddressDeliveryTerms     string                    `json:"shippingAddressDeliveryTerms"`
	ShippingAddressDeliveryTermPlace string                    `json:"shippingAddressDeliveryTermPlace"`
	ShippingService                  ShippingService           `json:"shippingService"`
	Totals                           Totals                    `json:"totals"`
	TotalsConverted                  Totals                    `json:"totalsConverted"`
	Seller                           Seller                    `json:"seller"`
	Accounting                       Accounting                `json:"accounting"`
	CalendarSummary                  CalendarSummary           `json:"calendarSummary"`
}

type RentalContractVersion struct {
	RentalContractVersionId   int64     `json:"rentalContractVersionId"`
	Version                   int       `json:"version"`
	Revision                  int       `json:"revision"`
	CreatedBy                 CreatedBy `json:"createdBy"`
	RentalContractVersionLink string    `json:"rentalContractVersionLink"`
}

type CreatedBy struct {
	UserId    *int64  `json:"userId"`
	Name      string  `json:"name"`
	FirstName string  `json:"firstName"`
	UserLink  *string `json:"userLink"`
}

type RentalContractQuote struct {
	RentalContractQuoteId *int64 `json:"rentalContractQuoteId"`
	Name                  string `json:"name"`
}

type InitialRentalContractItem struct {
	RentalContractItemId int64 `json:"rentalContractItemId"`
}

type Customer struct {
	CustomerId   int64  `json:"customerId"`
	Name         string `json:"name"`
	CustomerLink string `json:"customerLink"`
}

type Product struct {
	ProductId   int64  `json:"productId"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	ProductLink string `json:"productLink"`
}

type RecurringRate struct {
	RecurringRateId   int64  `json:"recurringRateId"`
	Status            string `json:"status"`
	Currency          string `json:"currency"`
	RateInterval      int    `json:"rateInterval"`
	Type              int    `json:"type"`
	NumberOfRateSteps int    `json:"numberOfRateSteps"`
	RecurringRateLink string `json:"recurringRateLink"`
}

type CurrentRateStep struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}

type Period struct {
	EffectiveDate   string  `json:"effectiveDate"`
	ExpirationDate  string  `json:"expirationDate"`
	TerminationDate *string `json:"terminationDate"`
}

type Check struct {
	ExpectedDate   string   `json:"expectedDate"`
	CheckOutStatus *int     `json:"checkOutStatus,omitempty"`
	Location       Location `json:"location"`
}

type Location struct {
	LocationId   int64   `json:"locationId"`
	ShortName    string  `json:"shortName"`
	Address      Address `json:"address"`
	LocationLink string  `json:"locationLink"`
}

type Address struct {
	AddressId    *int64 `json:"addressId,omitempty"`
	Name         string `json:"name"`
	FirstName    string `json:"firstname,omitempty"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	PostalCode   string `json:"postalCode"`
	City         string `json:"city"`
	Region       string `json:"region"`
	CountryCode  string `json:"countryCode"`
	Country      string `json:"country,omitempty"`
	AddressLink  string `json:"addressLink,omitempty"`
}

type ShippingService struct {
	ShippingServiceId   *int64  `json:"shippingServiceId"`
	Name                string  `json:"name"`
	Carrier             Carrier `json:"carrier"`
	ShippingServiceLink *string `json:"shippingServiceLink"`
}

type Carrier struct {
	CarrierId   *int64  `json:"carrierId"`
	Name        string  `json:"name"`
	CarrierLink *string `json:"carrierLink"`
}

type Totals struct {
	Currency          string  `json:"currency"`
	TotalAmount       float64 `json:"totalAmount"`
	DiscountAmount    float64 `json:"discountAmount"`
	TotalNetAmount    float64 `json:"totalNetAmount"`
	AccruedRevenue    float64 `json:"accruedRevenue"`
	InvoicedAmount    float64 `json:"invoicedAmount"`
	ForecastedRevenue float64 `json:"forecastedRevenue"`
	EarnedRevenue     float64 `json:"earnedRevenue"`
	DeferredRevenue   float64 `json:"deferredRevenue"`
}

type Seller struct {
	EmployeeId   int64  `json:"employeeId"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmployeeLink string `json:"employeeLink"`
}

type Accounting struct {
	VATCodeDefinition    VATCodeDefinition    `json:"vATCodeDefinition"`
	GeneralLedgerAccount GeneralLedgerAccount `json:"generalLedgerAccount"`
	Department           Department           `json:"department"`
	FinanceProject       FinanceProject       `json:"financeProject"`
	Asset                Asset                `json:"asset"`
}

type VATCodeDefinition struct {
	VATCode    string  `json:"vATCode"`
	VATPercent float64 `json:"vATPercent"`
}

type GeneralLedgerAccount struct {
	GeneralLedgerAccountId   int64  `json:"generalLedgerAccountId"`
	GeneralLedgerAccountLink string `json:"generalLedgerAccountLink"`
}

type Department struct {
	DepartmentId   *int64  `json:"departmentId"`
	DepartmentLink *string `json:"departmentLink"`
}

type FinanceProject struct {
	FinanceProjectId   *int64  `json:"financeProjectId"`
	FinanceProjectLink *string `json:"financeProjectLink"`
}

type Asset struct {
	AssetId   *int64  `json:"assetId"`
	AssetLink *string `json:"assetLink"`
}

type CalendarSummary struct {
	Quantity             Quantity             `json:"quantity"`
	RecurringRateSummary RecurringRateSummary `json:"recurringRate"`
	DiscountPercentage   DiscountPercentage   `json:"discountPercentage"`
	LastInvoicedDate     string               `json:"lastInvoicedDate"`
	InvoicedDays         int                  `json:"invoicedDays"`
}

type Quantity struct {
	Min             int `json:"min"`
	Max             int `json:"max"`
	Current         int `json:"current"`
	NumberOfChanges int `json:"numberOfChanges"`
}

type RecurringRateSummary struct {
	NumberOfChanges      int                  `json:"numberOfChanges"`
	CurrentRecurringRate CurrentRecurringRate `json:"currentRecurringRate"`
}

type CurrentRecurringRate struct {
	RecurringRateId   int64  `json:"recurringRateId"`
	RecurringRateLink string `json:"recurringRateLink"`
}

type DiscountPercentage struct {
	Min             float64  `json:"min"`
	Max             float64  `json:"max"`
	Current         *float64 `json:"current"`
	NumberOfChanges int      `json:"numberOfChanges"`
}
