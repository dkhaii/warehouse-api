package helpers

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"time"
)

type (
	NullSring   sql.NullString
	NullBool    sql.NullBool
	NullInt64   sql.NullInt64
	NullFloat64 sql.NullFloat64
	NullTime    sql.NullTime
)

func (ns *NullSring) Scan(value interface{}) error {
	var str sql.NullString

	err := str.Scan(value)
	if err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*ns = NullSring{
			str.String,
			false,
		}
	} else {
		*ns = NullSring{
			str.String,
			true,
		}
	}

	return nil
}

func (ns *NullSring) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}

	jm, err := json.Marshal(ns.String)
	if err != nil {
		return nil, err
	}

	return jm, nil
}

func (ns *NullSring) UnMarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	if err != nil {
		ns.Valid = false
		return err
	}
	ns.Valid = true

	return err
}

func (ni *NullInt64) Scan(value interface{}) error {
	var i sql.NullInt64

	err := i.Scan(value)
	if err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*ni = NullInt64{
			i.Int64,
			false,
		}
	} else {
		*ni = NullInt64{
			i.Int64,
			true,
		}
	}

	return nil
}

func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(ni.Int64)
}

func (ni *NullInt64) UnMarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int64)
	if err != nil {
		ni.Valid = false
		return err
	}
	ni.Valid = true
	return err
}

func (nb *NullBool) Scan(value interface{}) error {
	var b sql.NullBool

	err := b.Scan(value)
	if err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*nb = NullBool{
			b.Bool,
			false,
		}
	} else {
		*nb = NullBool{
			b.Bool,
			true,
		}
	}

	return nil
}

func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nb.Bool)
}

func (nb *NullBool) UnMarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nb.Bool)
	if err != nil {
		nb.Valid = false
		return err
	}
	nb.Valid = true
	return err
}

func (nf *NullFloat64) Scan(value interface{}) error {
	var f sql.NullFloat64

	err := f.Scan(value)
	if err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*nf = NullFloat64{
			f.Float64,
			false,
		}
	} else {
		*nf = NullFloat64{
			f.Float64,
			true,
		}
	}

	return nil
}

func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nf.Float64)
}

func (nf *NullFloat64) UnMarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nf.Float64)
	if err != nil {
		nf.Valid = false
		return err
	}
	nf.Valid = true

	return nil
}

func (nt *NullTime) Scan(value interface{}) error {
	var t sql.NullTime

	err := t.Scan(value)
	if err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*nt = NullTime{
			t.Time,
			false,
		}
	} else {
		*nt = NullTime{
			t.Time,
			true,
		}
	}

	return nil
}

func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}

	val := nt.Time.Format(time.RFC3339)
	return []byte(val), nil
}

func (nt *NullTime) UnMarshalJSON(b []byte) error {
	s := string(b)

	x, err := time.Parse(time.RFC3339, s)
	if err != nil {
		nt.Valid = false
		return err
	}

	nt.Time = x
	nt.Valid = true

	return nil
}

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
