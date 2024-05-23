package timing

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ttp/json_additions"
)

func TestIntervalInitStart(t *testing.T) {
	got := NewInterval("10h", "11h").Start
	want, _ := time.ParseDuration("10h")

	assert.Equal(t, want, got)
}

func TestIntervalInitEnd(t *testing.T) {
	got := NewInterval("10h", "12h").End
	want, _ := time.ParseDuration("12h")

	assert.Equal(t, want, got)
}

func TestIntervalInitSingleString(t *testing.T) {
	got := IntervalFromStr("10h-12h").End
	want, _ := time.ParseDuration("12h")

	assert.Equal(t, want, got)
}

func TestAvailableTimeInitGeneral(t *testing.T) {
	got := NewAt("10h-12h30m,13h-14h").Intervals[0]
	want := *IntervalFromStr("10h-12h30m")

	assert.Equal(t, want, got)
}

func TestAtInsertSameBeginning(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*IntervalFromStr("10h-11h"))

	got := testAt
	want := NewAt("11h-12h30m,13h-14h")

	assert.Equal(t, want, got)
}

func TestAtInsertEarlyBeginning(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*IntervalFromStr("9h-11h"))

	got := testAt
	want := NewAt("11h-12h30m,13h-14h")

	assert.Equal(t, want, got)
}

func TestAtInsertSameEnd(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*IntervalFromStr("11h-12h30m"))

	got := testAt
	want := NewAt("10h-11h,13h-14h")

	assert.Equal(t, want, got)
}

func TestAtInsertLaterEnd(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*IntervalFromStr("11h-13h30m"))

	got := testAt
	want := NewAt("10h-11h,13h30m-14h")

	assert.Equal(t, want, got)
}

func TestAtInsertMiddleFirstInterval(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*IntervalFromStr("11h-12h"))

	got := testAt
	want := NewAt("10h-11h,12h-12h30m,13h-14h")

	assert.Equal(t, want, got)
}

func TestInsertTwoNewIntervalsInTheMiddle(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")

	got := InsertTwoNewIntervalsInTheMiddle(testAt.Intervals, 0, *IntervalFromStr("10h-12h30m"), *IntervalFromStr("11h-12h"))
	want := NewAt("10h-11h,12h-12h30m,13h-14h").Intervals

	assert.Equal(t, want, got)
}

func TestAtInsertMiddleSecondInterval(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*IntervalFromStr("13h30m-13h50m"))

	got := testAt
	want := NewAt("10h-12h30m,13h-13h30m,13h50m-14h")

	assert.Equal(t, want, got)
}

func TestAtInsertSeveralRanges(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*IntervalFromStr("11h-13h30m"))

	got := testAt
	want := NewAt("10h-11h,13h30m-14h")

	assert.Equal(t, want, got)
}

func TestAtInsertIgnoring(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*IntervalFromStr("8h-10h"))

	got := testAt
	want := NewAt("10h-12h30m,13h-14h")

	assert.Equal(t, want, got)
}

func TestAtInsertExcluding(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*IntervalFromStr("10h-12h30m"))

	got := testAt
	want := NewAt("13h-14h")

	assert.Equal(t, want, got)
}

func TestClearNullIntervals(t *testing.T) {
	testAt := AvailableTime{Intervals: []Interval{{Start: 0, End: 0}, {Start: 46800000000000, End: 50400000000000}, {Start: 0, End: 0}}}
	testAt.ClearNullIntervals()

	got := testAt
	want := AvailableTime{Intervals: []Interval{{Start: 46800000000000, End: 50400000000000}}}

	assert.Equal(t, want, got)
}

func TestAtInsertMany(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.InsertMany([]Interval{*IntervalFromStr("13h30m-13h50m"), *IntervalFromStr("13h-13h30m")})

	got := testAt
	want := NewAt("10h-12h30m,13h50m-14h")

	assert.Equal(t, want, got)
}

func TestAtInsertManyBackwards(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.InsertMany([]Interval{*IntervalFromStr("13h-13h30m"), *IntervalFromStr("13h30m-13h50m")})

	got := testAt
	want := NewAt("10h-12h30m,13h50m-14h")

	assert.Equal(t, want, got)
}

