package validate

import (
	"fmt"

	"github.com/wilian1808/sqlmql/pkg/mongo"
	"github.com/wilian1808/sqlmql/pkg/sintaxsql"
)

// RedirectTransform func -
func RedirectTransform(data []interface{}) (string, error) {
	switch data[0] {
	case sintaxsql.SelectSQL:
		return mongo.TransformSelect(data), nil
	case sintaxsql.InsertSQL:
		return mongo.TransformInsert(data), nil
	case sintaxsql.UpdateSQL:
		return mongo.TransformUpdate(data), nil
	case sintaxsql.DeleteSQL:
		return mongo.TransformDelete(data), nil
	default:
		return "", fmt.Errorf("Error en la sintaxis: %s", data[0])
	}
}
