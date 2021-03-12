package mongo

import (
	"fmt"
	"strings"

	"github.com/wilian1808/sqlmql/pkg/sintaxsql"
)

// TransformSelect func - evalua la sintaxis para ver si esta correcta
func TransformSelect(data []interface{}) string {
	fmt.Println(data)

	var nextSelect []interface{}
	var nextFrom []interface{}
	var nextWhere []interface{}

	idFrom := 0
	idWhere := len(data) + 1

	// posiciones para las palabras reservadas
	for i, v := range data {
		switch v {
		case sintaxsql.FromSQL:
			idFrom = i
		case sintaxsql.WhereSQL:
			idWhere = i
		}
	}

	// valores de seleccion
	for i, v := range data {
		if i != 0 && i < idFrom {
			nextSelect = append(nextSelect, v)
		} else if i > idFrom && i < idWhere {
			nextFrom = append(nextFrom, v)
		} else if i > idWhere {
			nextWhere = append(nextWhere, v)
		}
	}

	// valores que se consultan
	var valuesQuery []string
	for i, v := range nextSelect {
		if v != "*" {
			if i == len(nextSelect)-1 {
				valuesQuery = append(valuesQuery, fmt.Sprintf(`%s:%d`, v, 1))
			} else {
				valuesQuery = append(valuesQuery, fmt.Sprintf(`%s:%d,`, v, 1))
			}
		}
	}

	// valores de condicion
	var arrWhere []string
	if len(nextWhere) != 0 {
		// formatear el where
		for _, v := range nextWhere {
			switch v {
			case "=":
				arrWhere = append(arrWhere, ":")
			default:
				arrWhere = append(arrWhere, fmt.Sprintf("%s", v))
			}
		}
	}

	var valueWhere string
	if len(nextWhere) != 0 {
		valueWhere = strings.Join(arrWhere, "")
	}

	// withWhere := `db.users.find({ "username " :  'mario' }, {})`
	// select username, paternal, maternal, email from users where data = 'hola'
	str := strings.Join(valuesQuery, " ")
	var newSintax string

	if len(nextWhere) != 0 {
		if len(str) == 0 {
			newSintax = fmt.Sprintf(`db.%s.find({%s},{})`, nextFrom[0], valueWhere)
		} else {
			newSintax = fmt.Sprintf(`db.%s.find({%s},{%s})`, nextFrom[0], valueWhere, str)
		}
	} else {
		if len(str) == 0 {
			newSintax = fmt.Sprintf(`db.%s.find({})`, nextFrom[0])
		} else {
			newSintax = fmt.Sprintf(`db.%s.find({},{%s})`, nextFrom[0], str)
		}
	}

	return newSintax
}
