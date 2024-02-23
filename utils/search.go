package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var query = ExplainBody{
	Query: Query{
		MultiMatch: MultiMatch{
			QueryString: "William Shatner Patrick Stewart",
			Fields: []string{
				"title",
				"overview",
				"cast.name",
				"directors.name",
			},
			Type: "cross_fields",
			// TieBreaker: "0.4",
		},
	},
}

var searchQuery = JsonBody{
	Size:  100,
	Query: query.Query,
}

func Search() {

	bBody, _ := json.Marshal(searchQuery)
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
	var res Response
	json.Unmarshal(bResp, &res)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(string(bResp))
		return
	}
	for _, data := range res.BigHit.Hits {
		fmt.Printf("\n%-10s%-50s\t%10f", data.ID, data.Source["original_title"].(string), data.Score)
	}

}
