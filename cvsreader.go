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
	if err != nil {
		log.Fatalln("Couldn't open the confirmed.csv file", err)
	}

	if err2 != nil {
		log.Fatalln("Couldn't open the recovered.csv file", err)
	}


	// Parse the file
	r := csv.NewReader(csvfile)
	r_recovered := csv.NewReader(csvfile_recovered)

	//read header
	header, err := r.Read()
	header2, _ := r_recovered.Read()
	fmt.Println("date;country;province;lat;lon;confirmed;recovered;current")	

	header2 = header2
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		record_recovered, err2 := r_recovered.Read()
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

		entry := Entry{province: record[0], country: record[1], lat: record[2], lon: record[3]}

		// Iterate through the records
		for i, s := range record {
			if i > 3 {
			current1, _ := strconv.Atoi(s)
			current2, _ := strconv.Atoi(record_recovered[i])
			current := current1 - current2	
			fmt.Printf("%s;%s;%s;%s;%s;%s;%s;%d\n",
				header[4 + i-4],
				entry.country,
				entry.province,
				entry.lat,
				entry.lon,
				s,
				record_recovered[i],
				current)
			}
		}
	}
}
