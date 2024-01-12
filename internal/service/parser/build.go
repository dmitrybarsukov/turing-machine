package parser

import (
	"errors"
	"fmt"
	"strconv"
	"turing-machine/internal/domain"
	"turing-machine/internal/service/validator"
)

func buildValidators(val yamlValidator) ([]domain.Validator, error) {
	if val.Compare != nil {
		return buildCompareValidators(val.Compare)
	}

	if val.Count != nil {
		return buildCountValidators(val.Count)
	}

	if val.Parity != nil {
		return buildParityValidators(val.Parity)
	}

	if val.MajorParity {
		return validator.HasMajorParity(), nil
	}

	if val.HasRepetitions {
		return validator.CountOfRepetitions(), nil
	}

	if val.HasPair {
		return validator.PairOfNumbersExist(), nil
	}

	if val.GreatestItem {
		return validator.ItemIsGreatest(), nil
	}

	if val.LeastItem {
		return validator.ItemIsLeast(), nil
	}

	if val.OutstandingItem {
		return validator.ItemIsOutstanding(), nil
	}

	if val.HasOrder {
		return validator.CodeHasOrder(), nil
	}

	if len(val.HasSequence) > 0 {
		return buildHasSequenceValidators(val.HasSequence)
	}

	return nil, errors.New("validator not specified")
}

func buildCompareValidators(data *yamlValidatorCompare) ([]domain.Validator, error) {
	if len(data.Item) > 0 {
		item, err := parseEnum(data.Item, mappingCodeItem)
		if err != nil {
			return nil, err
		}

		if targetItem, err := parseEnum(data.Target, mappingCodeItem); err == nil {
			return validator.ItemComparedToOtherItem(item, targetItem), nil
		}

		if targetNumber, err := strconv.Atoi(data.Target); err == nil {
			return validator.ItemComparedToConst(item, targetNumber), nil
		}

		return nil, errors.New("failed to recognize target")
	}

	if len(data.Sum) > 0 {
		items := make([]domain.CodeItem, 0, len(data.Sum))
		for _, itemStr := range data.Sum {
			item, err := parseEnum(itemStr, mappingCodeItem)
			if err != nil {
				return nil, err
			}

			items = append(items, item)
		}

		if targetNumber, err := strconv.Atoi(data.Target); err == nil {
			return validator.ItemsSumComparedToConst(items, targetNumber), nil
		}

		return nil, errors.New("failed to recognize target")
	}

	if data.AnyPair {
		return validator.AnyItemsPairCompared(), nil
	}

	if data.AnyItem {
		targetNumber, err := strconv.Atoi(data.Target)
		if err != nil {
			return nil, err
		}

		compare, err := parseEnum(data.Compare, mappingCompare)
		if err != nil {
			return nil, err
		}

		return validator.AnyItemComparedToConst(compare, targetNumber), nil
	}

	return nil, errors.New("comparator has nothing to compare")
}

func buildCountValidators(data *yamlValidatorCount) ([]domain.Validator, error) {
	if data.Number != 0 {
		return validator.CountOfNumber(data.Number), nil
	}

	if data.Parity != "" {
		parity, err := parseEnum(data.Parity, mappingParity)
		if err != nil {
			return nil, err
		}

		return validator.CountOfParity(parity), nil
	}

	if len(data.OneOf) > 0 {
		if len(data.OneOf) != 2 {
			return nil, errors.New("one_of should contain two elements")
		}

		return validator.CountOfAnyNumber(data.OneOf...), nil
	}

	return nil, errors.New("counter has nothing to count")
}

func buildParityValidators(data *yamlValidatorParity) ([]domain.Validator, error) {
	if len(data.Item) > 0 {
		item, err := parseEnum(data.Item, mappingCodeItem)
		if err != nil {
			return nil, err
		}

		return validator.ItemHasParity(item), nil
	}

	if data.Sum {
		return validator.SumHasParity(), nil
	}

	return nil, errors.New("parity checker has nothing to check")
}

func buildHasSequenceValidators(order string) ([]domain.Validator, error) {
	switch order {
	case "any":
		return validator.HasAnySequence(), nil
	case "asc":
		return validator.HasSequence(validator.Ascending), nil
	case "desc":
		return validator.HasSequence(validator.Descending), nil
	default:
		return nil, fmt.Errorf("unknown order %s", order)
	}
}
