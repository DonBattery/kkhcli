package main

import (
	"encoding/json"
	// "time"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
	"os"
)

type User struct {
	Username string `json:"username"`
	Enabled bool `json:"enabled"`
}

func connErr(msg string) {
	fmt.Println("\n", msg, "\n")
	os.Exit(3)
}

func seedDB() Response {
	var jsonStr = []byte(`{"Command":"seed"}`)
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("adminpassword", envPass)
	req.Header.Set("Content-Type", "application/json")	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			panic(err)
	}
	defer resp.Body.Close()	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}	
	msg:= Response{}
	json.Unmarshal(body, &msg)
	return msg
}

func getUsers(envPass string) []User {
	var jsonStr = []byte(`{"Command":"list"}`)
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("adminpassword", envPass)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		connErr("cannot connect to " + apiURL)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}	
	if resp.StatusCode != 200 {
		connErr("cannot connect to " + apiURL)
	}
	msg:= Response{}
	json.Unmarshal(body, &msg)
	if msg.Msg == "you have no rights to do this" {
		connErr("you have no rights to do this")
	}
	var users []User
	json.Unmarshal(body, &users)
	return users	
}

func addUser(envPass string, username string, password string) Response {	
	type Message struct {
		Command string
    Username string
    Password string
	}	
	m := Message{"add", username, password}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)		
	}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(b))
	req.Header.Set("adminpassword", envPass)
	req.Header.Set("Content-Type", "application/json")	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			panic(err)
	}
	defer resp.Body.Close()	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}	
	msg:= Response{}
	json.Unmarshal(body, &msg)
	return msg
}

func resetUser(envPass string, username string, password string) Response {	
	type Message struct {
		Command string
    Username string
    Password string
	}	
	m := Message{"reset", username, password}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)		
	}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(b))
	req.Header.Set("adminpassword", envPass)
	req.Header.Set("Content-Type", "application/json")	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			panic(err)
	}
	defer resp.Body.Close()	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}	
	msg:= Response{}
	json.Unmarshal(body, &msg)
	return msg
}