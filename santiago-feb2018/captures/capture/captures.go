package capture

import (
	"encoding/json"
	"sort"
	"sync"

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

type indexCapture struct {
	index int
	*Capture
}

type job struct {
	index int
	data  json.RawMessage
}

func worker(wg *sync.WaitGroup, jobs <-chan job, results chan<- indexCapture) {
	for job := range jobs {
		var c Capture
		if err := json.Unmarshal(job.data, &c); err == nil {
			results <- indexCapture{index: job.index, Capture: &c}
		}
		wg.Done()
	}
}

func (p *Captures) UnmarshalJSON(data []byte) error {
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(raw))
	jobs := make(chan job, len(raw))
	results := make(chan indexCapture, len(raw))

	for w := 0; w < WorkersNumber; w++ {
		go worker(&wg, jobs, results)
	}
	for i, v := range raw {
		jobs <- job{index: i, data: v}
	}
	close(jobs)
	wg.Wait()
	close(results)

	processed := make([]indexCapture, 0, len(results))
	for data := range results {
		processed = append(processed, data)
	}
	sort.Slice(processed, func(i, j int) bool {
		return processed[i].index < processed[j].index
	})
	for _, v := range processed {
		*p = append(*p, *v.Capture)
	}

	return nil
}
