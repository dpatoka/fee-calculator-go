package main

import (
	"fee-calculator-go/internal/pricing/interface/cli"
	"flag"
	"fmt"
	"log"
)

var (
	amountFlag   = flag.Float64("amount", 0, "Amount of money to lend")
	durationFlag = flag.Int("duration", 0, "Duration period for lending money")
)

func main() {
	flag.Parse()

	if *amountFlag <= 0 {
		log.Fatal("Amount must be above 0")
	}

	if *durationFlag <= 0 {
		log.Fatal("Duration can be 12 or 24")
	}

	command := cli.NewFeeCalculationCommand()
	result := command.Execute(*amountFlag, *durationFlag)

	fmt.Println(result)
}
