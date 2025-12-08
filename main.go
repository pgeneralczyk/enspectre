package main

import (
	"context"
	"enspectre/service/rental/contract/itemApi"
	"fmt"

	"github.com/fatih/color"
)

type RequestBody struct {
	Parameters struct {
		RentalContractItem struct {
			Quantity int `json:"Quantity"`
			Period   struct {
				EffectiveDate  string `json:"EffectiveDate"`
				ExpirationDate string `json:"ExpirationDate"`
			} `json:"Period"`
			RecurringRate struct {
				RecurringRateId int `json:"RecurringRateId"`
			} `json:"RecurringRate"`
			Product struct {
				ProductId int `json:"ProductId"`
			} `json:"Product"`
		} `json:"RentalContractItem"`
		RegisterAfterCreation bool `json:"RegisterAfterCreation"`
	} `json:"Parameters"`
}

type Product struct {
	ProductId       int
	RecurringRateId int
}

// func sendRequest(body *RequestBody) string {
// 	// AOP/101969 - Create rental contract subscription item
// 	url := "https://api.rambase.net/rental/contracts/101046/items/api-operations/101969/instances?$db=TEM-NO"
// 	token := "XDaOlxajUkGAwqtiUkB6Bg2"

// 	jsonBytes, err := json.Marshal(body)
// 	if err != nil {
// 		fmt.Println("JSON error:", err)
// 		return ""
// 	}

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
// 	if err != nil {
// 		fmt.Println("Request error:", err)
// 		return ""
// 	}

// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+token)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("HTTP error:", err)
// 		return ""
// 	}
// 	defer resp.Body.Close()

// 	respBody, _ := io.ReadAll(resp.Body)
// 	return string(respBody)
// }

// func main() {

// 	// artQuorsa := Product{
// 	// 	ProductId:       121941,
// 	// 	RecurringRateId: 102534,
// 	// }

// 	// artEasyGo := Product{
// 	// 	ProductId:       125330,
// 	// 	RecurringRateId: 102485,
// 	// }

// 	// artVanguard := Product{
// 	// 	ProductId:       121691,
// 	// 	RecurringRateId: 102479,
// 	// }

// 	artPrimeLicense := Product{
// 		ProductId:       121628,
// 		RecurringRateId: 102477,
// 	}

// 	body := RequestBody{}
// 	body.Parameters.RentalContractItem.Quantity = 2
// 	body.Parameters.RentalContractItem.RecurringRate.RecurringRateId = artPrimeLicense.RecurringRateId
// 	body.Parameters.RentalContractItem.Product.ProductId = artPrimeLicense.ProductId
// 	body.Parameters.RegisterAfterCreation = true

// 	nums := []int{2, 3, 4, 5, 6, 7, 8}

// 	for _, n := range nums {
// 		fromDate := time.Date(2025, time.Month(n), 1, 0, 0, 0, 0, time.UTC)
// 		toDate := fromDate.AddDate(0, 0, 9)

// 		body.Parameters.RentalContractItem.Period.EffectiveDate = fromDate.Format("2006-01-02")
// 		body.Parameters.RentalContractItem.Period.ExpirationDate = toDate.Format("2006-01-02")

// 		response := sendRequest(&body)
// 		if response == "" || response == "Invalid $access_token" {
// 			fmt.Println("Response:", response)
// 			break
// 		}
// 		fmt.Println("Response:", response)
// 	}

// }

func main() {
	ctx := context.Background()

	item, err := itemApi.GetRentalContractItem(ctx, 100490, 45)
	if err != nil {
		println(color.RedString(err.Error()))
		panic(err)
	}

	fmt.Println("Item id:", item.RentalContractItemId)
	fmt.Println("Line Number:", item.LineNumber)
	fmt.Println("Initial item id:", item.InitialRentalContractItem.RentalContractItemId)
	rarQuery := fmt.Sprintf("DOCNO==100490 && INITIALITM1REF==%d", item.InitialRentalContractItem.RentalContractItemId)
	println((color.GreenString(rarQuery)))

	rcaQuery := fmt.Sprintf("RCTNO==100490 && INITIALITM1REF==%d", item.InitialRentalContractItem.RentalContractItemId)
	println((color.YellowString(rcaQuery)))

}
