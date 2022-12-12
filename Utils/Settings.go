package Utils

var (
	QueryAuth = QueryURLAuth{
		ClientId:     "",
		ClientSecret: "",
		GrantType:    "client_credentials",
		Scope:        "",
	}
	UiPathToken *UipathToken
	UipathOrg   = UipathORG{
		OrganizationName: "studentfinis",
		TenantName:       "DefaultTenant",
		FolderID:         "3902201",
	}
)
