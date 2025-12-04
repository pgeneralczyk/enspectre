package apiclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const bearerToken = "_AOARALZrk2e--2bxswbuw2"
const baseUrl = "https://api.rambase.net/"
const db = "TEM-NO"

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func setReqQueryParams(fullUrl string) string {
	u, _ := url.Parse(fullUrl)
	q := u.Query()          // pobiera istniejące query parametry
	q.Set("$db", db)        // dodajemy / nadpisujemy db
	u.RawQuery = q.Encode() // zapisujemy z powrotem
	return u.String()
}

func concatUrl(endpoint string) string {
	base := baseUrl
	if base[len(base)-1] == '/' {
		base = base[:len(base)-1]
	}
	if endpoint[0] != '/' {
		endpoint = "/" + endpoint
	}
	return setReqQueryParams(base + endpoint)
}

func Get(ctx context.Context, endpoint string) ([]byte, error) {

	url := concatUrl(endpoint)
	fmt.Println(url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Bearer Authorization
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
