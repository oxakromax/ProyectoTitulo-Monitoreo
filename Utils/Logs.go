package Utils

import (
	"encoding/json"
	"github.com/oxakromax/Proyecto_Titulacion-Backend/SQL/Structs"
	"golang.org/x/exp/slices"
	"io"
	"time"
)

// Log es una estructura que representa un registro de log en el sistema.
// Los campos Level, WindowsIdentity, ProcessName, TimeStamp, Message, JobKey, RawMessage,
// RobotName, HostMachineName, MachineId y RuntimeType contienen información detallada sobre el log.
// El campo Id contiene un identificador único para el log.
type Log struct {
	Level           string      `json:"Level"`
	WindowsIdentity string      `json:"WindowsIdentity"`
	ProcessName     string      `json:"ProcessName"`
	TimeStamp       time.Time   `json:"TimeStamp"`
	Message         string      `json:"Message"`
	JobKey          string      `json:"JobKey"`
	RawMessage      string      `json:"RawMessage"`
	RobotName       string      `json:"RobotName"`
	HostMachineName string      `json:"HostMachineName"`
	MachineId       int         `json:"MachineId"`
	RuntimeType     interface{} `json:"RuntimeType"`
	Id              int         `json:"Id"`
}

// LogResponse es una estructura que contiene una lista de registros de log.
// Los campos OdataContext y OdataCount contienen información adicional sobre la respuesta.
// El campo Value contiene la lista de registros de log.
type LogResponse struct {
	OdataContext string `json:"@odata.context"`
	OdataCount   int    `json:"@odata.count"`
	Value        []Log  `json:"value"`
}

// Get es una función que permite obtener una lista de registros de log desde una API.
// La función hace una petición GET a la URL especificada en la constante UipathOrg.GetURL()
// y almacena la respuesta en una instancia de la estructura LogResponse.
func (r *LogResponse) Get(FolderID int) error {
	// Hacemos una petición GET a la URL especificada en UipathOrg.GetURL()
	resp, err := RequestAPI("GET", UipathOrg.GetURL()+"odata/RobotLogs", nil, FolderID)
	// Si hubo un error en la petición, lo devolvemos
	if err != nil {
		return err
	}

	// Leemos el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Convertimos el cuerpo de la respuesta a una instancia de la estructura LogResponse
	err = json.Unmarshal(body, r)
	if err != nil {
		return err
	}

	return nil
}

// FilterLogs es una función que permite filtrar una lista de registros de log.
// La función recibe una lista de procesos y devuelve una lista de registros de log que cumplen
// con dos condiciones: que la fecha y hora del log sea posterior a la fecha y hora especificada
// en la variable LastMonitoredTime, y que el nombre del proceso que generó el log esté incluido
// en la lista de procesos recibida como argumento.
func (r *LogResponse) FilterLogs(Processes []Structs.ProcesosBDD) []Log {
	// Inicializamos una lista vacía para almacenar los registros de log filtrados
	var filteredLogs []Log

	// Iteramos sobre cada registro de log en la lista Value de la estructura LogResponse
	for _, log := range r.Value {
		// Comprobamos si el nombre del proceso que generó el log está incluido en la lista de procesos
		isInProcess := slices.ContainsFunc(Processes, func(p Structs.ProcesosBDD) bool {
			if p.Nombre == log.ProcessName {
				return true
			}
			return false
		})

		// Si el log cumple con las dos condiciones mencionadas, lo añadimos a la lista de registros filtrados
		if log.TimeStamp.After(LastMonitoredTime) && isInProcess {
			filteredLogs = append(filteredLogs, log)
		}
	}
	// Devolvemos la lista de registros de log filtrados
	return filteredLogs
}
