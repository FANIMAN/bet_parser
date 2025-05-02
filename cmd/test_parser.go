package main

import (
	"fmt"
	"log"

	"betting-parser/internal/parser"
)

func mainn() {
	// Test parsing Cricket Prematch data
	cricketData, err := parser.ParseCricketPrematch("./testdata/cricket_prematch.json")
	if err != nil {
		log.Fatalf("Error parsing cricket prematch data: %v", err)
	}
	fmt.Printf("Parsed Cricket Prematch Data: %+v\n", cricketData)

	// Test parsing Cricket Result data
	cricketResultData, err := parser.ParseCricketResult("./testdata/cricket_result.json")
	if err != nil {
		log.Fatalf("Error parsing cricket result data: %v", err)
	}
	fmt.Printf("Parsed Cricket Result Data: %+v\n", cricketResultData)

	// Test parsing Volleyball Prematch data
	volleyballData, err := parser.ParseVolleyballPrematch("./testdata/volleyball_prematch.json")
	if err != nil {
		log.Fatalf("Error parsing volleyball prematch data: %v", err)
	}
	fmt.Printf("Parsed Volleyball Prematch Data: %+v\n", volleyballData)

	// Test parsing Volleyball Result data
	volleyballResultData, err := parser.ParseVolleyballResult("./testdata/volleyball_result.json")
	if err != nil {
		log.Fatalf("Error parsing volleyball result data: %v", err)
	}
	fmt.Printf("Parsed Volleyball Result Data: %+v\n", volleyballResultData)
}
