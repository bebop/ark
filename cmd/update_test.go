package cmd

import (
	"fmt"
	"log"
	"net/http"
)

func ExampleCreateCheckSumFile() {
	response, err := http.Get("https://ftp.expasy.org/databases/rhea/rdf/rhea.rdf.gz")

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	CreateCheckSumFile("../data/build", "rhea.rdf.gz", response.Body)
	fmt.Println("Hello")
	// Output: Hell
}

func ExampleIsEqualFiles() {
	response, err := http.Get("https://ftp.expasy.org/databases/rhea/rdf/rhea.rdf.gz")

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	fmt.Println(isEqualFiles("../data/build", "rhea.rdf.gz", response.Body))
	//Output: true
}
