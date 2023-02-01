package utils

import (
	"net/url"
)

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
