package Utils

import (
	"encoding/json"
	"io"
)

// ProcessResponse es una estructura que representa la respuesta de una petición a la API de UiPath para obtener información de los procesos.
// La respuesta contiene el contexto OData, el número de elementos incluidos en la respuesta, y un arreglo con la información de cada proceso.
type ProcessResponse struct {
	OdataContext string    `json:"@odata.context"`
	OdataCount   int       `json:"@odata.count"`
	Value        []Process `json:"value"`
}

// Process es una estructura que representa información de un proceso en UiPath.
// La estructura contiene el identificador, la versión, el nombre, la descripción, entre otros datos del proceso.
type Process struct {
	Key              string      `json:"Key"`
	ProcessKey       string      `json:"ProcessKey"`
	ProcessVersion   string      `json:"ProcessVersion"`
	IsLatestVersion  bool        `json:"IsLatestVersion"`
	IsProcessDeleted bool        `json:"IsProcessDeleted"`
	Description      interface{} `json:"Description"`
	Name             string      `json:"Name"`
	EnvironmentId    int         `json:"EnvironmentId"`
	EnvironmentName  string      `json:"EnvironmentName"`
	InputArguments   interface{} `json:"InputArguments"`
	Id               int         `json:"Id"`
	Arguments        interface{} `json:"Arguments"`
}

// Get es una función que permite realizar una petición a la API de Orchestrator de UiPath para obtener la información de todos los procesos definidos en una organización.
// La función recibe como argumento un puntero a una estructura de tipo ProcessResponse, en la que se almacenará la información de los procesos obtenida de la petición.
// La función retorna un error en caso de que ocurra algún problema al realizar la petición o al procesar la respuesta.

func (r *ProcessResponse) Get() error {
	resp, err := RequestAPI("GET", UipathOrg.GetURL()+"odata/Releases", nil)
	// print the json
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, r)
	if err != nil {
		return err
	}
	return nil
}

// GetProcesses es una función que permite obtener la información de todos los procesos incluidos en una estructura de tipo ProcessResponse.
// La función recibe como argumento un puntero a una estructura de tipo ProcessResponse, en la que se almacena la información de los procesos obtenida de la petición a la API de UiPath.
// La función retorna un arreglo con la información de cada proceso incluido en la estructura ProcessResponse.
func (r *ProcessResponse) GetProcesses() []Process {
	return r.Value
}
