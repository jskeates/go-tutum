package tutum

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

/*
func ListContainers
Returns : Array of Container objects
*/
func ListContainers() (CListResponse, error) {

	url := "container/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response CListResponse
	var finalResponse CListResponse

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
			var nextResponse CListResponse
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
func GetContainer
Argument : uuid
Returns : Container JSON object
*/
func GetContainer(uuid string) (Container, error) {

	url := ""
	if string(uuid[0]) == "/" {
		url = uuid[8:]
	} else {
		url = "container/" + uuid + "/"
	}

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Container

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

// Logs returns the container's logs through the given channel, wrapped in a Logs struct.
// See https://docs.tutum.co/v2/api/?go#get-the-logs-of-a-container for more info.
func (self *Container) Logs(c chan Logs) {

	endpoint := "container/" + self.Uuid + "/logs/?user=" + User + "&token=" + ApiKey
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
			//Type switches don't work here, so we use a type assertion instead
			_, ok := err.(*websocket.CloseError)
			if ok {
				close(c)
				break
			} else {
				log.Println(err)
			}
		}
		c <- msg
	}
}

// Exec runs the command in the container, and prints the output to the console.
// When the command exits and the final output has been printed, the function returns.

func (self *Container) Exec(command string, c chan Exec) {
	go self.Run(command, c)
Loop:
	for {
		select {
		case s, open := <-c:
			if open {
				fmt.Printf("%s", s.Output)
			} else {
				break Loop
			}
		}
	}
}

// Run executes the command in the container, and returns the output to the
// channel, wrapped in an Exec struct. When the command exits, the channel closes.
// See https://docs.tutum.co/v2/api/?go#execute-command-inside-a-container for more info.
func (self *Container) Run(command string, c chan Exec) {

	endpoint := "container/" + self.Uuid + "/exec/?user=" + User + "&token=" + ApiKey + "&command=" + url.QueryEscape(command)
	url := StreamUrl + endpoint

	header := http.Header{}
	header.Add("User-Agent", customUserAgent)

	var Dialer websocket.Dialer
	ws, _, err := Dialer.Dial(url, header)
	if err != nil {
		log.Println(err)
	}

	var msg Exec
Loop:
	for {
		if err = ws.ReadJSON(&msg); err != nil {
			//Type switches don't work here, so we use a type assertion instead
			_, ok := err.(*websocket.CloseError)
			if ok {
				close(c)
				break Loop
			} else {
				log.Println(err)
			}
		}
		c <- msg
	}
}

/*
func StartContainer
Returns : Error
*/
func (self *Container) Start() error {

	url := "container/" + self.Uuid + "/start/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)
	var response Container

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
func StopContainer
Returns : Error
*/
func (self *Container) Stop() error {

	url := "container/" + self.Uuid + "/stop/"
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
func RedeployContainer
Returns : Error
*/
func (self *Container) Redeploy() error {

	url := "container/" + self.Uuid + "/redeploy/"
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
func TerminateContainer
Returns : Error
*/
func (self *Container) Terminate() error {

	url := "container/" + self.Uuid + "/"
	request := "DELETE"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}
