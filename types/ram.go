package types

import (
	"encoding/json"
	"fmt"
)

type RAM int

func (o *RAM) UnmarshalJSON(d []byte) (err error) {
	var s string
	err = json.Unmarshal(d, &s)
	if err != nil {
		return
	}

	_, err = fmt.Sscanf(s, "%d MB", o)
	if err != nil {
		return
	}

	return
}
