package Utils

import (
	"io"
	"net/http"
	"time"
)

func RequestAPI(method, url string, body io.Reader) (*http.Response, error) {
	for IsRefreshingToken.Get() {
		time.Sleep(1 * time.Second)
	}
	client := new(http.Client)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+UiPathToken.AccessToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-UIPATH-OrganizationUnitId", UipathOrg.FolderID)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
