package date

import (
	"strings"
	"time"
)

type Date time.Time

func (date *Date) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value) //parse time
	if err != nil {
		return err
	}
	*date = Date(t) //set result using the pointer
	return nil
}

func (date Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(date).Format("2006-01-02") + `"`), nil
}
