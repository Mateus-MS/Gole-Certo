package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func GetQueryParam[T any](r *http.Request, key string, required bool, defaultVal T) (T, error) {
	var zero T

	val := r.URL.Query().Get(key)
	if val == "" {
		if required {
			return zero, fmt.Errorf("missing required query parameter: %s", key)
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

func GetProductFilters(r *http.Request, includeBrand, includePrice bool) (filter bson.M, err error) {
	filter = bson.M{}

	if includeBrand {
		// Get the "brand" filter from the URL
		var brandRaw string
		if brandRaw, err = GetQueryParam(r, "brands", false, ""); err != nil {
			return filter, err
		}
		if brandRaw != "" {
			// Assume multiple brands are separated by commas
			brands := strings.Split(brandRaw, ":")
			filter["brand"] = bson.M{
				"$in": brands,
			}
		}
	}

	if includePrice {
		// Get the "price" filter from the URL
		var priceRaw string
		if priceRaw, err = GetQueryParam(r, "price", false, ""); err != nil {
			return filter, err
		}
		if priceRaw != "" {
			priceMinMax := strings.Split(priceRaw, "-")
			if len(priceMinMax) == 2 {
				min, _ := strconv.ParseFloat(priceMinMax[0], 64)
				max, _ := strconv.ParseFloat(priceMinMax[1], 64)
				filter["price"] = bson.M{
					"$gte": min,
					"$lte": max,
				}
			}
		}
	}

	return filter, nil
}
