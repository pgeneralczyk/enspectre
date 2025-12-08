package itemApi

import (
	"context"
	"encoding/json"
	"enspectre/apiclient"
	"enspectre/model/rental/contract/item"
	"fmt"
)

func GetRentalContractItem(ctx context.Context, rentalContractId int64, itemId int64) (*item.RentalContractItem, error) {

	var resp item.RentalContractItemResponse
	var err error

	path := fmt.Sprintf("rental/contracts/%d/items/%d", rentalContractId, itemId)

	var data []byte
	data, err = apiclient.Get(ctx, path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.RentalContractItem, nil

}
