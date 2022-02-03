package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"encoding/csv"
	"github.com/jszwec/csvutil"
)

type Person struct {
	Name        string `csv:"Name"`
	PhoneNumber string `csv:"Phone Number"`
	Age         int    `csv:"Age"`
}

func main() {
	file, err := os.Open("csvToStructParser/main/people.csv")
	if err != nil {
		log.Fatal(err)
	}
	err = ParseCsvUsingHeaders(file, func(person Person) {
		fmt.Println(fmt.Sprintf(
			"Person: {name: '%s', phoneNumber: '%s', age='%d'}",
			person.Name, person.PhoneNumber, person.Age))
	})
	if err != nil {
		log.Fatal(err)
	}
}

type ProcessorFunction func(person Person)

func ParseCsvUsingHeaders(reader io.Reader, function ProcessorFunction) error {
	decoder, err := csvutil.NewDecoder(csv.NewReader(reader))
	if err != nil {
		log.Fatal(err)
	}
	waitGroup := sync.WaitGroup{}
	for {
		person := &Person{}
		if err = decoder.Decode(person); err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return err
		}
		waitGroup.Add(1)
		go func(person Person) {
			function(person)
			waitGroup.Done()
		}(*person)
	}
	waitGroup.Wait()
	return err
}
