package main

import (
	"database/sql"
	"fmt"

	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	product := "pyruvate"
	substrate := "glucose"
	recurse(product, substrate)



}

type Rxns struct {
	reactants, products string
}

// in: product, function outputs reactions that give you that product
func dig(product string) []Rxns {
	//fmt.Println("Testing")
	database, _ := sql.Open("sqlite3", "./allbase.db")
	rows, _ := database.Query("SELECT equation FROM reaction WHERE equation LIKE '%" + product + "%'")
	var equation string
	rxns := make([]Rxns, 0, 30)

	for rows.Next() {
		rows.Scan(&equation)
		reversable := strings.Contains(equation, "<=>")
		if reversable {
			vec := strings.Split(equation, "<=>")
			if strings.Contains(vec[0], product) {
				rxn := Rxns{vec[1], vec[0]}
				rxns = append(rxns, rxn)
			} else {
				rxn := Rxns{vec[0], vec[1]}
				rxns = append(rxns, rxn)
			}
		} else {
			vec := strings.Split(equation, "=>")
			if strings.Contains(vec[1], product) {
				rxn := Rxns{vec[0], vec[1]}
				rxns = append(rxns, rxn)

			}

		}
	}
	fmt.Println(rxns)
	return rxns
}


//base case: if product is in a level of rxns, you've won. else, repeat the function. reactants are the below level's products, set each reactant as the product. if base case, return list of pathways
// tester recursion function supposed to just verify if it can go down the tree and find a match
func recurse(product string, substrate string) {
	// database, _ := sql.Open("sqlite3", "./allbase.db")
	if product == substrate {
		fmt.Println("found:", substrate)
	}
	rxns := dig(product)
	for i := range rxns {
		reactants := strings.Split(rxns[i].reactants, " + ")
		for j := range reactants {
			product = reactants[j]
			// count = count + 1
			recurse(product, substrate)

		}
	}
}
