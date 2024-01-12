package parser

import (
	"errors"
	"strconv"
	"turing-machine/internal/domain"
	"turing-machine/internal/domain/validator"
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

	if val.HasMoreParity {
		return validator.HasMoreNumbersWithParity(), nil
	}

	if val.HasRepetitions {
		return validator.HasSomeRepeatingNumbers(), nil
	}

	if val.HasPair {
		return validator.PairOfNumbersExist(), nil
	}

	if val.GreatestItem {
		return validator.OneItemIsGreater(), nil
	}

	if val.LeastItem {
		return validator.OneItemIsLess(), nil
	}

	if val.OutlierItem {
		return validator.OneItemIsOutlier(), nil
	}

	if val.HasOrder {
		return validator.CodeIsOrdered(), nil
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

	if data.Multi {
		return validator.ItemsMultiComparable(), nil
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

	if len(data.OneOfNumbers) > 0 {
		if len(data.OneOfNumbers) != 2 {
			return nil, errors.New("one_of_numbers should contain two elements")
		}

		return validator.CountOfNumberOfOneOfTwo(data.OneOfNumbers[0], data.OneOfNumbers[1]), nil
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
