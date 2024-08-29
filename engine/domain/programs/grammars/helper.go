package grammars

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

func blockName(
	data []byte,
	firstBytes []byte,
	remainingAcceptableBytes []byte,
	filterBytes []byte,
) ([]byte, []byte, error) {
	data = filterPrefix(data, filterBytes)
	retFirstMatches, retRemaining := matchBytes(data, firstBytes, filterBytes)
	if len(retFirstMatches) <= 0 {
		return nil, nil, errors.New("the bytes did not match any of the firstBytes")
	}

	retSecondMatches, retSecondRemaining := matchBytes(retRemaining, remainingAcceptableBytes, filterBytes)
	return append(retFirstMatches, retSecondMatches...), retSecondRemaining, nil
}

func bytesToMinMax(
	data []byte,
	possibleNumbers []byte,
	cardinalityOpen byte,
	cardinalityClose byte,
	cardinalitySeparator byte,
	cardinalityZeroPlus byte,
	cardinalityOnePlus byte,
	cardinalityOptional byte,
	filterBytes []byte,
) (uint, *uint, []byte, error) {
	data = filterPrefix(data, filterBytes)
	if len(data) <= 0 {
		str := fmt.Sprintf("the bytes must contain at least 1 value in order to convert it to cardinality's min/max, %d provided", len(data))
		return 0, nil, nil, errors.New(str)
	}

	firstValue := data[0]
	if firstValue == cardinalityOnePlus {
		return 1, nil, data[1:], nil
	}

	if firstValue == cardinalityZeroPlus {
		return 0, nil, data[1:], nil
	}

	if firstValue == cardinalityOptional {
		max := uint(1)
		return 0, &max, data[1:], nil
	}

	return bytesToBracketsMinMax(
		data,
		possibleNumbers,
		cardinalityOpen,
		cardinalityClose,
		cardinalitySeparator,
		filterBytes,
	)
}

func bytesToBracketsIndex(
	data []byte,
	possibleNumbers []byte,
	indexOpen byte,
	indexClose byte,
	filterBytes []byte,
) (uint, []byte, error) {
	data = filterPrefix(data, filterBytes)
	if len(data) <= 0 {
		str := fmt.Sprintf("the bytes must contain at least 1 value in order to convert it to an index, %d provided", len(data))
		return 0, nil, errors.New(str)
	}

	data = filterPrefix(data, filterBytes)
	firstValue := data[0]
	if firstValue != indexOpen {
		return 0, nil, errors.New("the provided bytes could not be converted to an index because it does not contain the indexOpen byte")
	}

	data = filterPrefix(data[1:], filterBytes)
	retMinBytes, retRemaining := matchBytes(data, possibleNumbers, filterBytes)
	iValue, err := strconv.Atoi(string(retMinBytes))
	if err != nil {
		return 0, nil, err
	}

	uiValue := uint(iValue)
	if len(retRemaining) <= 0 {
		return 0, nil, errors.New("the remaining bytes, after fetching the minimum, was empty and therefore could not be converted to cardinality's min/max")
	}

	nextValue := retRemaining[0]
	if nextValue != indexClose {
		return 0, nil, errors.New("the provided bytes could not be converted to an index because it does not contain the indexClose byte")
	}

	return uiValue, filterPrefix(retRemaining[1:], filterBytes), nil
}

