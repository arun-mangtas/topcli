package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Flag to read port
	port := flag.Int("port", 8080, "Server listens on this port")
	flag.Parse()

	if len(os.Args) < 2 {
		log.Fatal("Missing text argument")
	}

	text := os.Args[1]

	client := &http.Client{}

	b, err := json.Marshal(struct {
		Text string `json:"text"`
	}{
		Text: text,
	})
	if err != nil {
		log.Fatalf("Error marshalling body %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:"+strconv.Itoa(*port)+"/top/ten/words", bytes.NewBuffer(b))
	if err != nil {
		log.Fatalf("Error while creating a new request %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body %v", err)
	}

	var res []struct {
		Word  string `json:"word"`
		Count int    `json:"count"`
	}

	err = json.Unmarshal(respBody, &res)
	if err != nil {
		log.Fatalf("Error unmarshalling response body %v", err)
	}

	fmt.Printf("%+v", res)
}
