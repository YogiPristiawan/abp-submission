package primitive

import (
	"database/sql"
	"database/sql/driver"
	"reflect"

	"github.com/goccy/go-json"
)

type Bool sql.NullBool

func (b *Bool) Scan(value interface{}) error {
	var nb sql.NullBool
	if err := nb.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*b = Bool{nb.Bool, false}
	} else {
		*b = Bool{nb.Bool, true}
	}
	return nil
}

func (b Bool) Value() (driver.Value, error) {
	if !b.Valid {
		return nil, nil
	}
	return b.Bool, nil
}

func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(b.Bool)
}

func (b *Bool) UnmarshalJSON(by []byte) error {
	if str := string(by); str == `null` {
		return nil
	}
	err := json.Unmarshal(by, &b.Bool)
	b.Valid = (err == nil)
	return err
}

type Float64 sql.NullFloat64

func (f *Float64) Scan(value interface{}) error {
	var nf sql.NullFloat64
	if err := nf.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*f = Float64{nf.Float64, false}
	} else {
		*f = Float64{nf.Float64, true}
	}
	return nil
}

func (f Float64) Value() (driver.Value, error) {
	if !f.Valid {
		return nil, nil
	}
	return f.Float64, nil
}

func (f Float64) MarshalJSON() ([]byte, error) {
	if !f.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(f.Float64)
}

func (f *Float64) UnmarshalJSON(b []byte) error {
	if str := string(b); str == `null` {
		return nil
	}
	err := json.Unmarshal(b, &f.Float64)
	f.Valid = (err == nil)
	return err
}

type Int16 sql.NullInt16

func (i *Int16) Scan(value interface{}) error {
	var ni sql.NullInt16
	if err := ni.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*i = Int16{ni.Int16, false}
	} else {
		*i = Int16{ni.Int16, true}
	}
	return nil
}

func (i Int16) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return i.Int16, nil
}

func (i Int16) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i.Int16)
}

func (i *Int16) UnmarshalJSON(b []byte) error {
	if str := string(b); str == `null` {
		return nil
	}
	err := json.Unmarshal(b, &i.Int16)
	i.Valid = (err == nil)
	return err
}

type Int32 sql.NullInt32

func (i *Int32) Scan(value interface{}) error {
	var ni sql.NullInt32
	if err := ni.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*i = Int32{ni.Int32, false}
	} else {
		*i = Int32{ni.Int32, true}
	}
	return nil
}

func (i Int32) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return i.Int32, nil
}

func (i Int32) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i.Int32)
}

func (i *Int32) UnmarshalJSON(b []byte) error {
	if str := string(b); str == `null` {
		return nil
	}
	err := json.Unmarshal(b, &i.Int32)
	i.Valid = (err == nil)
	return err
}

type Int64 sql.NullInt64

func (i *Int64) Scan(value interface{}) error {
	var ni sql.NullInt64
	if err := ni.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*i = Int64{ni.Int64, false}
	} else {
		*i = Int64{ni.Int64, true}
	}
	return nil
}

func (i Int64) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return i.Int64, nil
}

func (i Int64) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i.Int64)
}

func (i *Int64) UnmarshalJSON(b []byte) error {
	if str := string(b); str == `null` {
		return nil
	}
	err := json.Unmarshal(b, &i.Int64)
	i.Valid = (err == nil)
	return nil
}

type String sql.NullString

func (s *String) Scan(value interface{}) error {
	var ns sql.NullString
	if err := ns.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*s = String{ns.String, false}
	} else {
		*s = String{ns.String, true}
	}
	return nil
}

func (s String) Value() (driver.Value, error) {
	if !s.Valid {
		return nil, nil
	}
	return s.String, nil
}

func (s String) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *String) UnmarshalJSON(b []byte) error {
	if str := string(b); str == `null` {
		return nil
	}
	err := json.Unmarshal(b, &s.String)
	s.Valid = (err == nil)
	return err
}

type Time sql.NullTime

func (t *Time) Scan(value interface{}) error {
	var nt sql.NullTime
	if err := nt.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*t = Time{nt.Time, false}
	} else {
		*t = Time{nt.Time, true}
	}
	return nil
}

func (t Time) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(t.Time)
}

func (t *Time) UnmarshalJSON(b []byte) error {
	if str := string(b); str == `null` {
		return nil
	}
	err := json.Unmarshal(b, &t.Time)
	t.Valid = (err == nil)
	return err
}
