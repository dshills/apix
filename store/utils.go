package store

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

// AsNullString converts a string into a sql.NullString
func AsNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: true}
}

// AsNullInt64 converts a int64 into a sql.NullInt64
func AsNullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: true}
}

// AsNullTime converts a time into a sql.NullTime
func AsNullTime(t time.Time) mysql.NullTime {
	return mysql.NullTime{Time: t, Valid: true}
}

// NullStringVal converts a sql.NullString to a string if valid
// and to default if not
func NullStringVal(s sql.NullString, def string) string {
	if s.Valid {
		return s.String
	}
	return def
}

// NullInt64Val converts a sql.NullInt64 to a int64 if valid
// and to default if not
func NullInt64Val(i sql.NullInt64, def int64) int64 {
	if i.Valid {
		return i.Int64
	}
	return def
}
