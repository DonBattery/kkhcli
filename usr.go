package main

import (
	"encoding/json"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
	"os"
)

type Response struct {
	Msg string `json:"msg"`
}

type User struct {
	Username string `json:"username"`
}

type Collection struct {
	Name string `json:"name"`
}

func connErr(msg string) {
	fmt.Println("\n", msg, "\n")
	os.Exit(3)
}

func flushColllection(collectionToFlush string) Response {
	type Message struct {
		Command string
		Collection string
	}	
	m := Message{"flush", collectionToFlush}
	fmt.Println(m)
	b, err := json.Marshal(m)
	if err != nil {
		connErr("JSON Error")		
	}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(b))
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
		connErr("Caannot read response")
	}	
	msg:= Response{}
	json.Unmarshal(body, &msg)
	return msg	
}

func seedDB() Response {
	var jsonStr = []byte(`{"Command":"seed"}`)
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
		connErr("Caannot read response")		
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
		connErr("Cannot connect to " + apiURL)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		connErr("Cannot read response")
	}	
	if resp.StatusCode != 200 {
		connErr("Error connecting to " + apiURL)
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

func getCollections(envPass string) []Collection {
	var jsonStr = []byte(`{"Command":"collections"}`)
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
		connErr("Cannot read response")
	}	
	if resp.StatusCode != 200 {
		connErr("Error connecting to " + apiURL)
	}
	msg:= Response{}
	json.Unmarshal(body, &msg)
	if msg.Msg == "you have no rights to do this" {
		connErr("you have no rights to do this")
	}
	var collections []Collection
	json.Unmarshal(body, &collections)
	return collections
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