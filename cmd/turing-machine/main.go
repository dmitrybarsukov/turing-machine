package main

import (
	"fmt"
	"os"
	"turing-machine/internal/analyze"
	"turing-machine/internal/domain"
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
	result, err := loadFile(fileName)
	if err != nil {
		return err
	}

	machines := result.HyperMachine.GetAllMachines()
	fmt.Printf("Having total %d machines\n", len(machines))

	machines = lo.Filter(machines, func(m turing.Machine, _ int) bool {
		return m.HasSolution()
	})

	fmt.Printf("Found %d possible machines:\n", len(machines))

	for _, t := range result.Tests {
		machines = t.FilterMachines(machines)
	}

	fmt.Printf("%d machines remained after filtering:\n", len(machines))

	for _, m := range machines {
		fmt.Println(m)
	}

	if len(machines) == 1 {
		fmt.Printf("\nFound solution: %v\n\n", machines[0].Solution())
		return nil
	}

	codes := lo.Map(machines, func(it turing.Machine, _ int) domain.Code {
		return it.Solution()
	})

	codeStats := analyze.Codes(codes)
	fmt.Println("\nCodes stats:")
	for _, c := range codeStats {
		fmt.Printf("\t%v\n", c)
	}

	validatorStats := analyze.Validators(machines)
	fmt.Println("\nValidators stats:")
	for _, v := range validatorStats {
		fmt.Printf("\t%v\n", v)
	}

	fmt.Println()

	if len(result.Tests)%3 == 0 {
		recommendCode(machines, validatorStats)
	}
	recommendValidator(machines, validatorStats)

	fmt.Println()

	return nil
}

func loadFile(fileName string) (parser.ParseResult, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return parser.ParseResult{}, err
	}

	defer func() {
		_ = file.Close()
	}()

	result, err := parser.Parse(file)
	if err != nil {
		return parser.ParseResult{}, err
	}

	return result, nil
}

func recommendCode(machines []turing.Machine, validatorStats []analyze.ValidatorStats) {
	if len(machines) == 0 {
		return
	}

	for _, stat := range validatorStats {
		newMachines := lo.Filter(machines, func(it turing.Machine, _ int) bool {
			return it.Validators[stat.Key] == stat.Validator
		})

		if len(newMachines) == 0 {
			break
		}

		machines = newMachines
	}

	fmt.Printf("Try code %v\n", machines[0].Solution())
}

func recommendValidator(_ []turing.Machine, validatorStats []analyze.ValidatorStats) {
	leastConfident := lo.MinBy(validatorStats, func(stats, stats2 analyze.ValidatorStats) bool {
		return stats.Confidence < stats2.Confidence
	})

	fmt.Printf("Check validator %c\n", leastConfident.Key)
}
