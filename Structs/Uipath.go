package Structs

type UipathToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type QueryURLAuth struct {
	GrantType    string `json:"grant_type" url:"grant_type"`
	ClientId     string `json:"client_id" url:"client_id"`
	ClientSecret string `json:"client_secret" url:"client_secret"`
	Scope        string `json:"scope" url:"scope"`
}

type UipathORG struct {
	OrganizationName string `json:"org"`
	TenantName       string `json:"tenant"`
	FolderID         string `json:"folderId"`
}

func (u *UipathORG) GetURL() string {
	return "https://cloud.uipath.com/" + u.OrganizationName + "/" + u.TenantName + "/orchestrator_/"
}
