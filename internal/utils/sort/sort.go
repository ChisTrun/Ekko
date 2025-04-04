package utils

import (
	ekko "ekko/api"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
)

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func GetSort(columns []string, table string, sorts []*ekko.SortMethod) (func(s *sql.Selector), error) {
	var values []string
	for _, data := range sorts {
		if Contains(columns, data.Name) {
			switch data.Type {
			case ekko.SortType_SORT_TYPE_ASC:
				values = append(values, sql.Asc(fmt.Sprintf("`%s`.`%s`", table, data.Name)))
			case ekko.SortType_SORT_TYPE_DESC:
				values = append(values, sql.Desc(fmt.Sprintf("`%s`.`%s`", table, data.Name)))
			}
			continue
		}
		return nil, errors.New("column not found")
	}
	return func(s *sql.Selector) {
		s.OrderBy(values...)
	}, nil
}
