package main

import (
	"fmt"
	"os"
	"turing-machine/internal/domain/turing"
	"turing-machine/internal/parser"

	"github.com/samber/lo"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Use: %s <task file name>\n", os.Args[0])
		return
	}

	err := run(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func run(fileName string) error {
	hyperMachine, err := loadHyperMachine(fileName)
	if err != nil {
		return err
	}

	machines := hyperMachine.GetAllMachines()
	fmt.Printf("Having total %d machines\n", len(machines))

	possibleMachines := lo.Filter(machines, func(m turing.Machine, _ int) bool {
		return m.HasSolution()
	})

	fmt.Printf("Found %d POSSIBLE machines:\n", len(possibleMachines))
	for _, m := range possibleMachines {
		fmt.Println(m)
	}

	return nil
}

func loadHyperMachine(fileName string) (turing.HyperMachine, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return turing.HyperMachine{}, err
	}

	defer func() {
		_ = file.Close()
	}()

	hyperMachine, err := parser.ParseHyperMachine(file)
	if err != nil {
		return turing.HyperMachine{}, err
	}

	return hyperMachine, nil
}
