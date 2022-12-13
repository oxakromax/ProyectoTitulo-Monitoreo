package Utils

// UipathToken es una estructura que representa la respuesta de una petición de autenticación en la API de UiPath.
// La respuesta contiene un token de acceso, la fecha de expiración del token, el tipo de token y el alcance del mismo.
type UipathToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

// QueryURLAuth es una estructura que representa los parámetros necesarios para realizar una petición de autenticación en la API de UiPath.
// Estos parámetros son el tipo de autorización, el identificador y secreto del cliente, y el alcance del token.
type QueryURLAuth struct {
	GrantType    string `json:"grant_type" url:"grant_type"`
	ClientId     string `json:"client_id" url:"client_id"`
	ClientSecret string `json:"client_secret" url:"client_secret"`
	Scope        string `json:"scope" url:"scope"`
}

// UipathORG es una estructura que representa la información de la organización de UiPath en la que se realizarán las peticiones a la API.
// La estructura contiene el nombre de la organización, el nombre del inquilino y el identificador de la carpeta.
type UipathORG struct {
	OrganizationName string `json:"org"`
	TenantName       string `json:"tenant"`
	Id               string `json:"id"`
}

// GetURL es una función que devuelve la URL base para hacer peticiones a la API de Orchestrator de UiPath.
// La función utiliza la información almacenada en una instancia de la estructura UipathORG para construir la URL.
func (u *UipathORG) GetURL() string {
	return "https://cloud.uipath.com/" + u.OrganizationName + "/" + u.TenantName + "/orchestrator_/"
}
