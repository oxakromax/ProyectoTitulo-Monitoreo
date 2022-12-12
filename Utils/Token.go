package Utils

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"strings"
)

func RefreshUiPathToken() error {
	IsRefreshingToken.Set(true)
	defer IsRefreshingToken.Set(false)
	url := "https://cloud.uipath.com/identity_/connect/token"
	method := "POST"
	vals, err := query.Values(QueryAuth)
	if err != nil {
		return err
	}
	payload := strings.NewReader(vals.Encode())
	client := new(http.Client)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &UiPathToken) // Refresh the token
	if err != nil {
		return err
	}
	return nil
}
