package bnrsqlx

import (
	"fmt"

	"github.com/oneononex/oolib/ooerrors"
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

func (h *UpdateHelper) CommitUpdateQuery(tableName string) ooerrors.Error {

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
		return ooerrors.ErrNoWhereClauseInDatabaseQuery
	}

	query := fmt.Sprintf("update %v set %v where %v", tableName, updateQuery, whereQuery)
	res, err := h.db.DB().Exec(query, args...)
	if err != nil {
		return ooerrors.NewDatabaseError(err)
	}
	_, err = res.RowsAffected()
	if err != nil {
		return ooerrors.NewDatabaseError(err)
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
