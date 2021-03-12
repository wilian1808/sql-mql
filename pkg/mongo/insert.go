package mongo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wilian1808/sqlmql/pkg/sintaxsql"
)

// TransformInsert func -
func TransformInsert(data []interface{}) string {
	fmt.Println(data...)

	idValues := 0
	idInto := 0

	// posiciones de las key
	for i, v := range data {
		if sintaxsql.IntoSQL == v {
			idInto = i
		} else if sintaxsql.ValuesSQL == v {
			idValues = i
		}
	}

	// obtenemos los datos a insertar con sus valores
	var nextInto []interface{}
	var nextValues []interface{}
	var tableName string
	// ><
	for i, v := range data {
		if i == idInto+1 {
			tableName = fmt.Sprintf("%s", v)
		} else if i > idInto && i < idValues {
			nextInto = append(nextInto, v)
		} else if i > idValues {
			nextValues = append(nextValues, v)
		}
	}

	// evaluamos campos numericos
	var arrIntegers []int
	for i, v := range nextValues {
		str := fmt.Sprintf("%s", v)
		_, err := strconv.Atoi(str)
		if err == nil {
			arrIntegers = append(arrIntegers, i)
		}
	}

	// formateamos los datos a insertar
	var arrGeneral []string
	for i := 0; i < len(nextInto); i++ {
		for _, v := range arrIntegers {
			if v == i {
				v, _ := strconv.Atoi(fmt.Sprintf("%s", nextValues[i]))
				arrGeneral = append(arrGeneral, fmt.Sprintf("%s:%d,", nextInto[i], v))
			} else {
				arrGeneral = append(arrGeneral, fmt.Sprintf("%s:'%s',", nextInto[i], nextValues[i]))
			}
		}
	}

	// convertimos a un solo valor
	str := strings.Join(arrGeneral, " ")
	// formateamos resultado
	result := fmt.Sprintf("db.%s.insert({%s})", tableName, str)
	return result
}
