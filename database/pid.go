package database

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strconv"
)

var (
	ErrInvalidPID = errors.New("ID is not valid")
)

// PID Primary ID
type PID int64

// NilPID Null Primary ID
var NilPID = PID(0)

// NullPID can be used with the standard sql package to represent a
// UUID value that can be NULL in the database
type NullPID struct {
	PID   PID
	ValID bool
}

// Value - Implementation of valuer for database/sql
func (ID PID) Value() (driver.Value, error) {
	// value needs to be a base driver.Value type
	// such as string, bool and ...
	return int64(ID), nil
}

// Scan implements the sql.Scanner interface.
// A 16-byte slice is handled by UnmarshalBinary, while
// a longer byte slice or a string is handled by UnmarshalText.
func (ID *PID) Scan(src interface{}) error {
	if src == nil {
		*ID = PID(0)
		return nil
	}

	// ns := sql.NullInt64{}
	// if err := ns.Scan(src); err != nil {
	//     return err
	// }
	//
	// if !ns.ValID {
	//     return errors.New("scan not valid")
	// }
	//
	// nsv, _ := ns.Value()
	// *ID = PID(nsv.(int64))

	*ID = PID(src.(int64))

	return nil
}

func (ID PID) String() string {
	return strconv.Itoa(int(ID))
}

// CheckPID ...
func (ID PID) CheckPID() bool {
	return true
}

func (ID PID) IsValid() bool {
	return int64(ID) > 0
}

// ParsePID , parses a string ID to a PID one
func ParsePID(ID interface{}) (pID PID, err error) {
	switch ID.(type) {
	case string:
		var d int
		if d, err = strconv.Atoi(ID.(string)); err != nil {
			return 0, err
		}
		pID = PID(d)
	case int:
		pID = PID(ID.(int))
	case float64:
		pID = PID(ID.(float64))
	case PID:
		pID = ID.(PID)
	}

	if !pID.IsValid() {
		err = ErrInvalidPID
	}

	return pID, err
}

// Parse ...
func Parse(ID string) PID {
	pid, _ := ParsePID(ID)
	return pid
}

// Validate ...
func Validate(ID string) (PID, bool) {
	pid, err := ParsePID(ID)
	return pid, err == nil
}

// String ...
func String(ID PID) string {
	return ID.String()
}

// CheckPID ...
func CheckPID(ID PID) bool {
	return ID.CheckPID()
}

// Value implements the driver.Valuer interface.
func (u NullPID) Value() (driver.Value, error) {
	if !u.ValID {
		return nil, nil
	}
	// Delegate to int64 Value function
	return u.PID.Value()
}

// Scan implements the sql.Scanner interface.
func (u *NullPID) Scan(src interface{}) error {
	if src == nil {
		u.PID, u.ValID = NilPID, false
		return nil
	}

	// Delegate to int64 Scan function
	u.ValID = true
	return u.PID.Scan(src)
}

// MarshalJSON ...
func (u NullPID) MarshalJSON() ([]byte, error) {
	if u.ValID {
		return json.Marshal(u.PID)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON ...
func (u *NullPID) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *PID
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		u.ValID = true
		u.PID = *x
	} else {
		u.ValID = false
	}
	return nil
}
