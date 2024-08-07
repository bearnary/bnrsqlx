package bnrsqlx

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
)

type UpdateHelper struct {
	params      map[string]interface{}
	whereParams map[string]interface{}
	db          Client
	paramCount  int
}

func (c *defaultClient) NewUpdateHelper() *UpdateHelper {
	return &UpdateHelper{
		params:      make(map[string]interface{}),
		whereParams: make(map[string]interface{}),
		db:          c,
	}
}

func (h *UpdateHelper) SetWhereParam(key string, value interface{}) {
	h.whereParams[key] = value
}

func (h *UpdateHelper) SetParam(key string, value interface{}) {
	h.params[key] = value
}

func (h *UpdateHelper) CommitUpdateQuery(tableName string) error {

	h.paramCount = 1
	var args []interface{}

	if !h.HasChanged() {
		return nil
	}

	updateQuery := ""
	isNotFirstUpdateClause := false
	for k, v := range h.params {
		if isNotFirstUpdateClause {
			updateQuery = fmt.Sprintf("%v%v", updateQuery, ", ")
		} else {
			isNotFirstUpdateClause = true
		}
		av := h.ParamArgumentVariable()
		updateQuery = fmt.Sprintf("%v%v = %v", updateQuery, k, av)

		args = append(args, v)
	}

	whereQuery := ""
	if len(h.whereParams) > 0 {
		isNotFirstWhereClause := false
		for k, v := range h.whereParams {
			if isNotFirstWhereClause {
				whereQuery = fmt.Sprintf("%v and ", whereQuery)
			} else {
				isNotFirstWhereClause = true
			}
			av := h.ParamArgumentVariable()
			whereQuery = fmt.Sprintf("%v%v = %v", whereQuery, k, av)

			args = append(args, v)
		}
	} else {
		return errors.New("no where clause in database query")
	}

	query := fmt.Sprintf("update %v set %v where %v", tableName, updateQuery, whereQuery)
	res, err := h.db.DB().Exec(query, args...)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (h *UpdateHelper) CommitUpdateInQuery(tableName string) error {

	h.paramCount = 1
	var args []interface{}

	if !h.HasChanged() {
		return nil
	}

	whereQuery := []string{}
	if len(h.whereParams) > 0 {
		for key, value := range h.whereParams {
			if isArrayOrSlice(value) {
				inQuery, inArgs, err := sqlx.In(fmt.Sprintf("%v in (?)", key), value)
				if err != nil {
					return errors.New("no where clause in database query")
				}
				whereQuery = append(whereQuery, inQuery)
				args = append(args, inArgs...)
				h.paramCount += len(inArgs)
			} else {
				whereQuery = append(whereQuery, fmt.Sprintf("%v = $%d", key, h.paramCount))
				args = append(args, value)
				h.paramCount++
			}
		}
	} else {
		return errors.New("no where clause in database query")
	}

	updateQuery := []string{}
	//isNotFirstUpdateClause := false
	for key, value := range h.params {
		updateQuery = append(updateQuery, fmt.Sprintf("%v = $%d", key, h.paramCount))
		args = append(args, value)
		h.paramCount++
	}

	whereSQL := strings.Join(whereQuery, " AND ")
	updateSQL := strings.Join(updateQuery, ", ")

	query := fmt.Sprintf("update %v set %v where %v", tableName, updateSQL, whereSQL)

	query = h.db.DB().Rebind(query)

	res, err := h.db.DB().Exec(query, args...)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (h *UpdateHelper) ParamArgumentVariable() string {
	if h.db.DatabaseType() == DatabaseTypePostgres {
		v := fmt.Sprintf("$%d", h.paramCount)
		h.paramCount++
		return v
	} else {
		return "?"
	}
}

func (h *UpdateHelper) HasChanged() bool {
	return len(h.params) > 0
}

func isArrayOrSlice(input interface{}) bool {
	typeOf := reflect.TypeOf(input)
	if typeOf == nil {
		return false
	}
	return typeOf.Kind() == reflect.Array || typeOf.Kind() == reflect.Slice
}