func TestIsIntervalAvailableTrue(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	got := testAt.IsIntervalAvailable(*IntervalFromStr("10h-12h"))
	assert.True(t, got)
}

func TestIsIntervalAvailableFalse(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	got := testAt.IsIntervalAvailable(*IntervalFromStr("9h-10h"))
	assert.False(t, got)
}

func TestGetWeekday(t *testing.T) {
	testDay := Day{At: *new(AvailableTime), Date: json_additions.RFC3339DATE{Year: 2024, Month: time.May, Day: 22}}

	got := testDay.GetWeekday()
	want := time.Wednesday

	assert.Equal(t, want, got)
}

func TestTimetableInit(t *testing.T) {
	testDays := [7]Day{
		{At: *NewAt("10h-18h"), Date: json_additions.RFC3339DATE{Year: 2024, Month: time.May, Day: 24}},
		{At: *NewAt("10h-18h"), Date: json_additions.RFC3339DATE{Year: 2024, Month: time.May, Day: 27}},
		{At: *NewAt("10h-18h"), Date: json_additions.RFC3339DATE{Year: 2024, Month: time.May, Day: 26}},
		{At: *NewAt("10h-18h"), Date: json_additions.RFC3339DATE{Year: 2024, Month: time.May, Day: 23}},
		{At: *NewAt("10h-18h"), Date: json_additions.RFC3339DATE{Year: 2024, Month: time.May, Day: 25}},
		{At: *NewAt("10h-18h"), Date: json_additions.RFC3339DATE{Year: 2024, Month: time.May, Day: 22}},
		{At: *NewAt("10h-18h"), Date: json_additions.RFC3339DATE{Year: 2024, Month: time.May, Day: 28}},
	}

	got := *NewTimetable(testDays)
	want := Timetable{
		Monday:    Day{At: *NewAt("10h-18h"), Date: json_additions.RFC3339DATE{Year: 2024, Month: time.May, Day: 27}},
		Tuesday:   Day{At: *NewAt("10h-18h"), Date: json_additions.RFC3339DATE{Year: 2024, Month: time.May, Day: 28}},
		Wednesday: Day{At: *NewAt("10h-18h"), Date: json_additions.RFC3339DATE{Year: 2024, Month: time.May, Day: 22}},
		Thursday:  Day{At: *NewAt("10h-18h"), Date: json_additions.RFC3339DATE{Year: 2024, Month: time.May, Day: 23}},
		Friday:    Day{At: *NewAt("10h-18h"), Date: json_additions.RFC3339DATE{Year: 2024, Month: time.May, Day: 24}},
		Saturday:  Day{At: *NewAt("10h-18h"), Date: json_additions.RFC3339DATE{Year: 2024, Month: time.May, Day: 25}},
		Sunday:    Day{At: *NewAt("10h-18h"), Date: json_additions.RFC3339DATE{Year: 2024, Month: time.May, Day: 26}},
	}

	assert.Equal(t, want, got)
}

func TestGetNearDates(t *testing.T) {
	got := GetNextSevenDaysDates(time.Date(2024, time.May, 22, 0, 0, 0, 0, time.Local))
	want := [7]json_additions.RFC3339DATE{
		{Year: 2024, Month: time.May, Day: 23},
		{Year: 2024, Month: time.May, Day: 24},
		{Year: 2024, Month: time.May, Day: 25},
		{Year: 2024, Month: time.May, Day: 26},
		{Year: 2024, Month: time.May, Day: 27},
		{Year: 2024, Month: time.May, Day: 28},
		{Year: 2024, Month: time.May, Day: 29},
	}

	assert.Equal(t, want, got)
}

func TestGetNearDatesEndOfWeek(t *testing.T) {
	got := GetNextSevenDaysDates(time.Date(2024, time.May, 26, 0, 0, 0, 0, time.Local))
	want := [7]json_additions.RFC3339DATE{
		{Year: 2024, Month: time.May, Day: 27},
		{Year: 2024, Month: time.May, Day: 28},
		{Year: 2024, Month: time.May, Day: 29},
		{Year: 2024, Month: time.May, Day: 30},
		{Year: 2024, Month: time.May, Day: 31},
		{Year: 2024, Month: time.June, Day: 1},
		{Year: 2024, Month: time.June, Day: 2},
	}

	assert.Equal(t, want, got)
}
