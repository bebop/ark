package cmd

import (
	"fmt"
	"log"
	"net/http"
)

func ExampleIsEqualFiles() {
	response, err := http.Get("https://ftp.expasy.org/databases/rhea/rdf/rhea.rdf.gz")

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	fmt.Println(isEqualFiles("../data/build", "rhea.rdf.gz", response.Body))
	//Output: true
}
