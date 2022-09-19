package jtime

import "time"

// count time
// add by yjh 211125
type TimeCount struct {
	_startTime time.Time
	_stopTime  time.Time
}

func (t *TimeCount) _calc() time.Duration {
	return t._stopTime.Sub(t._startTime)
}
func (t *TimeCount) Start() {
	t._startTime = time.Now()
}

func (t *TimeCount) Minutes() float64 {
	t._stopTime = time.Now()

	return t._calc().Minutes()
}
func (t *TimeCount) Seconds() float64 {
	t._stopTime = time.Now()

	return t._calc().Seconds()
}
func (t *TimeCount) Hours() float64 {
	t._stopTime = time.Now()

	return t._calc().Hours()
}

func (t *TimeCount) Microseconds() int64 {
	t._stopTime = time.Now()

	return t._calc().Microseconds()
}
func (t *TimeCount) Milliseconds() int64 {
	t._stopTime = time.Now()

	return t._calc().Milliseconds()
}