func bytesToBracketsMinMax(
	data []byte,
	possibleNumbers []byte,
	cardinalityOpen byte,
	cardinalityClose byte,
	cardinalitySeparator byte,
	filterBytes []byte,
) (uint, *uint, []byte, error) {
	data = filterPrefix(data, filterBytes)
	if len(data) <= 0 {
		str := fmt.Sprintf("the bytes must contain at least 1 value in order to convert it to cardinality's min/max, %d provided", len(data))
		return 0, nil, nil, errors.New(str)
	}

	data = filterPrefix(data, filterBytes)
	firstValue := data[0]
	if firstValue != cardinalityOpen {
		return 0, nil, nil, errors.New("the provided bytes could not be converted to cardinality's min/max")
	}

	data = filterPrefix(data[1:], filterBytes)
	retMinBytes, retRemaining := matchBytes(data, possibleNumbers, filterBytes)
	iMin, err := strconv.Atoi(string(retMinBytes))
	if err != nil {
		return 0, nil, nil, err
	}

	uiMin := uint(iMin)
	if len(retRemaining) <= 0 {
		return 0, nil, nil, errors.New("the remaining bytes, after fetching the minimum, was empty and therefore could not be converted to cardinality's min/max")
	}

	nextValue := retRemaining[0]
	if nextValue == cardinalityClose {
		// max same as min:
		return uiMin, &uiMin, retRemaining[1:], nil
	}

	if nextValue != cardinalitySeparator {
		return 0, nil, nil, errors.New("the provided bytes could not be converted to cardinality's min/max, no separator found")
	}

	retRemaining = filterPrefix(retRemaining[1:], filterBytes)
	if len(retRemaining) <= 0 {
		str := fmt.Sprintf("the remaining bytes, after fetching the cardinality separator (%s), was expected to not be empty", string([]byte{cardinalityClose}))
		return 0, nil, nil, errors.New(str)
	}

	nextValueAfterSeparator := retRemaining[0]
	if nextValueAfterSeparator == cardinalityClose {
		// min, no max
		return uiMin, nil, retRemaining[1:], nil
	}

	retRemaining = filterPrefix(retRemaining, filterBytes)
	retMaxBytes, retRemainingAfterMax := matchBytes(retRemaining, possibleNumbers, filterBytes)
	if len(retRemainingAfterMax) <= 0 {
		str := fmt.Sprintf("the remaining bytes, after fetching the cardinality's max (%s), was expected to contain the cardinality's close byte (%s).  Emty bytes returned", retMaxBytes, string([]byte{cardinalityClose}))
		return 0, nil, nil, errors.New(str)
	}

	if retRemainingAfterMax[0] != cardinalityClose {
		str := fmt.Sprintf("the remaining bytes, after fetching the cardinality's max, was expected to contain the close byte (%s)", string([]byte{cardinalityClose}))
		return 0, nil, nil, errors.New(str)
	}

	iMax, err := strconv.Atoi(string(retMaxBytes))
	if err != nil {
		return 0, nil, nil, err
	}

	uiMax := uint(iMax)
	return uiMin, &uiMax, filterPrefix(retRemainingAfterMax[1:], filterBytes), nil
}

func bytesToRuleNameAndValue(
	data []byte,
	ruleNameValueSeparator byte,
	possibleNameCharacters []byte,
	ruleNameSeparator byte,
	ruleValuePrefix byte,
	ruleValueSuffix byte,
	ruleValueEscape byte,
	filterBytes []byte,
) ([]byte, []byte, []byte, error) {
	retRuleName, retRemaining, err := bytesToRuleName(data, possibleNameCharacters, ruleNameSeparator, filterBytes)
	if err != nil {
		return nil, nil, nil, err
	}

	if len(retRemaining) < 1 {
		return nil, nil, nil, errors.New("the remaining data after fetching the ruleName was expected to contain at least 1 byte")
	}

	retRemaining = filterPrefix(retRemaining, filterBytes)
	if retRemaining[0] != ruleNameValueSeparator {
		str := fmt.Sprintf("the byte after fetching the ruleName was expected to be %d, %d provided", ruleNameValueSeparator, retRemaining[0])
		return nil, nil, nil, errors.New(str)
	}

	retRemaining = filterPrefix(retRemaining[1:], filterBytes)
	retRuleValue, retRemainingAfterValue, err := extractBetween(retRemaining, ruleValuePrefix, ruleValueSuffix, &ruleValueEscape)
	if err != nil {
		return nil, nil, nil, err
	}

	return retRuleName, retRuleValue, filterPrefix(retRemainingAfterValue, filterBytes), nil
}

func bytesToRuleName(data []byte, possibleBytes []byte, separator byte, filter []byte) ([]byte, []byte, error) {
	data = filterPrefix(data, filter)
	output := []byte{}
	for idx, oneByte := range data {
		isValid := false
		if oneByte == separator {
			if idx == 0 {
				str := fmt.Sprintf("the first character of a ruleName cannot be the separator (%s)", string(separator))
				return nil, nil, errors.New(str)
			}

			isValid = true
		}

		if !isValid {
			for _, onePossibleBytes := range possibleBytes {
				if oneByte == onePossibleBytes {
					isValid = true
					break
				}
			}
		}

		if isValid {
			output = append(output, oneByte)
			continue
		}

		break
	}

	remaining := filterPrefix(data[len(output):], filter)
	sepBytes := []byte{separator}
	if bytes.HasSuffix(output, sepBytes) {
		output = output[:len(output)-1]
		remaining = append(sepBytes, remaining...)
	}

	if len(output) <= 0 {
		return nil, nil, errors.New("the rule name must contain at least 1 byte")
	}

	return output, filterPrefix(remaining, filter), nil
}

