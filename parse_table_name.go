package bnrsqlx

import (
	"errors"
	"fmt"
	"reflect"
)

func parseTableName(arg interface{}) (string, error) {
	st := reflect.TypeOf(arg)
	tableName := ""
	if obj, ok := arg.(interface{ TableName() string }); ok {
		tableName = obj.TableName()
	} else {
		eMsg := fmt.Sprintf("dao %v: has no TableName function defined", st.Name())
		return "", errors.New(eMsg)
	}
	return tableName, nil
}
