package tutum

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
func ListServices
Returns : Array of Service objects
*/
func ListServices() (SListResponse, error) {
	url := "service/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response SListResponse
	var finalResponse SListResponse

	data, err := TutumCall(url, request, body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	finalResponse = response

Loop:
	for {
		if response.Meta.Next != "" {
			var nextResponse SListResponse
			data, err := TutumCall(response.Meta.Next[8:], request, body)
			if err != nil {
				return nextResponse, err
			}
			err = json.Unmarshal(data, &nextResponse)
			if err != nil {
				return nextResponse, err
			}
			finalResponse.Objects = append(finalResponse.Objects, nextResponse.Objects...)
			response = nextResponse

		} else {
			break Loop
		}

	}

	return finalResponse, nil

}

/*
func GetService
Argument : uuid
Returns : Service JSON object
*/
func GetService(uuid string) (Service, error) {

	url := ""
	if string(uuid[0]) == "/" {
		url = uuid[8:]
	} else {
		url = "service/" + uuid + "/"
	}

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Service

	data, err := TutumCall(url, request, body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil

}

/*
func CreateService
Argument : Service JSON object (see documentation)
Returns : Service JSON object
*/
func CreateService(createRequest ServiceCreateRequest) (Service, error) {

	url := "service/"
	request := "POST"
	var response Service

	newService, err := json.Marshal(createRequest)
	if err != nil {
		return response, err
	}

	data, err := TutumCall(url, request, newService)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

/*
func Logs
Argument : a channel of type string for the output
*/

func (self *Service) Logs(c chan Logs) {

	endpoint := "service/" + self.Uuid + "/logs/?user=" + User + "&token=" + ApiKey
	url := StreamUrl + endpoint

	header := http.Header{}
	header.Add("User-Agent", customUserAgent)

	var Dialer websocket.Dialer
	ws, _, err := Dialer.Dial(url, header)
	if err != nil {
		log.Println(err)
	}

	var msg Logs
	for {
		if err = ws.ReadJSON(&msg); err != nil {
			if err != nil && err.Error() != "EOF" {
				log.Println(err)
			} else {
				break
			}
		}
		c <- msg
	}
}

/*
func UpdateService
Argument : updatedService JSON object
Returns : Error
*/
func (self *Service) Scale() error {

	url := "service/" + self.Uuid + "/scale/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)
	var response Service

	data, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return err
	}

	return nil
}

/*
func UpdateService
Argument : updatedService JSON object
Returns : Error
*/
func (self *Service) Update(createRequest ServiceCreateRequest) error {

	url := "service/" + self.Uuid + "/"
	request := "PATCH"

	updatedService, err := json.Marshal(createRequest)
	if err != nil {
		return err
	}

	_, err = TutumCall(url, request, updatedService)
	if err != nil {
		return err
	}

	return nil
}

/*
func StartService
Returns : Error
*/
func (self *Service) Start() error {
	url := "service/" + self.Uuid + "/start/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)
	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}

/*
func StopService
Returns : Error
*/
func (self *Service) StopService() error {

	url := "service/" + self.Uuid + "/stop/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)
	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}

/*
func RedeployService
Returns : Error
*/
func (self *Service) Redeploy() error {

	url := "service/" + self.Uuid + "/redeploy/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}

/*
func TerminateService
Returns : Error
*/
func (self *Service) TerminateService() error {
	url := "service/" + self.Uuid + "/"
	request := "DELETE"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}
