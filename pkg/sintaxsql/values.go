package sintaxsql

import "github.com/wilian1808/sqlmql/pkg/models"

var (
	// SelectSQL - select
	SelectSQL = models.Reserved{Data: "select"}
	// InsertSQL - insert
	InsertSQL = models.Reserved{Data: "insert"}
	// UpdateSQL - update
	UpdateSQL = models.Reserved{Data: "update"}
	// DeleteSQL - delete
	DeleteSQL = models.Reserved{Data: "delete"}
	// WhereSQL - where
	WhereSQL = models.Reserved{Data: "where"}
	// FromSQL - from
	FromSQL = models.Reserved{Data: "from"}
)
