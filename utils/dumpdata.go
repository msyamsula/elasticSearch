package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func DumpData() {
	f, err := os.Open("tmdb.json")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	bData, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
		return
	}

	x := map[string]interface{}{}
	json.Unmarshal(bData, &x)
	c := &http.Client{}
	// log.Println([]byte(x["93837"].(map[string]interface{}{})))
	// body := []byte{}

	fullCmd := ""
	i := 0
	for id := range x {
		command := fmt.Sprintf(`{"index": {"_index": "tmdb","_type": "_doc","_id":%s}}`, id)
		bMovie, _ := json.Marshal(x[id])
		var cmd string
		cmd += fmt.Sprintf("%s\n%s\n", command, string(bMovie))
		fullCmd += cmd
		i++
	}

	body := bytes.NewBuffer([]byte(fullCmd))
	req, err := http.NewRequest("POST", "http://localhost:9200/_bulk", body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {

		log.Println(err)
		return
	}

	response, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	bRes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bRes))

}
