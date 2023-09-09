package window

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/viant/gmetric/counter"
)

//OnCounterDone represents on on counter callback
type OnCounterDone func(end time.Time, flag bool, err error) int64

//Operation represents operation metrics
type Operation struct {
	*counter.Operation
	Recent             []*counter.Operation
	index              int32
	UnitDuration       time.Duration `json:"-"`
	Unit               string
	RecentUnitDuration time.Duration `json:"-"`
	RecentUnit         string
	Provider           counter.Provider `json:"-"`

	recentCoverageDuration time.Duration // should be < 0
}

//Begin start metrics
func (c *Operation) Begin(started time.Time) counter.OnDone {
	totalOnDone := c.Operation.Begin(started)

	index := c.Index(started)
	currentIndex := atomic.LoadInt32(&c.index)
	if currentIndex != index {
		if atomic.CompareAndSwapInt32(&c.index, currentIndex, index) {
			c.Recent[index] = c.newOperation()
		}
	} else if c.Recent[index].Before(time.Now().Add(c.recentCoverageDuration)) {
		// if Recent[index] was created for a period before the oldest
		// possible Recent, clear it out.
		c.Recent[index] = c.newOperation()
	}

	recentAtTime := c.Recent[index]
	recentOnDone := recentAtTime.Begin(started)
	return func(end time.Time, values ...interface{}) int64 {
		values = counter.NormalizeValue(values)
		count := totalOnDone(end, values...)

		if recentAtTime.Before(end.Add(c.recentCoverageDuration)) {
			_ = recentOnDone(end, values...)
		}

		return count
	}
}

func (c *Operation) GetRecent(when time.Time) *counter.Operation {
	index := c.Index(when)
	if c.Recent[index].Before(time.Now().Add(c.recentCoverageDuration)) {
		// if Recent[index] was created for a period before the oldest
		// possible Recent, clear it out.
		c.Recent[index] = c.newOperation()
	}

	return c.Recent[index]
}

// Clean will replace each outdated Recent *counter.Operation with a new instance.
// This is mildly expensive, so should only be used when reading all Recent values.
func (c *Operation) Clean() {
	oldest := time.Now().Add(c.recentCoverageDuration)
	for i, r := range c.Recent {
		if r.Before(oldest) {
			c.Recent[i] = c.newOperation()
		}
	}
}

func (c *Operation) newOperation() *counter.Operation {
	return counter.NewOperation(c.UnitDuration, c.Provider)
}

//Index returns recent bucket index for supplied time
func (c *Operation) Index(atTime time.Time) int32 {
	recentIndex := int(time.Duration(atTime.UnixNano()) / c.RecentUnitDuration)
	index := int32(recentIndex % len(c.Recent))
	return index
}

//NewOperation creates basic rolling window counter
func NewOperation(recentBuckets int, recentUnit time.Duration, unit time.Duration, provider counter.Provider) Operation {
	if recentBuckets < 1 {
		recentBuckets = 2
	}
	if recentUnit == 0 {
		recentUnit = time.Minute
	}
	result := Operation{
		Operation:          counter.NewOperation(unit, provider),
		Recent:             make([]*counter.Operation, recentBuckets),
		index:              0,
		UnitDuration:       unit,
		Unit:               fmt.Sprintf("%s", unit),
		RecentUnitDuration: recentUnit,
		RecentUnit:         fmt.Sprintf("%s", recentUnit),
		Provider:           provider,

		recentCoverageDuration: -recentUnit * time.Duration(recentBuckets),
	}
	for i := range result.Recent {
		result.Recent[i] = counter.NewOperation(unit, provider)
	}
	return result
}
