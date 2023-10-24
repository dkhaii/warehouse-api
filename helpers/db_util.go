package helpers

import (
	"database/sql"
)

type (
	NullSring   sql.NullString
	NullBool    sql.NullBool
	NullInt64   sql.NullInt64
	NullFloat64 sql.NullFloat64
	NullTime    sql.NullTime
)

func NullStringToString(nullStr sql.NullString) string {
	if nullStr.Valid {
		return nullStr.String
	}
	return ""
}

func HandleTXRollBack(tx *sql.Tx) error {
	err := recover()
	if err != nil {
		errRollBack := tx.Rollback()
		if errRollBack != nil {
			return errRollBack
		}
		return err.(error)
	}
	return nil
}

func HandleTXCommit(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func CommitOrRollBack(tx *sql.Tx) error {
	err := recover()
	if err != nil {
		errRollBack := tx.Rollback()
		if errRollBack != nil {
			return errRollBack
		}
		return err.(error)
	}

	errCommit := tx.Commit()
	if errCommit != nil {
		return errCommit
	}

	return nil
}
