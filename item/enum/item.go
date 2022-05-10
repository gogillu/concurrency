package enum

import "database/sql/driver"

//go:generate enumer -type=ItemType

type ItemType int

const (
	Raw ItemType = iota
	Manufactured
	Imported
)

func (i *ItemType) Scan(value interface{}) error {
	expr, _ := value.([]byte)
	var err error
	*i, err = ItemTypeString(string(expr))
	if err != nil {
		return err
	}

	return nil
}

func (i ItemType) Value() (driver.Value, error) {
	return i.String(), nil
}
