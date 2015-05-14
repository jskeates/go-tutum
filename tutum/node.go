package tutum

import "encoding/json"

/*
func ListNodes
Returns : Array of Node objects
*/
func ListNodes() (NodeListResponse, error) {

	url := "node/"
	request := "GET"

	//Empty Body Request
	body := []byte(`{}`)
	var response NodeListResponse

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
func GetNode
Argument : uuid
Returns : Node JSON object
*/
func GetNode(uuid string) (Node, error) {

	url := ""
	if string(uuid[0]) == "/" {
		url = uuid[8:]
	} else {
		url = "node/" + uuid + "/"
	}

	request := "GET"
	body := []byte(`{}`)
	var response Node

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
func UpdateNode
Argument : uuid
Returns : Node JSON object
*/
func (self *Node) Update(requestBody string) error {

	url := "node/" + self.Uuid + "/"
	request := "PATCH"

	updatedNode := []byte(requestBody)

	_, err := TutumCall(url, request, updatedNode)
	if err != nil {
		return err
	}

	return nil
}

/*
func UpgradeDaemon
Argument : uuid
Returns : Node JSON object
*/
func (self *Node) Upgrade() error {

	url := "node/" + self.Uuid + "/docker-upgrade/"
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
func TerminateNode
Argument : uuid
Returns : Node JSON object
*/
func (self *Node) Terminate() error {

	url := "node/" + self.Uuid + "/"
	request := "DELETE"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}