func extractBetween(data []byte, prefix byte, suffix byte, pEscape *byte) ([]byte, []byte, error) {
	if len(data) < 2 {
		str := fmt.Sprintf("the input was expected to contain at least 2 bytes, %d provided", len(data))
		return nil, nil, errors.New(str)
	}

	if data[0] != prefix {
		str := fmt.Sprintf("the first byte of the input (%d) was expected to contain the prefix (%d)", data[0], prefix)
		return nil, nil, errors.New(str)
	}

	output := []byte{}
	escapeReached := false
	isSuffixReached := false
	lastIndex := 1
	for _, oneByte := range data[1:] {
		lastIndex++
		if pEscape != nil && !escapeReached {
			escapeReached = oneByte == *pEscape
			if escapeReached {
				continue
			}
		}

		if escapeReached {
			output = append(output, oneByte)
			escapeReached = false
			continue
		}

		if oneByte == suffix {
			isSuffixReached = true
			break
		}

		output = append(output, oneByte)
	}

	if !isSuffixReached {
		str := fmt.Sprintf("the suffix byte (%d) was never reached", suffix)
		return nil, nil, errors.New(str)
	}

	return output, data[lastIndex:], nil
}

func matchBytes(data []byte, possibleValues []byte, filterBytes []byte) ([]byte, []byte) {
	output := []byte{}
	for _, oneByte := range data {
		isMatch := false
		for _, onePossibleValue := range possibleValues {
			if onePossibleValue == oneByte {
				isMatch = true
				break
			}
		}

		if isMatch {
			output = append(output, oneByte)
			continue
		}

		break
	}

	return output, filterPrefix(data[len(output):], filterBytes)
}

func filterPrefix(data []byte, possibleBytes []byte) []byte {
	if len(data) <= 0 {
		return data
	}

	first := data[0]
	for _, oneByte := range possibleBytes {
		if first == oneByte {
			return filterPrefix(data[1:], possibleBytes)
		}
	}

	return data
}

func createPossibleFuncNameCharacters() []byte {
	letters := createBlockNameCharacters()
	output := append(letters, []byte(funcNameSeparator)...)
	return append(output, createPossibleNumbers()...)
}

func createBlockNameCharacters() []byte {
	numbers := createPossibleNumbers()
	lowerCaseLetters := createPossibleLowerCaseLetters()
	upperCaseLetters := createPossibleUpperCaseLetters()
	output := append(lowerCaseLetters, upperCaseLetters...)
	output = append(output, numbers...)
	return output
}

func createPossibleUpperCaseLetters() []byte {
	return []byte{
		[]byte(ulA)[0],
		[]byte(ulB)[0],
		[]byte(ulC)[0],
		[]byte(ulD)[0],
		[]byte(ulE)[0],
		[]byte(ulF)[0],
		[]byte(ulG)[0],
		[]byte(ulH)[0],
		[]byte(ulI)[0],
		[]byte(ulJ)[0],
		[]byte(ulK)[0],
		[]byte(ulL)[0],
		[]byte(ulM)[0],
		[]byte(ulN)[0],
		[]byte(ulO)[0],
		[]byte(ulP)[0],
		[]byte(ulQ)[0],
		[]byte(ulR)[0],
		[]byte(ulS)[0],
		[]byte(ulT)[0],
		[]byte(ulU)[0],
		[]byte(ulV)[0],
		[]byte(ulW)[0],
		[]byte(ulX)[0],
		[]byte(ulY)[0],
		[]byte(ulZ)[0],
	}
}

func createPossibleLowerCaseLetters() []byte {
	return []byte{
		[]byte(llA)[0],
		[]byte(llB)[0],
		[]byte(llC)[0],
		[]byte(llD)[0],
		[]byte(llE)[0],
		[]byte(llF)[0],
		[]byte(llG)[0],
		[]byte(llH)[0],
		[]byte(llI)[0],
		[]byte(llJ)[0],
		[]byte(llK)[0],
		[]byte(llL)[0],
		[]byte(llM)[0],
		[]byte(llN)[0],
		[]byte(llO)[0],
		[]byte(llP)[0],
		[]byte(llQ)[0],
		[]byte(llR)[0],
		[]byte(llS)[0],
		[]byte(llT)[0],
		[]byte(llU)[0],
		[]byte(llV)[0],
		[]byte(llW)[0],
		[]byte(llX)[0],
		[]byte(llY)[0],
		[]byte(llZ)[0],
	}
}

func createPossibleNumbers() []byte {
	return []byte{
		[]byte(nZero)[0],
		[]byte(nOne)[0],
		[]byte(nTwo)[0],
		[]byte(nTree)[0],
		[]byte(nFour)[0],
		[]byte(nFive)[0],
		[]byte(nSix)[0],
		[]byte(nSeven)[0],
		[]byte(nHeight)[0],
		[]byte(nNine)[0],
	}
}
