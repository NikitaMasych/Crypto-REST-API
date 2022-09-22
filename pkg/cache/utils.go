package cache

import (
	"GenesisTask/pkg/errors"
	"reflect"
)

func convertInterfaceToFloat64(unk interface{}) (float64, error) {
	float64Type := reflect.TypeOf(float64(0))
	valuePtr := reflect.ValueOf(unk)
	value := reflect.Indirect(valuePtr)
	if !value.Type().ConvertibleTo(float64Type) {
		return 0, errors.ErrIsNotFloat64
	}
	result := value.Convert(float64Type)
	return result.Float(), nil
}
