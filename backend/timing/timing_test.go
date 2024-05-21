package timing

import (
	"testing"
	"time"
)

func TestIntervalInitStart(t *testing.T) {
	got := NewInterval("10h","11h").Start
	want, _ := time.ParseDuration("10h")

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestIntervalInitEnd(t *testing.T) {
	got := NewInterval("10h","12h").End
	want, _ := time.ParseDuration("12h")

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestIntervalInitSingleString(t *testing.T) {
	got := NewIntervalFromString("10h-12h").End
	want, _ := time.ParseDuration("12h")

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}


func TestAvailableTimeInitGeneral(t *testing.T) {
	got := NewAt("10h-12h30m,13h-14h").Intervals[0]
	want := Interval{Start: time.Duration(1000000000*60*60*10), End: time.Duration(1000000000*60*60*12 + 1000000000*60*30)}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}