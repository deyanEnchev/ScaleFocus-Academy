// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package sqlcDB

import (
	"database/sql"
)

type TopStory struct {
	ID         int32
	Title      sql.NullString
	Score      sql.NullInt32
	TimeStored sql.NullString
}
