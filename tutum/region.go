package tutum

import "encoding/json"

type RegionListResponse struct {
	Objects []Region `json:"objects"`
}

type Region struct {
	Available    bool     `json:"available"`
	Label        string   `json:"label"`
	Name         string   `json:"name"`
	Node_types   []string `json:"node_types"`
	Provider     string   `json:"provider"`
	Resource_uri string   `json:"resource_uri"`
}

/*
func ListRegions
Returns : Array of Region objects
*/
func ListRegions() (RegionListResponse, error) {

	url := "region/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response RegionListResponse

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
func GetRegion
Argument : provider name and location name
Returns : Region JSON object
*/
func GetRegion(id string) (Region, error) {

	url := ""
	if string(id[0]) == "/" {
		url = id[8:]
	} else {
		url = "region/" + id + "/"
	}

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Region

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
