package parser

import (
	"errors"
	"strconv"
	"turing-machine/internal/domain"
	"turing-machine/internal/domain/validator"
)

func buildValidators(validator yamlValidator) ([]domain.Validator, error) {
	if validator.Compare != nil {
		return buildCompareValidators(validator.Compare, validator.Clues)
	}

	if validator.Counter != nil {
		return buildCounterValidators(validator.Counter, validator.Clues)
	}

	if validator.Parity != nil {
		return buildParityValidators(validator.Parity, validator.Clues)
	}

	if validator.HasMoreParity {
		return buildHasMoreParityValidators(validator.Clues)
	}

	if validator.HasRepetitions {
		return buildHasRepetitionsValidators(validator.Clues)
	}

	if validator.HasPair {
		return buildHasPairValidators(validator.Clues)
	}

	if validator.GreatestItem {
		return buildGreatestItemValidators(validator.Clues)
	}

	if validator.LeastItem {
		return buildLeastItemValidators(validator.Clues)
	}

	if validator.HasOrder {
		return buildHasOrderValidators(validator.Clues)
	}

	return nil, errors.New("validator not specified")
}

func buildCompareValidators(data *yamlValidatorCompare, clues []string) ([]domain.Validator, error) {
	compareVariants, err := parseVariants(clues, mappingCompare)
	if err != nil {
		return nil, err
	}

	if len(data.Item) > 0 {
		item, err := parseEnum(data.Item, mappingCodeItem)
		if err != nil {
			return nil, err
		}

		if targetItem, err := parseEnum(data.Target, mappingCodeItem); err == nil {
			return validator.ItemComparedToOtherItem(item, targetItem, compareVariants), nil
		}

		if targetNumber, err := strconv.Atoi(data.Target); err == nil {
			return validator.ItemComparedToConst(item, targetNumber, compareVariants), nil
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
			return validator.ItemsSumComparedToConst(items, targetNumber, compareVariants), nil
		}

		return nil, errors.New("failed to recognize target")
	}

	return nil, errors.New("comparator has nothing to compare")
}

func buildCounterValidators(data *yamlValidatorCounter, clues []string) ([]domain.Validator, error) {
	countVariants, err := parseVariants(clues, mappingCount)
	if err != nil {
		return nil, err
	}

	if data.Number != 0 {
		return validator.CountOfNumber(data.Number, countVariants), nil
	}

	if data.Parity != "" {
		parity, err := parseEnum(data.Parity, mappingParity)
		if err != nil {
			return nil, err
		}

		return validator.CountOfParity(parity, countVariants), nil
	}

	return nil, errors.New("counter has nothing to count")
}

func buildParityValidators(data *yamlValidatorParity, clues []string) ([]domain.Validator, error) {
	parityVariants, err := parseVariants(clues, mappingParity)
	if err != nil {
		return nil, err
	}

	if len(data.Item) > 0 {
		item, err := parseEnum(data.Item, mappingCodeItem)
		if err != nil {
			return nil, err
		}

		return validator.ItemHasParity(item, parityVariants), nil
	}

	if data.Sum {
		return validator.SumHasParity(parityVariants), nil
	}

	return nil, errors.New("parity checker has nothing to check")
}

func buildHasMoreParityValidators(clues []string) ([]domain.Validator, error) {
	parityVariants, err := parseVariants(clues, mappingParity)
	if err != nil {
		return nil, err
	}

	return validator.HasMoreNumbersWithParity(parityVariants), nil
}

func buildHasRepetitionsValidators(clues []string) ([]domain.Validator, error) {
	countVariants, err := parseVariants(clues, mappingCount)
	if err != nil {
		return nil, err
	}

	return validator.HasSomeRepeatingNumbers(countVariants), nil
}

func buildHasPairValidators(clues []string) ([]domain.Validator, error) {
	boolVariants, err := parseVariants(clues, mappingBool)
	if err != nil {
		return nil, err
	}

	return validator.PairOfNumbersExist(boolVariants), nil
}

func buildGreatestItemValidators(clues []string) ([]domain.Validator, error) {
	codeItemVariants, err := parseVariants(clues, mappingCodeItem)
	if err != nil {
		return nil, err
	}

	return validator.OneItemIsDifferent(validator.More, codeItemVariants), nil
}

func buildLeastItemValidators(clues []string) ([]domain.Validator, error) {
	codeItemVariants, err := parseVariants(clues, mappingCodeItem)
	if err != nil {
		return nil, err
	}

	return validator.OneItemIsDifferent(validator.Less, codeItemVariants), nil
}

func buildHasOrderValidators(clues []string) ([]domain.Validator, error) {
	orderVariants, err := parseVariants(clues, mappingOrder)
	if err != nil {
		return nil, err
	}

	return validator.CodeIsOrdered(orderVariants), nil
}
