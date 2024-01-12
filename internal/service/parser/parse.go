package parser

import (
	"errors"
	"fmt"
	"io"
	"turing-machine/internal/domain"
	"turing-machine/internal/service/turing"
	"unicode"

	"gopkg.in/yaml.v3"
)

func Parse(reader io.Reader) (ParseResult, error) {
	root, err := parseYaml(reader)
	if err != nil {
		return ParseResult{}, err
	}

	machine := turing.HyperMachine{
		Validators: make(map[rune][]domain.Validator),
		Codes:      domain.AllCodes,
	}

	for key, validator := range root.Validators {
		if len(key) > 1 {
			return ParseResult{}, fmt.Errorf("validator '%s' has too long name", key)
		}

		letter := unicode.ToUpper([]rune(key)[0])
		if letter < 'A' || letter > 'Z' {
			return ParseResult{}, fmt.Errorf("validator '%s' has wrong name", key)
		}

		validators, err := buildValidators(validator)
		if err != nil {
			return ParseResult{}, fmt.Errorf("validator '%s' failed build: %w", key, err)
		}

		machine.Validators[letter] = validators
	}

	if len(machine.Validators) == 0 {
		return ParseResult{}, errors.New("no validators specified")
	}

	tests := make([]Test, 0, len(root.Tests)*3)
	for codeInt, checks := range root.Tests {
		for key, result := range checks {
			tests = append(tests, Test{
				Code:      domain.Code{codeInt / 100 % 10, codeInt / 10 % 10, codeInt % 10},
				Validator: unicode.ToUpper([]rune(key)[0]),
				Result:    result,
			})
		}
	}

	return ParseResult{
		HyperMachine: machine,
		Tests:        tests,
	}, nil
}

func parseYaml(reader io.Reader) (yamlRoot, error) {
	var result yamlRoot
	err := yaml.NewDecoder(reader).Decode(&result)
	return result, err
}
