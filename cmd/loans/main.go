package main

import (
	"fee-calculator-go/internal/pricing/interface/cli"
	"flag"
	"fmt"
	"log"
)

var (
	amountFlag = flag.Float64("amount", 0, "Amount of money to lend")
	termFlag   = flag.Int("term", 0, "Term period for lending money")
)

func main() {
	flag.Parse()

	if *amountFlag <= 0 {
		log.Fatal("Amount must be above 0")
	}

	if *termFlag <= 0 {
		log.Fatal("Term can be 12 or 24")
	}

	command := cli.NewFeeCalculationCommand()
	result, err := command.Execute(*amountFlag, *termFlag)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
