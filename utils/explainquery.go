package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Explain() {
	bBody, _ := json.Marshal(query)
	req, err := http.NewRequest("GET", "http://localhost:9200/tmdb/_validate/query?explain", bytes.NewBuffer(bBody))
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
	var res ExplainResponse
	json.Unmarshal(bResp, &res)
	fmt.Println(string(bResp))
	fmt.Println(res.Explanations)

}
