package invoiceperiod

import "time"

type InvoicePeriod struct {
	InvoicePeriodID int       `json:"invoicePeriodId"`
	CreatedAt       time.Time `json:"createdAt"`
	Status          string    `json:"status"`
	Period          struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	} `json:"period"`
	PlannedInvoiceDate       string `json:"plannedInvoiceDate"`
	CustomersReferenceNumber string `json:"customersReferenceNumber"`
	TotalDays                int    `json:"totalDays"`
	TotalOffRentDays         int    `json:"totalOffRentDays"`
	TotalBillingEntries      int    `json:"totalBillingEntries"`
	LatestBillingEntries     []struct {
		BillingEntryID int    `json:"billingEntryId"`
		Status         string `json:"status"`
		Period         struct {
			StartDate string `json:"startDate"`
			EndDate   string `json:"endDate"`
		} `json:"period"`
		Totals struct {
			TotalNetAmount float64 `json:"totalNetAmount"`
			InvoicedAmount float64 `json:"invoicedAmount"`
		} `json:"totals"`
	} `json:"latestBillingEntries"`
	TotalSalesInvoices  int `json:"totalSalesInvoices"`
	LatestSalesInvoices []struct {
		SalesInvoiceID   int       `json:"salesInvoiceId"`
		CreatedAt        time.Time `json:"createdAt"`
		IssueDate        string    `json:"issueDate"`
		Status           string    `json:"status"`
		SalesInvoiceLink string    `json:"salesInvoiceLink"`
	} `json:"latestSalesInvoices"`
	TotalSalesCreditNotes  int `json:"totalSalesCreditNotes"`
	LatestSalesCreditNotes []struct {
		SalesCreditNoteID   int       `json:"salesCreditNoteId"`
		CreatedAt           time.Time `json:"createdAt"`
		IssueDate           string    `json:"issueDate"`
		Status              string    `json:"status"`
		SalesCreditNoteLink string    `json:"salesCreditNoteLink"`
	} `json:"latestSalesCreditNotes"`
	Totals struct {
		Currency             string  `json:"currency"`
		TotalAmount          float64 `json:"totalAmount"`
		DiscountAmount       float64 `json:"discountAmount"`
		TotalNetAmount       float64 `json:"totalNetAmount"`
		InvoicedAmount       float64 `json:"invoicedAmount"`
		InvoiceableAmount    float64 `json:"invoiceableAmount"`
		ProformaAmount       float64 `json:"proformaAmount"`
		PendingPaymentAmount float64 `json:"pendingPaymentAmount"`
		PaidAmount           float64 `json:"paidAmount"`
		CreditedAmount       float64 `json:"creditedAmount"`
	} `json:"totals"`
	TotalsConverted struct {
		Currency          string  `json:"currency"`
		TotalAmount       float64 `json:"totalAmount"`
		DiscountAmount    float64 `json:"discountAmount"`
		TotalNetAmount    float64 `json:"totalNetAmount"`
		InvoicedAmount    float64 `json:"invoicedAmount"`
		InvoiceableAmount float64 `json:"invoiceableAmount"`
	} `json:"totalsConverted"`
}

type InvoicePeriodResponse struct {
	InvoicePeriod InvoicePeriod `json:"invoicePeriod"`
}
