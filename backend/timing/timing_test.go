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
	got := NewIntervalFromString("10h-12h").End
	want, _ := time.ParseDuration("12h")

	assert.Equal(t, want, got)
}

func TestAvailableTimeInitGeneral(t *testing.T) {
	got := NewAt("10h-12h30m,13h-14h").Intervals[0]
	want := *NewIntervalFromString("10h-12h30m")

	assert.Equal(t, want, got)
}

func TestAtInsertSameBeginning(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*NewIntervalFromString("10h-11h"))

	got := testAt
	want := NewAt("11h-12h30m,13h-14h")

	assert.Equal(t, want, got)
}

func TestAtInsertEarlyBeginning(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*NewIntervalFromString("9h-11h"))

	got := testAt
	want := NewAt("11h-12h30m,13h-14h")

	assert.Equal(t, want, got)
}

func TestAtInsertSameEnd(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*NewIntervalFromString("11h-12h30m"))

	got := testAt
	want := NewAt("10h-11h,13h-14h")

	assert.Equal(t, want, got)
}

func TestAtInsertLaterEnd(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*NewIntervalFromString("11h-13h30m"))

	got := testAt
	want := NewAt("10h-11h,13h30m-14h")

	assert.Equal(t, want, got)
}

func TestAtInsertMiddleFirstInterval(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*NewIntervalFromString("11h-12h"))

	got := testAt
	want := NewAt("10h-11h,12h-12h30m,13h-14h")

	assert.Equal(t, want, got)
	//интервал из нуля!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
}

func TestInsertTwoNewIntervalsInTheMiddle(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")

	got := InsertTwoNewIntervalsInTheMiddle(testAt.Intervals, 0, *NewIntervalFromString("10h-12h30m"), *NewIntervalFromString("11h-12h"))
	want := NewAt("10h-11h,12h-12h30m,13h-14h").Intervals

	assert.Equal(t, want, got)
}

func TestAtInsertMiddleSecondInterval(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*NewIntervalFromString("13h30m-13h50m"))

	got := testAt
	want := NewAt("10h-12h30m,13h-13h30m,13h50m-14h")

	assert.Equal(t, want, got)
}

func TestAtInsertSeveralRanges(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*NewIntervalFromString("11h-13h30m"))

	got := testAt
	want := NewAt("10h-11h,13h30m-14h")

	assert.Equal(t, want, got)
}

func TestAtInsertIgnoring(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*NewIntervalFromString("8h-10h"))

	got := testAt
	want := NewAt("10h-12h30m,13h-14h")

	assert.Equal(t, want, got)
}

func TestAtInsertExcluding(t *testing.T) {
	testAt := NewAt("10h-12h30m,13h-14h")
	testAt.Insert(*NewIntervalFromString("10h-12h30m"))

	got := testAt
	want := NewAt("13h-14h")

	assert.Equal(t, want, got)
}

func TestClearNullIntervals(t *testing.T) {
	testAt := AvailableTime{Intervals:[]Interval{{Start:0, End:0}, {Start:46800000000000, End:50400000000000},{Start:0, End:0}}}
	testAt.ClearNullIntervals()

	got := testAt
	want := AvailableTime{Intervals:[]Interval{{Start:46800000000000, End:50400000000000}}}


	assert.Equal(t, want, got)
}
