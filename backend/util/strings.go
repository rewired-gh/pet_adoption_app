package util

import "database/sql"

func ToNullString(s string) sql.NullString {
	return sql.NullString{
		Valid:  s != "",
		String: s,
	}
}
