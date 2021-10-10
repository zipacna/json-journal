package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

var jsonFile = "journal.json"

func jsonFilePresent() bool {
	_, err := os.Stat(jsonFile)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

type Entry struct {
	Data      string
	Timestamp time.Time
}

// wtf writes to file.
func wtf(fileData []Entry) bool {
	res, err := json.Marshal(fileData)
	if err != nil {
		return false
	}

	err2 := ioutil.WriteFile(jsonFile, res, 0644)
	if err2 != nil {
		return false
	}

	//log.Println(string(res))
	return true
}

func writeJSON(data Entry) bool {
	var fileData []Entry
	if jsonFilePresent() {
		fileData = append(readJSON(), data)
	} else {
		fileData = append(fileData, data)
	}
	return wtf(fileData)
}

func readJSON() []Entry {
	content, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil
	}

	var data []Entry
	err2 := json.Unmarshal(content, &data)
	if err2 != nil {
		return nil
	}

	//log.Printf("%+v\n", data) // default formatting + struct fields https://pkg.go.dev/log
	return data
}

// commonUD are the shared parts of Update and Delete. update=true will update, while update=false will delete.
func commonUD(data Entry, update bool) bool {
	var fileData []Entry
	if jsonFilePresent() {
		fileData = readJSON()
		// Replace object in array on matching timestamps
		for i := 0; i < len(fileData); i++ {
			if fileData[i].Timestamp == data.Timestamp {
				if update {
					fileData[i] = data
				} else {
					fileData = append(fileData[:i], fileData[i+1:]...)
				}
				return wtf(fileData)
			}
		}
	}
	return false
}

func updateJSON(data Entry) bool {
	return commonUD(data, true)
}

func deleteJSON(data Entry) bool {
	return commonUD(data, false)
}
