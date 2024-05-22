package timing

import (
	"log"
	"strings"
	"time"

	"github.com/ttp/json_additions"
)

type Interval struct {
	Start time.Duration
	End   time.Duration
}

func NewInterval(start, end string) *Interval {
	interval := new(Interval)
	var err error
	interval.Start, err = time.ParseDuration(start)
	if err != nil {
		log.Fatal(err)
	}
	interval.End, err = time.ParseDuration(end)
	if err != nil {
		log.Fatal(err)
	}
	return interval
}

func NewIntervalFromString(intervalString string) *Interval {
	splitted := strings.Split(intervalString, "-")
	return NewInterval(splitted[0], splitted[1])
}

type AvailableTime struct {
	Intervals []Interval
}

func NewAt(commaSeparatedIntervals string) *AvailableTime {
	at := new(AvailableTime)
	intervals := strings.Split(commaSeparatedIntervals, ",")
	for i := 0; i < len(intervals); i++ {
		interval := NewIntervalFromString(intervals[i])
		at.Intervals = append(at.Intervals, *interval)
	}
	return at
}

func (at *AvailableTime) Insert(interval Interval) {
	for i := 0; i < len(at.Intervals); i++ {
		atInterval := at.Intervals[i]
		if HaveSameBeginnning(atInterval, interval) {
			atInterval.Start = interval.End
		} else if HaveSameEnding(atInterval, interval) {
			atInterval.End = interval.Start
		} else if SecondIntervalIsInTheMiddle(atInterval, interval) {
			at.Intervals = InsertTwoNewIntervalsInTheMiddle(at.Intervals, i, atInterval, interval)
			return
		} else if SecondIntervalIsBigger(atInterval, interval) {
			atInterval = *new(Interval)
		}
		at.Intervals[i] = atInterval
	}
	at.ClearNullIntervals()
}

func (at *AvailableTime) ClearNullIntervals() {
	var newIntervals []Interval
	for i := 0; i < len(at.Intervals); i++ {
		atInterval := at.Intervals[i]
		if atInterval == *new(Interval) {
			continue
		}
		newIntervals = append(newIntervals, atInterval)
	}
	at.Intervals = newIntervals
}

func HaveSameBeginnning(first, second Interval) bool {
	return first.Start >= second.Start && first.End > second.End && first.Start < second.End
}

func HaveSameEnding(first, second Interval) bool {
	return first.Start < second.Start && first.End <= second.End && first.End > second.Start
}

func SecondIntervalIsInTheMiddle(first, second Interval) bool {
	return first.Start < second.Start && first.End > second.End
}

func InsertTwoNewIntervalsInTheMiddle(where []Interval, place int, first, second Interval) []Interval {
	before := where[:place]
	var after []Interval
	if len(where[place+1:]) > 0 {
		after = make([]Interval, len(where[place+1:]))
		copy(after, where[place+1:])
	}

	newFirstInterval := NewInterval(first.Start.String(), second.Start.String())
	newSecondInterval := NewInterval(second.End.String(), first.End.String())
	newIntervals := append(before, *newFirstInterval)
	newIntervals = append(newIntervals, *newSecondInterval)
	newIntervals = append(newIntervals, after...)
	return newIntervals
}

func SecondIntervalIsBigger(first, second Interval) bool {
	return first.Start >= second.Start && first.End <= second.End
}

type Day struct {
	At   AvailableTime
	Date json_additions.RFC3339DATE
}

type Timetable struct {
	Mon Day
	Tue Day
	Wed Day
	Thu Day
	Fri Day
	Sat Day
	Sun Day
}
