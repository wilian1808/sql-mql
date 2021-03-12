package mongo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wilian1808/sqlmql/pkg/sintaxsql"
)

// TransformarUpdate func -
func TransformUpdate(data []interface{}) string {
	fmt.Println(data...)

	// traemos los indices
	idSet := 0
	idWhere := len(data)

	for i, v := range data {
		switch v {
		case sintaxsql.SetSQL:
			idSet = i
		case sintaxsql.WhereSQL:
			idWhere = i
		}
	}

	// obtenemos los valores
	var valueTable string
	var nextSet []interface{}
	var nextWhere []interface{}
	for i, v := range data {
		if i == 1 {
			valueTable = fmt.Sprintf("%s", v)
		} else if i > idSet && i < idWhere {
			nextSet = append(nextSet, v)
		} else if i > idWhere {
			nextWhere = append(nextWhere, v)
		}
	}

	// obtenemos los campos numericos
	var arrIntegers []int
	for i, v := range nextSet {
		str := fmt.Sprintf("%s", v)
		_, err := strconv.Atoi(str)
		if err == nil {
			arrIntegers = append(arrIntegers, i)
		}
	}

	// formateamos el dato - {"age": 55} -next set
	count := 0
	var formatSet []string
	for i, v := range nextSet {
		count = count + 1
		if count == 3 && i != len(nextSet)-1 {
			formatSet = append(formatSet, fmt.Sprintf("'%s',", v))
			count = 0
		} else {
			formatSet = append(formatSet, fmt.Sprintf("%s", v))
		}
	}

	// convertimos los campos numericos
	for _, v := range arrIntegers {
		val, _ := strconv.Atoi(fmt.Sprintf("%s", nextSet[v]))
		formatSet[v] = fmt.Sprintf("%d", val)
	}

	// datos finales a actualizar
	updateValues := strings.ReplaceAll(strings.Join(formatSet, ""), "=", ":")

	// formateamos el where
	var valuesWhere []string
	count = 0
	for _, v := range nextWhere {
		count = count + 1
		if count == 3 {
			valuesWhere = append(valuesWhere, fmt.Sprintf("'%s'", v))
			count = 0
		} else {
			valuesWhere = append(valuesWhere, fmt.Sprintf("%s", v))
		}
	}

	valWhere := strings.ReplaceAll(strings.Join(valuesWhere, ""), "=", ":")

	// guia - db.users.update({"username": "pedro"}, {$set: {"age": 55}})
	var response string
	// formateamos la respuesta
	if len(nextWhere) != 0 {
		response = fmt.Sprintf(`db.%s.updateMany({%s}, {$set: {%s}})`, valueTable, valWhere, updateValues)
	} else {
		response = fmt.Sprintf(`db.%s.updateMany({}, {$set: {%s}})`, valueTable, updateValues)
	}

	return response
}
