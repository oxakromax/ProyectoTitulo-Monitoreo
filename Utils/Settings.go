package Utils

import "github.com/oxakromax/Proyecto_Titulacion-Backend/SQL/Structs"

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
	}
	ProcessBDDList = new(Structs.ProcessBDDArray)
)
