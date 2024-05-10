package json_additions

import (
	"fmt"
	"strings"
	"time"
)

const rfc3339 string = "2006-01-02"

// RFC3339DATE represents a date without a time component, encoded as a string
// in the "YYYY-MM-DD" format.
type RFC3339DATE struct {
	Year  int
	Month time.Month
	Day   int
}

// UnmarshalJSON implements json.Unmarshaler inferface.
func (d *RFC3339DATE) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(rfc3339, strings.Trim(string(b), `"`))
	if err != nil {
		return err
	}
	d.Year, d.Month, d.Day = t.Date()
	return nil
}

// MarshalJSON implements json.Marshaler interface.
func (d RFC3339DATE) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"%04d-%02d-%02d"`, d.Year, d.Month, d.Day)
	return []byte(s), nil
}

// String defines a string representation.
// It will be called automatically when you try to convert struct instance
// to a string.
func (d RFC3339DATE) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}
