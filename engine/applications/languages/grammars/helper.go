package grammars

import (
	"bytes"
	"errors"
	"fmt"
)

func bytesToRuleNameAndValue(
	data []byte,
	ruleNameValueSeparator byte,
	possibleNameCharacters []byte,
	ruleNameSeparator byte,
	ruleValuePrefix byte,
	ruleValueSuffix byte,
	ruleValueEscape byte,
) ([]byte, []byte, []byte, error) {
	retRuleName, retRemaining, err := bytesToRuleName(data, possibleNameCharacters, ruleNameSeparator)
	if err != nil {
		return nil, nil, nil, err
	}

	if len(retRemaining) < 1 {
		return nil, nil, nil, errors.New("the remaining data after fetching the ruleName was expected to contain at least 1 byte")
	}

	if retRemaining[0] != ruleNameValueSeparator {
		str := fmt.Sprintf("the byte after fetching the ruleName was expected to be %d, %d provided", ruleNameValueSeparator, retRemaining[0])
		return nil, nil, nil, errors.New(str)
	}

	retRuleValue, retRemainingAfterValue, err := extractBetween(retRemaining[1:], ruleValuePrefix, ruleValueSuffix, ruleValueEscape)
	if err != nil {
		return nil, nil, nil, err
	}

	return retRuleName, retRuleValue, retRemainingAfterValue, nil
}

func bytesToRuleName(data []byte, possibleBytes []byte, separator byte) ([]byte, []byte, error) {
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

	remaining := data[len(output):]
	sepBytes := []byte{separator}
	if bytes.HasSuffix(output, sepBytes) {
		output = output[:len(output)-1]
		remaining = append(sepBytes, remaining...)
	}

	return output, remaining, nil
}

func extractBetween(data []byte, prefix byte, suffix byte, escape byte) ([]byte, []byte, error) {
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
	lastByteInInt := -1
	isSuffixReached := false
	lastIndex := 1
	for _, oneByte := range data[1:] {
		lastIndex++
		if !escapeReached && lastByteInInt != -1 {
			escapeReached = oneByte == escape
			if escapeReached {
				continue
			}
		}

		if escapeReached {
			output = append(output, oneByte)
			lastByteInInt = int(oneByte)
			escapeReached = false
			continue
		}

		if oneByte == suffix {
			isSuffixReached = true
			break
		}

		output = append(output, oneByte)
		lastByteInInt = int(oneByte)
	}

	if !isSuffixReached {
		str := fmt.Sprintf("the suffix byte (%d) was never reached", suffix)
		return nil, nil, errors.New(str)
	}

	return output, data[lastIndex:], nil
}

func createPossibleRuleNameCharactersList() []byte {
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
