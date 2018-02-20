package capture

import (
	"github.com/ifreddyrondon/go-talks/santiago-feb2018/structs/point"
	"github.com/ifreddyrondon/go-talks/santiago-feb2018/timestamp/timeutils"
)

const WorkersNumber = 4

type Capture struct {
	Point     point.Point
	Timestamp timeutils.Timestamp
}

type Captures []Capture

func (c *Capture) UnmarshalJSON(data []byte) error {
	var p point.Point
	if err := p.UnmarshalJSON(data); err != nil {
		panic(err)
	}
	var t timeutils.Timestamp
	if err := t.UnmarshalJSON(data); err != nil {
		return err
	}
	*c = Capture{Point: p, Timestamp: t}
	return nil
}
