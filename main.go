package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/oxakromax/ProyectoTitulo-Monitoreo/Utils"
	"github.com/oxakromax/Proyecto_Titulacion-Backend/SQL/Structs"
	"io"
	"net/http"
	"os"
	"sort"
	"time"
)

func GetProcessesBDD(ProcessBDDArray *Structs.ProcessBDDArray) error {
	// Get /Proceses from APIHost, with query organization_id = UipathOrg.Id
	client := new(http.Client)
	req, err := http.NewRequest("GET", "http://"+Utils.APIHost+"/Processes", nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// add query parameters
	q := req.URL.Query()
	q.Add("organization_id", Utils.UipathOrg.Id)
	req.URL.RawQuery = q.Encode()
	// make request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()
	// read response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// unmarshal response

	err = json.Unmarshal(body, &ProcessBDDArray)
	if err != nil {
		return nil
	}
	return err
}

func MonitorThread() {
	ProcessBDDArray := new(Structs.ProcessBDDArray)
	err := GetProcessesBDD(ProcessBDDArray)
	if err != nil {
		fmt.Println(err)
		return
	}
	FoldersID := ProcessBDDArray.FilterUniqueFoldersID()
	for _, FolderID := range FoldersID {
		ProcessBDDARRByFolder := new(Structs.ProcessBDDArray)
		for _, ProcessBDD := range ProcessBDDArray.Processes {
			if ProcessBDD.Folderid == FolderID {
				ProcessBDDARRByFolder.Add(ProcessBDD)
			}
		}
		LogsResponse := new(Utils.LogResponse)
		err = LogsResponse.Get(FolderID)
		if err != nil {
			fmt.Println(err)
		}
		Logs := LogsResponse.FilterLogs(ProcessBDDARRByFolder.Get())
		// Sort the logs by time, older first
		sort.Slice(Logs, func(i, j int) bool {
			return Logs[i].TimeStamp.Before(Logs[j].TimeStamp)
		})
		mapJobKey := make(map[string][]Utils.Log)
		for _, Log := range Logs {
			// check if jobKey is in map
			if _, ok := mapJobKey[Log.JobKey]; ok {
				mapJobKey[Log.JobKey] = append(mapJobKey[Log.JobKey], Log)
			} else {
				mapJobKey[Log.JobKey] = make([]Utils.Log, 0)
			}
		}
		errors := make(map[string]Structs.Incidentes_Procesos)
		for _, LogsK := range mapJobKey {
			ProcessNameLog := LogsK[0].ProcessName
			var ProcessBDD Structs.ProcesosBDD
			for _, bdd := range ProcessBDDARRByFolder.Get() {
				if bdd.Nombre == ProcessNameLog {
					ProcessBDD = bdd
					break
				}
			}
			// Check if this process is in errors
			if _, ok := errors[ProcessBDD.Nombre]; ok {
				continue
			}
			warningTolerance := ProcessBDD.WarningTolerance
			errorTolerance := ProcessBDD.ErrorTolerance
			fatalTolerance := ProcessBDD.FatalTolerance
			warningCount := 0
			errorCount := 0
			fatalCount := 0
			for _, Log := range LogsK {
				switch Log.Level {
				case "Warning":
					warningCount++
				case "Error":
					errorCount++
				case "Fatal":
					fatalCount++
				}
				if fatalCount > fatalTolerance {
					errors[ProcessNameLog] = Structs.Incidentes_Procesos{
						ProcesoID: ProcessBDD.ID,
						Incidente: "Error Fatal detectado",
						Tipo:      1,
						Estado:    1,
						Detalles: []Structs.Incidentes_Detalle{
							{
								Detalle:      Log.Message,
								Fecha_Inicio: time.Now(),
								Fecha_Fin:    time.Now(),
							},
						},
					}
					break
				} else if errorCount > errorTolerance {
					errors[ProcessNameLog] = Structs.Incidentes_Procesos{
						ProcesoID: ProcessBDD.ID,
						Incidente: "Exceso de errores detectado, por favor revisar consistencia del proceso",
						Tipo:      1,
						Estado:    1,
						Detalles: []Structs.Incidentes_Detalle{
							{
								Detalle:      Log.Message,
								Fecha_Inicio: time.Now(),
								Fecha_Fin:    time.Now(),
							},
						},
					}
					break
				} else if warningCount > warningTolerance {
					errors[ProcessNameLog] = Structs.Incidentes_Procesos{
						ProcesoID: ProcessBDD.ID,
						Incidente: "Exceso de advertencias detectado, por favor revisar consistencia del proceso",
						Tipo:      1,
						Estado:    1,
						Detalles: []Structs.Incidentes_Detalle{
							{
								Detalle:      Log.Message,
								Fecha_Inicio: time.Now(),
								Fecha_Fin:    time.Now(),
							},
						},
					}
					break
				}
			}
		}
		for _, ErroresDetectados := range errors {
			// Make POST request to API
			// Route /incidentes
			client := new(http.Client)
			incidentesJSON, err := json.Marshal(ErroresDetectados)
			if err != nil {
				fmt.Println(err)
			}
			req, err := http.NewRequest("POST", "http://"+Utils.APIHost+"/incidentes", bytes.NewBuffer(incidentesJSON))
			if err != nil {
				fmt.Println(err)
			}
			req.Header.Set("Content-Type", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(body))
		}
	}
}

func main() {
	Utils.QueryAuth.ClientId = os.Getenv("APP-ID")
	Utils.QueryAuth.ClientSecret = os.Getenv("APP-Secret")
	Utils.QueryAuth.Scope = os.Getenv("APP-Scope")
	Utils.UipathOrg.Id = os.Getenv("Organization-ID")
	Utils.LastMonitoredTime = time.Now()
	_ = Utils.RefreshUiPathToken()
	go func() {
		for {
			err := Utils.RefreshUiPathToken()
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(59 * time.Minute) // Refresh the token every 10 minutes
		}
	}()
	for {
		MonitorThread()
		Utils.LastMonitoredTime = time.Now()
		time.Sleep(15 * time.Second)
	}
}
