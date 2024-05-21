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