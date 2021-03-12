package mongo

import (
	"fmt"
	"strings"

	"github.com/wilian1808/sqlmql/pkg/sintaxsql"
)

// transformDelete func -
func TransformDelete(data []interface{}) string {
	fmt.Println(data...)
	idFrom := 0
	idWhere := len(data)

	for i, v := range data {
		switch v {
		case sintaxsql.FromSQL:
			idFrom = i
		case sintaxsql.WhereSQL:
			idWhere = i
		}
	}

	// delete from users where eeee = eee
	// obtenemos valores <>
	var valueTable string
	var valuesWhere []string
	count := 0
	for i, v := range data {
		if i > idFrom && i < idWhere {
			valueTable = fmt.Sprintf("%s", v)
		} else if i > idWhere {
			count = count + 1
			if count == 3 {
				valuesWhere = append(valuesWhere, fmt.Sprintf("'%s'", v))
				count = 0
			} else {
				valuesWhere = append(valuesWhere, fmt.Sprintf("%s", v))
			}
		}
	}

	// formateamos la condicion
	valWhere := strings.ReplaceAll(strings.Join(valuesWhere, ""), "=", ":")
	var response string

	// db.users.deleteMany({ username: "mario" })
	if len(valuesWhere) != 0 {
		response = fmt.Sprintf(`db.%s.deleteMany({%s})`, valueTable, valWhere)
	} else {
		response = fmt.Sprintf(`db.%s.deleteMany({})`, valueTable)
	}

	return response
}
