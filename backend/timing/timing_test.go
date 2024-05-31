package timing

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func TestEqualIntervalSlicesTrue(t *testing.T) {
	testSlice1 := []Interval{*IntervalFromStr("12h-14h"), *IntervalFromStr("15h-18h")}
	testSlice2 := []Interval{*IntervalFromStr("12h-14h"), *IntervalFromStr("15h-18h")}

	got := EqualIntervalSlices(testSlice1, testSlice2)

	assert.True(t, got)
}

func TestEqualIntervalSlicesFalse(t *testing.T) {
	testSlice1 := []Interval{*IntervalFromStr("12h-14h"), *IntervalFromStr("15h-18h")}
	testSlice2 := []Interval{*IntervalFromStr("9h-14h"), *IntervalFromStr("15h-20h")}

	got := EqualIntervalSlices(testSlice1, testSlice2)

	assert.False(t, got)
}

func TestEqualIntervalSlicesDifferentLength(t *testing.T) {
	testSlice1 := []Interval{*IntervalFromStr("12h-14h"), *IntervalFromStr("15h-18h")}
	testSlice2 := []Interval{*IntervalFromStr("9h-14h"), *IntervalFromStr("15h-20h"), *IntervalFromStr("19h-22h")}

	got := EqualIntervalSlices(testSlice1, testSlice2)

	assert.False(t, got)
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

func TestIsAtEqualTrue(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	got := testAt.IsEqual(*NewAt("10h-12h30m,13h-14h"))
	assert.True(t, got)
}

func TestIsAtEqualFalse(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	got := testAt.IsEqual(*NewAt("9h-10h,13h-14h"))
	assert.False(t, got)
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

func TestGetDate(t *testing.T) {
	got := GetDate(2024, time.May, 26)
	want := time.Date(2024, time.May, 26, 0, 0, 0, 0, time.Local)

	assert.Equal(t, want, got)
}

func TestGetTime(t *testing.T) {
	got := GetTime(12, 30, 0)
	want := time.Date(0, time.January, 0, 12, 30, 0, 0, time.Local)

	assert.Equal(t, want, got)
}

func TestGetWeekday(t *testing.T) {
	testDay := Day{At: *new(AvailableTime), Date: GetDate(2024, time.May, 22)}

	got := testDay.GetWeekday()
	want := time.Wednesday

	assert.Equal(t, want, got)
}

func TestTimetableInit(t *testing.T) {
	testDays := [7]Day{
		{At: *NewAt("10h-18h"), Date: GetDate(2024, time.May, 24)},
		{At: *NewAt("10h-18h"), Date: GetDate(2024, time.May, 27)},
		{At: *NewAt("10h-18h"), Date: GetDate(2024, time.May, 26)},
		{At: *NewAt("10h-18h"), Date: GetDate(2024, time.May, 23)},
		{At: *NewAt("10h-18h"), Date: GetDate(2024, time.May, 25)},
		{At: *NewAt("10h-18h"), Date: GetDate(2024, time.May, 22)},
		{At: *NewAt("10h-18h"), Date: GetDate(2024, time.May, 28)},
	}

	got := *NewTimetable(testDays)
	want := Timetable{
		Monday:    Day{At: *NewAt("10h-18h"), Date: time.Date(2024, time.May, 27, 0, 0, 0, 0, time.Local)},
		Tuesday:   Day{At: *NewAt("10h-18h"), Date: time.Date(2024, time.May, 28, 0, 0, 0, 0, time.Local)},
		Wednesday: Day{At: *NewAt("10h-18h"), Date: time.Date(2024, time.May, 22, 0, 0, 0, 0, time.Local)},
		Thursday:  Day{At: *NewAt("10h-18h"), Date: time.Date(2024, time.May, 23, 0, 0, 0, 0, time.Local)},
		Friday:    Day{At: *NewAt("10h-18h"), Date: time.Date(2024, time.May, 24, 0, 0, 0, 0, time.Local)},
		Saturday:  Day{At: *NewAt("10h-18h"), Date: time.Date(2024, time.May, 25, 0, 0, 0, 0, time.Local)},
		Sunday:    Day{At: *NewAt("10h-18h"), Date: time.Date(2024, time.May, 26, 0, 0, 0, 0, time.Local)},
	}

	assert.Equal(t, want, got)
}

func TestGetNextSevenDaysDates(t *testing.T) {
	got := GetNextSevenDaysDates(time.Date(2024, time.May, 22, 0, 0, 0, 0, time.Local))
	want := [7]time.Time{
		time.Date(2024, time.May, 23, 0, 0, 0, 0, time.Local),
		time.Date(2024, time.May, 24, 0, 0, 0, 0, time.Local),
		time.Date(2024, time.May, 25, 0, 0, 0, 0, time.Local),
		time.Date(2024, time.May, 26, 0, 0, 0, 0, time.Local),
		time.Date(2024, time.May, 27, 0, 0, 0, 0, time.Local),
		time.Date(2024, time.May, 28, 0, 0, 0, 0, time.Local),
		time.Date(2024, time.May, 29, 0, 0, 0, 0, time.Local),
	}

	assert.Equal(t, want, got)
}

func TestGetNextSevenDaysDatesEndOfWeek(t *testing.T) {
	got := GetNextSevenDaysDates(time.Date(2024, time.May, 26, 0, 0, 0, 0, time.Local))
	want := [7]time.Time{
		time.Date(2024, time.May, 27, 0, 0, 0, 0, time.Local),
		time.Date(2024, time.May, 28, 0, 0, 0, 0, time.Local),
		time.Date(2024, time.May, 29, 0, 0, 0, 0, time.Local),
		time.Date(2024, time.May, 30, 0, 0, 0, 0, time.Local),
		time.Date(2024, time.May, 31, 0, 0, 0, 0, time.Local),
		time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
		time.Date(2024, time.June, 2, 0, 0, 0, 0, time.Local),
	}

	assert.Equal(t, want, got)
}
