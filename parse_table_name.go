package bnrsqlx

import (
	"fmt"
	"reflect"

	"github.com/oneononex/oolib/ooerrors"
)

func parseTableName(arg interface{}) (string, ooerrors.Error) {
	st := reflect.TypeOf(arg)
	tableName := ""
	if obj, ok := arg.(interface{ TableName() string }); ok {
		tableName = obj.TableName()
	} else {
		eMsg := fmt.Sprintf("dao %v: has no TableName function defined", st.Name())
		return "", ooerrors.NewDatabaseErrorWithMessage(eMsg)
	}
	return tableName, nil
}
