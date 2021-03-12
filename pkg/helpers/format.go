package helpers

import (
	"strings"

	"github.com/wilian1808/sqlmql/pkg/models"
)

// FormatData func .
func FormatData(query string) ([]interface{}, error) {
	var arr []interface{}

	new := strings.Split(query, " ")

	for _, v := range new {
		aux := strings.Split(v, ",")

		if len(aux) == 1 {
			au := strings.Join(aux, "")
			switch strings.ToUpper(au) {
			case "SELECT":
				s := models.Reserved{
					Data: strings.TrimSpace(strings.ToLower(au)),
				}
				arr = append(arr, s)
			case "FROM":
				s := models.Reserved{
					Data: strings.TrimSpace(strings.ToLower(au)),
				}
				arr = append(arr, s)
			case "INSERT":
				s := models.Reserved{
					Data: strings.TrimSpace(strings.ToLower(au)),
				}
				arr = append(arr, s)
			case "UPDATE":
				s := models.Reserved{
					Data: strings.TrimSpace(strings.ToLower(au)),
				}
				arr = append(arr, s)
			case "DELETE":
				s := models.Reserved{
					Data: strings.TrimSpace(strings.ToLower(au)),
				}
				arr = append(arr, s)
			case "SET":
				s := models.Reserved{
					Data: strings.TrimSpace(strings.ToLower(au)),
				}
				arr = append(arr, s)
			case "INTO":
				s := models.Reserved{
					Data: strings.TrimSpace(strings.ToLower(au)),
				}
				arr = append(arr, s)
			case "WHERE":
				s := models.Reserved{
					Data: strings.TrimSpace(strings.ToLower(au)),
				}
				arr = append(arr, s)
			default:
				// s := models.Identifier{
				// 	Data: strings.TrimSpace(strings.ToLower(au)),
				// }
				arr = append(arr, strings.TrimSpace(strings.ToLower(au)))
			}

		} else {
			// evaluamos para los que tenian ,
			for _, vaux := range aux {
				if len(vaux) != 0 {
					// s := models.Identifier{
					// 	Data: strings.TrimSpace(strings.ToLower(vaux)),
					// }
					arr = append(arr, strings.TrimSpace(strings.ToLower(vaux)))
				}
			}
		}
	}

	return arr, nil
}
