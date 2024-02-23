package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var boostingQuery = BoostingQuery{
	S: 100,
	Q: struct {
		B struct {
			S []interface{} "json:\"should\""
		} "json:\"bool\""
	}{
		B: struct {
			S []interface{} "json:\"should\""
		}{
			S: []interface{}{
				Query{
					MultiMatch{
						QueryString: "William Shatner Patrick Stewart",
						Fields: []string{
							"title",
							"overview",
							"cast.name",
							"directors.name",
						},
						Type: "cross_fields",
						// TieBreaker: "",
					},
				},
				Boost{
					MP: struct {
						Title struct {
							Q string  "json:\"query\""
							B float64 "json:\"boost,omitempty\""
						} "json:\"title\""
					}{
						Title: struct {
							Q string  "json:\"query\""
							B float64 "json:\"boost,omitempty\""
						}{
							Q: "star trek",
							B: 0.3,
						},
					},
				},
			},
		},
	},
}

func BoostingSearch() {
	bBody, _ := json.Marshal(boostingQuery)
	fmt.Println(string(bBody))
	req, err := http.NewRequest("GET", "http://localhost:9200/tmdb/_doc/_search", bytes.NewBuffer(bBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		return
	}

	c := &http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	bResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(string(bResp))
		return
	}
	var res Response
	json.Unmarshal(bResp, &res)
	for _, data := range res.BigHit.Hits {
		fmt.Printf("\n%-10s%-50s\t%10f", data.ID, data.Source["original_title"].(string), data.Score)
	}
}
