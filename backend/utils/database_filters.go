package utils

import (
	"math"
	"net/url"
	"strconv"
)

func GetLimitOffset(queryParams url.Values) (int, int, error) {
	limit := 50

	if queryParams.Get("limit") != "" {
		var err error
		limit, err = strconv.Atoi(queryParams.Get("limit"))
		if err != nil {
			return -1, -1, err
		}
		limit = int(math.Min(float64(500), float64(limit)))
		limit = int(math.Max(float64(1), float64(limit)))
	}

	offset := 0

	if queryParams.Get("offset") != "" {
		var err error
		offset, err = strconv.Atoi(queryParams.Get("offset"))
		if err != nil {
			return -1, -1, err
		}
		offset = int(math.Min(float64(9223372036854775807), float64(offset)))
		offset = int(math.Max(float64(0), float64(offset)))
	}

	return limit, offset, nil
}

func GenerateWhereFilter(filters []string, queryParams url.Values) (map[string]interface{}, error) {
	whereClause := make(map[string]interface{})

	for _, filter := range filters {
		if queryParams.Get(filter) != "" {
			var err error
			whereClause[filter] = queryParams.Get(filter)
			if err != nil {
				return nil, err
			}
		}
	}

	return whereClause, nil
}
