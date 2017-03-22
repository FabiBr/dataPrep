package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Paths struct {
	TweetList       string
	PartyHandleList string
	PartyResultList string
	OldResultList   string
	NewResultList   string
}

var paths = *NewPaths()

func main() {
	paths.TweetList = "data/all-tweets.csv"
	paths.PartyHandleList = "data/BundestagKurz.csv"
	paths.PartyResultList = "data/results.csv"
	paths.OldResultList = "data/normalized_dataset.csv"
	paths.NewResultList = "data/fixed_normalized_dataset.csv"

	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	//check if results.csv exists, if not create it
	//if _, err := os.Stat(paths.PartyResultList); os.IsNotExist(err) {
	readDataset(getHandles())
	//}

}

func fixShittyResults() {

}

func readDataset(handles [][]string) {
	readFile, _ := os.Open(paths.TweetList)
	defer readFile.Close()
	reader := csv.NewReader(bufio.NewReader(readFile))

	writeFile, err := os.Create(paths.PartyResultList)
	Debug.Println(err)
	defer writeFile.Close()
	writer := csv.NewWriter(writeFile)
	//var counter = 0
	firstline, _ := reader.Read()
	writer.Write(firstline)
	for {
		line, err := reader.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		var party = ""
		var partyFlag = false
		var currentDude = strings.ToLower(line[2])

		for _, values := range handles {

			//break when handles.csv is shit
			if len(values) <= 5 {
				Debug.Println(values)
				break
			}

			twitterHandle := values[5]

			//remove whitespace from shitty csv
			twitterHandle = SpaceMap(twitterHandle)

			//lowercase
			twitterHandle = strings.ToLower(twitterHandle)

			//remove @ from handles why ever
			if len(twitterHandle) > 0 {
				if twitterHandle[0] == '@' {
					twitterHandle = twitterHandle[1:]
				}
			}

			//check name in line with each handle in List
			if currentDude == twitterHandle {
				//Debug.Println(line[2] + ": " + values[1])
				partyFlag = true
				party = values[1]
				line = append(line, party)
				writer.Write(line)
				break
			}
		}

		if !partyFlag {
			Debug.Println("Not in list : " + currentDude)
			line = append(line, "NOTINLIST")
			writer.Write(line)
		}

		// counter++
		// if counter > 10 {
		// 	break
		// }
	}
	defer writer.Flush()
}

func getHandles() [][]string {
	file, _ := os.Open(paths.PartyHandleList)
	reader := csv.NewReader(bufio.NewReader(file))
	reader.Comma = ';'

	handles, _ := reader.ReadAll()
	return handles
}
