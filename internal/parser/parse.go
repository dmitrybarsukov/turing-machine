package parser

import (
	"errors"
	"fmt"
	"io"
	"turing-machine/internal/domain"
	"turing-machine/internal/domain/turing"
	"unicode"

	"gopkg.in/yaml.v3"
)

func ParseHyperMachine(reader io.Reader) (turing.HyperMachine, error) {
	root, err := parseYaml(reader)
	if err != nil {
		return turing.HyperMachine{}, err
	}

	machine := turing.HyperMachine{
		Validators: make(map[rune][]domain.Validator),
		Codes:      domain.AllCodes,
	}

	for key, validator := range root.Validators {
		if validator.Disabled {
			continue
		}

		if len(key) > 1 {
			return machine, fmt.Errorf("validator '%s' has too long name", key)
		}

		letter := unicode.ToUpper([]rune(key)[0])
		if letter < 'A' || letter > 'Z' {
			return machine, fmt.Errorf("validator '%s' has wrong name", key)
		}

		validators, err := buildValidators(validator)
		if err != nil {
			return machine, fmt.Errorf("validator '%s' failed build: %w", key, err)
		}

		machine.Validators[letter] = validators
	}

	if len(machine.Validators) == 0 {
		return machine, errors.New("no validators specified")
	}

	return machine, nil
}

func parseYaml(reader io.Reader) (yamlRoot, error) {
	var result yamlRoot
	err := yaml.NewDecoder(reader).Decode(&result)
	return result, err
}
