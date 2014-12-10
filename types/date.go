package types

import (
	"encoding/json"
	"fmt"
	"time"
)

// Mon Jan 2 15:04:05 -0700 MST 2006
const dateLayout = "2006-01-02 15:04:05 -0700"

type Date time.Time

func (o *Date) UnmarshalJSON(d []byte) (err error) {
	var s string
	err = json.Unmarshal(d, &s)
	if err != nil {
		return
	}

	t, err := time.Parse(dateLayout, fmt.Sprintf("%s -0500", s))
	*o = Date(t.UTC())

	return
}

func (o Date) String() string {
	return time.Time(o).Format(dateLayout)
}
