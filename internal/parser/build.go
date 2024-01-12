package parser

import (
	"errors"
	"strconv"
	"turing-machine/internal/domain"
	"turing-machine/internal/domain/validator"
)

func buildValidators(validator yamlValidator) ([]domain.Validator, error) {
	if validator.Compare != nil {
		return buildCompareValidators(validator.Compare)
	}

	if validator.Counter != nil {
		return buildCounterValidators(validator.Counter)
	}

	if validator.Parity != nil {
		return buildParityValidators(validator.Parity)
	}

	if validator.HasMoreParity {
		return buildHasMoreParityValidators()
	}

	if validator.HasRepetitions {
		return buildHasRepetitionsValidators()
	}

	if validator.HasPair {
		return buildHasPairValidators()
	}

	if validator.GreatestItem {
		return buildGreatestItemValidators()
	}

	if validator.LeastItem {
		return buildLeastItemValidators()
	}

	if validator.HasOrder {
		return buildHasOrderValidators()
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
			return validator.ItemComparedToOtherItem(item, targetItem, allCompare), nil
		}

		if targetNumber, err := strconv.Atoi(data.Target); err == nil {
			return validator.ItemComparedToConst(item, targetNumber, allCompare), nil
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
			return validator.ItemsSumComparedToConst(items, targetNumber, allCompare), nil
		}

		return nil, errors.New("failed to recognize target")
	}

	return nil, errors.New("comparator has nothing to compare")
}

func buildCounterValidators(data *yamlValidatorCounter) ([]domain.Validator, error) {
	if data.Number != 0 {
		return validator.CountOfNumber(data.Number, allCount), nil
	}

	if data.Parity != "" {
		parity, err := parseEnum(data.Parity, mappingParity)
		if err != nil {
			return nil, err
		}

		return validator.CountOfParity(parity, allCount), nil
	}

	return nil, errors.New("counter has nothing to count")
}

func buildParityValidators(data *yamlValidatorParity) ([]domain.Validator, error) {
	if len(data.Item) > 0 {
		item, err := parseEnum(data.Item, mappingCodeItem)
		if err != nil {
			return nil, err
		}

		return validator.ItemHasParity(item, allParity), nil
	}

	if data.Sum {
		return validator.SumHasParity(allParity), nil
	}

	return nil, errors.New("parity checker has nothing to check")
}

func buildHasMoreParityValidators() ([]domain.Validator, error) {
	return validator.HasMoreNumbersWithParity(allParity), nil
}

func buildHasRepetitionsValidators() ([]domain.Validator, error) {
	return validator.HasSomeRepeatingNumbers(allCount), nil
}

func buildHasPairValidators() ([]domain.Validator, error) {
	return validator.PairOfNumbersExist(allBool), nil
}

func buildGreatestItemValidators() ([]domain.Validator, error) {
	return validator.OneItemIsDifferent(validator.More, allCodeItem), nil
}

func buildLeastItemValidators() ([]domain.Validator, error) {
	return validator.OneItemIsDifferent(validator.Less, allCodeItem), nil
}

func buildHasOrderValidators() ([]domain.Validator, error) {
	return validator.CodeIsOrdered(allOrder), nil
}
