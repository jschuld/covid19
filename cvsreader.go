package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Entry struct {
	province string
	country string
	lat string
	lon string
	day1 int
	day2 int
	day3 int
}

func main() {
	// Open the file
	csvfile, err := os.Open("confirmed.csv")
	csvfile_recovered, err2 := os.Open("recovered.csv")
	csvfile_deaths, err3 := os.Open("deaths.csv")
	if err != nil {
		log.Fatalln("Couldn't open the confirmed.csv file", err)
	}

	if err2 != nil {
		log.Fatalln("Couldn't open the recovered.csv file", err)
	}

	if err3 != nil {
		log.Fatalln("Couldn't open the deaths.csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	r_recovered := csv.NewReader(csvfile_recovered)
	r_deaths := csv.NewReader(csvfile_deaths)

	//read header
	header, err := r.Read()
	header2, _ := r_recovered.Read()
	header3, _ := r_deaths.Read()
	fmt.Println("date;country;province;lat;lon;confirmed;recovered;current;deaths;confirmed-increase")	
	header2 = header2
	header3 = header3
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		record_recovered, err2 := r_recovered.Read()
		record_deaths, err3 := r_deaths.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if err2 == io.EOF {
			break
		}
		if err2 != nil {
			log.Fatal(err2)
		}

		if err3 == io.EOF {
			break
		}
		if err3 != nil {
			log.Fatal(err3)
		}

		entry := Entry{province: record[0], country: record[1], lat: record[2], lon: record[3]}

		diagnosed_yesterday := 0

		// Iterate through the records
		for i, s := range record {
			if i > 3 {
				current1, _ := strconv.Atoi(s)
				current2, _ := strconv.Atoi(record_recovered[i])
				current := current1 - current2
				increase := current1 - diagnosed_yesterday
				diagnosed_yesterday = current1
									
				fmt.Printf("%s;%s;%s;%s;%s;%s;%s;%d;%s;%d\n",
					header[4 + i-4],
					entry.country,
					entry.province,
					entry.lat,
					entry.lon,
					s,
					record_recovered[i],
					current,
					record_deaths[i],
					increase)
			}
		}
	}
}
