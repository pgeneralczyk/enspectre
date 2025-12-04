package item

type RentalContractItem struct {
	RentalContractItemId      int64  `json:"RentalContractItemId"`
	LineNumber                int64  `json:"LineNumber"`
	Status                    string `json:"Status"`
	Quantity                  int64  `json:"Quantity"`
	InitialRentalContractItem struct {
		RentalContractItemId int64 `json:"RentalContractItemId"`
	} `json:"InitialRentalContractItem"`
}
