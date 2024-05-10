package json_additions

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDate_UnmarshalJSON(t *testing.T) {
	in := `"2022-12-31"`
	want := time.Date(2022, time.December, 31, 0, 0, 0, 0, time.UTC)

	var got RFC3339DATE
	if err := got.UnmarshalJSON([]byte(in)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !(got.Year == want.Year() && got.Month == want.Month() && got.Day == want.Day()) {
		t.Errorf("got date = %s, want %s", got, want)
	}
}

func TestDate_UnmarshalJSON_badFormat(t *testing.T) {
	in := `"31 Dec 22"`

	var got RFC3339DATE
	err := got.UnmarshalJSON([]byte(in))

	if err, ok := err.(*time.ParseError); !ok {
		t.Errorf("expected a time parse error, got: %v", err)
	}
}

func TestDate_MarshalJSON(t *testing.T) {
	testcases := map[string]struct {
		in   RFC3339DATE
		want string
	}{
		"without zero padding": {
			in:   RFC3339DATE{2022, time.December, 31},
			want: `"2022-12-31"`,
		},
		"with zero padding": {
			in:   RFC3339DATE{2022, time.July, 1},
			want: `"2022-07-01"`,
		},
		"initial value": {
			in:   RFC3339DATE{},
			want: `"0000-00-00"`,
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			got, err := json.Marshal(tc.in)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if string(got) != tc.want {
				t.Errorf("got date = %s, want %s", got, tc.want)
			}
		})
	}
}

func TestDate_String(t *testing.T) {
	testcases := map[string]struct {
		in   RFC3339DATE
		want string
	}{
		"without zero padding": {
			in:   RFC3339DATE{2022, time.December, 31},
			want: "2022-12-31",
		},
		"with zero padding": {
			in:   RFC3339DATE{2022, time.July, 1},
			want: "2022-07-01",
		},
		"initial value": {
			in:   RFC3339DATE{},
			want: "0000-00-00",
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			if got := tc.in.String(); got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
