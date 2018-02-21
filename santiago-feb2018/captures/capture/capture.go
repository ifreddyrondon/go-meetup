package capture

import (
	"github.com/ifreddyrondon/go-talks/santiago-feb2018/structs/point"
	"github.com/ifreddyrondon/go-talks/santiago-feb2018/timestamp/timeutils"
)

type Capture struct {
	Point     point.Point
	Timestamp timeutils.Timestamp
}

type Captures []Capture

// START OMIT
func (c *Capture) UnmarshalJSON(data []byte) error {
	var p point.Point // HL
	if err := p.UnmarshalJSON(data); err != nil {
		panic(err)
	}
	var t timeutils.Timestamp // HL
	if err := t.UnmarshalJSON(data); err != nil {
		return err
	}
	*c = Capture{Point: p, Timestamp: t}
	return nil
}

// END OMIT
