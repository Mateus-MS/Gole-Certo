package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var ParameterNotInURL = errors.New("missing required query parameter")

func GetQueryParam[T any](r *http.Request, key string, required bool, defaultVal T) (T, error) {
	var zero T

	val := r.URL.Query().Get(key)
	if val == "" {
		if required {
			return zero, ParameterNotInURL
		}
		return defaultVal, nil
	}

	switch any(zero).(type) {
	case string:
		return any(val).(T), nil
	case int:
		i, err := strconv.Atoi(val)
		if err != nil {
			return zero, fmt.Errorf("invalid int value for %s: %v", key, err)
		}
		return any(i).(T), nil
	case float64:
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return zero, fmt.Errorf("invalid float value for %s: %v", key, err)
		}
		return any(f).(T), nil
	case bool:
		b, err := strconv.ParseBool(val)
		if err != nil {
			return zero, fmt.Errorf("invalid bool value for %s: %v", key, err)
		}
		return any(b).(T), nil
	default:
		return zero, fmt.Errorf("unsupported type")
	}
}
