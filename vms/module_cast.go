package vms

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/steve-care-software/webx/programs/domain/programs/modules"
)

type moduleCast struct {
}

func createModuleCast() *moduleCast {
	out := moduleCast{}
	return &out
}

// Execute executes the application
func (app *moduleCast) Execute() map[uint]modules.ExecuteFn {
	return app.castTo()
}

func (app *moduleCast) castTo() map[uint]modules.ExecuteFn {
	toInt := app.castToInt()
	toUint := app.castToUint()
	toBool := app.castToBool()
	toFloat32 := app.castToFloat32()
	toFloat64 := app.castToFloat64()
	return map[uint]modules.ExecuteFn{
		ModuleCastToInt:     toInt,
		ModuleCastToUint:    toUint,
		ModuleCastToBool:    toBool,
		ModuleCastToFloat32: toFloat32,
		ModuleCastToFloat64: toFloat64,
	}
}

func (app *moduleCast) castToInt() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		if ins, ok := input[0]; ok {
			if casted, ok := ins.(string); ok {
				return strconv.Atoi(casted)
			}

			if casted, ok := ins.(uint); ok {
				return int(casted), nil
			}

			str := fmt.Sprintf("the value was expected to contain a string or uint")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the value was expected to be valid")
		return nil, errors.New(str)
	}
}

func (app *moduleCast) castToUint() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		if ins, ok := input[0]; ok {
			str := strings.TrimSpace(fmt.Sprintf("%s", ins))
			intValue, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}

			return uint(intValue), nil
		}

		str := fmt.Sprintf("the value was expected to be valid")
		return nil, errors.New(str)
	}
}

func (app *moduleCast) castToBool() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		if ins, ok := input[0]; ok {
			if casted, ok := ins.(string); ok {
				if strings.TrimSpace(casted) == "true" {
					return true, nil
				}

				if strings.TrimSpace(casted) == "false" {
					return false, nil
				}

				str := fmt.Sprintf("the value was expected to contain true/false when a string is provided")
				return nil, errors.New(str)
			}

			if casted, ok := ins.(int); ok {
				if casted == 0 {
					return false, nil
				}

				return true, nil
			}

			if casted, ok := ins.(uint); ok {
				if casted == 0 {
					return false, nil
				}

				return true, nil
			}

			str := fmt.Sprintf("the value was expected to contain a string, int or uint")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the value was expected to be valid")
		return nil, errors.New(str)
	}
}

func (app *moduleCast) castToFloat32() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		if ins, ok := input[0]; ok {
			if casted, ok := ins.(string); ok {
				floatSixtyFour, err := strconv.ParseFloat(casted, 32)
				if err != nil {
					return nil, err
				}

				return float32(floatSixtyFour), nil
			}

			if casted, ok := ins.(int); ok {
				return float32(casted), nil
			}

			if casted, ok := ins.(uint); ok {
				return float32(casted), nil
			}

			str := fmt.Sprintf("the value was expected to contain a string, int or uint")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the value was expected to be valid")
		return nil, errors.New(str)
	}
}

func (app *moduleCast) castToFloat64() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		if ins, ok := input[0]; ok {
			if casted, ok := ins.(string); ok {
				return strconv.ParseFloat(casted, 64)
			}

			if casted, ok := ins.(int); ok {
				return float64(casted), nil
			}

			if casted, ok := ins.(uint); ok {
				return float64(casted), nil
			}

			str := fmt.Sprintf("the value was expected to contain a string, int or uint")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the value was expected to be valid")
		return nil, errors.New(str)
	}
}
