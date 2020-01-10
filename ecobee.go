package p

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Eco struct {
	Page struct {
		Page       int `json:"page"`
		TotalPages int `json:"totalPages"`
		PageSize   int `json:"pageSize"`
		Total      int `json:"total"`
	} `json:"page"`
	ThermostatList []struct {
		Identifier     string `json:"identifier"`
		Name           string `json:"name"`
		ThermostatRev  string `json:"thermostatRev"`
		IsRegistered   bool   `json:"isRegistered"`
		ModelNumber    string `json:"modelNumber"`
		LastModified   string `json:"lastModified"`
		ThermostatTime string `json:"thermostatTime"`
		UtcTime        string `json:"utcTime"`
		Runtime        struct {
			RuntimeRev         string `json:"runtimeRev"`
			Connected          bool   `json:"connected"`
			FirstConnected     string `json:"firstConnected"`
			ConnectDateTime    string `json:"connectDateTime"`
			DisconnectDateTime string `json:"disconnectDateTime"`
			LastModified       string `json:"lastModified"`
			LastStatusModified string `json:"lastStatusModified"`
			RuntimeDate        string `json:"runtimeDate"`
			RuntimeInterval    int    `json:"runtimeInterval"`
			ActualTemperature  int    `json:"actualTemperature"`
			ActualHumidity     int    `json:"actualHumidity"`
			DesiredHeat        int    `json:"desiredHeat"`
			DesiredCool        int    `json:"desiredCool"`
			DesiredHumidity    int    `json:"desiredHumidity"`
			DesiredDehumidity  int    `json:"desiredDehumidity"`
			DesiredFanMode     string `json:"desiredFanMode"`
		} `json:"runtime"`
	} `json:"thermostatList"`
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
}

func temp() string {
	client := http.Client{}
	url := `https://api.ecobee.com/1/thermostat?format=json&body={"selection":{"selectionType":"registered","selectionMatch":"","includeRuntime":true}}`
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}

	token := "Bearer " + os.Getenv("ECOBEE_TOKEN")
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "text/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(body))

	var eco Eco

	err = json.Unmarshal(body, &eco)

	temp := eco.ThermostatList[0].Runtime.ActualTemperature / 10

	return string(temp)

}

func sendTemp(m Message) {
	log.Println("setting up message")
	var response Response

	response.ChatID = m.Message.Chat.ID

	botURL := "https://api.telegram.org/bot" + token + "/sendMessage"

	temp := temp()

	message := fmt.Sprintf("The temperature is %s degrees F", temp)

	response.Text = message

	var body []byte

	body, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(body))

	req, err := http.Post(botURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()
}
