package timeutils

import (
	"encoding/json"
	"time"
)

// START OMIT
type Timestamp struct{ time.Time } // HL

type timestampJSON struct {
	Date      json.Number `json:"date"`
	Timestamp json.Number `json:"timestamp"`
}

func (t *Timestamp) UnmarshalJSON(data []byte) error { // HL
	t.Time = time.Now()
	var model timestampJSON
	if err := json.Unmarshal(data, &model); err != nil {
		return nil
	}
	date := getDate(&model)
	parsedTime, err := ParseDateString(date) // HL
	if err != nil {
		return nil
	}
	t.Time = parsedTime.UTC() // HL
	return nil
}

// END OMIT

func getDate(model *timestampJSON) string {
	if model.Date == "" {
		return model.Timestamp.String()
	}
	return model.Date.String()
}
