package main

import (
	"ProyectoTItulo/Structs"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"strings"
	"time"
)

var (
	QueryAuth = Structs.QueryURLAuth{
		ClientId:     "c68d9f9d-abe4-4f84-8178-4267ad6fe447",
		ClientSecret: "PyRahZliqlAc3)Q(",
		GrantType:    "client_credentials",
		Scope:        "OR.Webhooks OR.Monitoring OR.Monitoring OR.ML OR.Tasks OR.Analytics OR.Folders OR.BackgroundTasks OR.TestSets OR.TestSetExecutions OR.TestSetSchedules OR.TestDataQueues OR.Audit OR.License OR.Settings OR.Robots OR.Machines OR.Execution OR.Assets OR.Administration OR.Users OR.Jobs OR.Queues OR.Hypervisor",
	}
	UiPathToken *Structs.UipathToken
	UipathORG   = Structs.UipathORG{OrganizationName: "studentfinis",
		TenantName: "DefaultTenant",
		FolderID:   "3321402"}
	LastMonitoredTime time.Time
	IsRefreshingToken bool
)

func RefreshUiPathToken() error {
	IsRefreshingToken = true
	defer func() {
		IsRefreshingToken = false
	}()
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

func RequestAPI(method, url string, body io.Reader) (*http.Response, error) {
	for IsRefreshingToken {
		time.Sleep(1 * time.Second)
	}
	client := new(http.Client)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+UiPathToken.AccessToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-UIPATH-OrganizationUnitId", UipathORG.FolderID)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func main() {
	LastMonitoredTime = time.Now()
	go func() {
		for {
			err := RefreshUiPathToken()
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(59 * time.Minute) // Refresh the token every 10 minutes
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		resp, err := RequestAPI("GET", UipathORG.GetURL()+"odata/RobotLogs", nil)
		// print the json
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		JsonResponse := new(Structs.LogResponse)
		err = json.Unmarshal(body, &JsonResponse)
		if err != nil {
			fmt.Println(err)
		}
		byTime := JsonResponse.FilterLogsByTime(LastMonitoredTime)
		LastMonitoredTime = time.Now()
		fmt.Println(byTime)
	}()
	app := fiber.New()
	app.Get("/AuthReturn", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
