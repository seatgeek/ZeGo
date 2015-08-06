package zego

import (
    "encoding/json"
)

type Search_Results struct {
    Count            int `json:"count"`
    NextPage         string `json:"next_page"`
    PrevPage         string `json:"prev_page"`
    Results          []*Result `json:"results"`
}


type Result struct {
    Name             string `json:"name"`
    CreatedAt        string `json:"created_at"`
    UpdatedAt        string `json:"updated_at"`
    Id               int `json:"id"`
    ResultType       string `json:"result_type"`
    Url              string `json:"url"`
}

func (a Auth) Search(query string) (*ResultsStruct, error) {

    ResultsStruct := &Search_Results{}

    path := "/search.json?query=" + query
    resource, err := api(a, "GET", path, "")
    if err != nil {
        return nil, err
    }

    json.Unmarshal([]byte(resource.Raw), ResultsStruct)

    return ResultsStruct, nil

}



