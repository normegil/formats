package time_test

import (
	"time"
	"testing"
	timeFormat "github.com/normegil/formats/time"
)

func TestTimeToString(t *testing.T) {
	cases := []time.Time{
		time.Now(),
		time.Now().Add(48 * time.Hour),
	}
	for _, testTime := range cases {
		value := timeFormat.MarshallableTime{Time: &testTime}.String()
		expected := testTime.Format(time.RFC3339)
		if expected != value {
			t.Errorf("Returned message (%s) don't correspond to expected message (%s)", expected, value)
		}
	}
}

func TestTimeMarshallJSON(t *testing.T) {
	cases := []struct {
		testName string
		input    time.Time
	}{
		{"JSON - Classic case", time.Now()},
	}
	for _, params := range cases {
		value, err := timeFormat.MarshallableTime{Time: &params.input}.MarshalJSON()
		if nil != err {
			t.Fatal(err.Error())
		}
		output := "\"" + params.input.Format(time.RFC3339) + "\""
		if output != string(value) {
			t.Errorf("Malformed JSON {Expected:%s;Got:%s}", output, string(value))
		}
	}
}

func TestTimeUnmarshallJSON(t *testing.T) {
	cases := []struct {
		testName string
		input    string
	}{
		{"JSON - Classic case", "2018-09-01T01:02:03Z"},
	}
	for _, params := range cases {
		var value timeFormat.MarshallableTime
		err := value.UnmarshalJSON([]byte(params.input))
		if err != nil {
			t.Fatal(err.Error())
		}
		expected, err := time.Parse(time.RFC3339, params.input)
		if expected != *value.Time {
			t.Errorf("Wrong date {Expected:%s;Got:%s}", expected, *value.Time)
		}
	}
}