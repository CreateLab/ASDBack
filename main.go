package main

import (
	"ASDBack/ASDBack"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "http://127.0.0.1:7860/"

func getPing(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got ping request\n")
	_, err := io.WriteString(w, "pong\n")
	if err != nil {
		fmt.Println(err)
	}
}

func postRedirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}
	var requestData ASDBack.RedirectRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	result := ASDBack.CheckApiKey(requestData.ApiKey)

	if !result {
		fmt.Println("API key does not exist in the database")
		return
	}

	var resp *http.Response
	if requestData.Method == "GET" {
		resp, err = http.Get(url + requestData.Url)
	}

	if requestData.Method == "POST" {
		resp, err = http.Post(url+requestData.Url, "application/json", strings.NewReader(requestData.Data))
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	w.Write(responseBody)

}

func main() {
	http.HandleFunc("/ping", getPing)
	http.HandleFunc("/postRedirect", postRedirect)

	ASDBack.SetUp()

	err := http.ListenAndServe(":9898", nil)
	if err != nil {
		fmt.Println(err)
	}

}
