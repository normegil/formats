package time

import (
	"time"

	"github.com/pkg/errors"
)

// Time is just an alias for time.Time, which will be marshalled in a standard time representation (RFC3339)
type MarshallableTime struct {
	*time.Time
}

func (j MarshallableTime) MarshalJSON() ([]byte, error) {
	json := "\"" + j.String() + "\""
	return []byte(json), nil
}

func (j *MarshallableTime) UnmarshalJSON(b []byte) error {
	toUnmarshal := j.clean(string(b))
	t, err := time.Parse(time.RFC3339, toUnmarshal)
	if nil != err {
		return errors.Wrapf(err, "Could not Unmarshall %s into Time", toUnmarshal)
	}
	j.Time = &t
	return nil
}

func (j *MarshallableTime) clean(toClean string) string {
	toReturn := toClean
	if '"' == toReturn[0] {
		toReturn = toReturn[1:]
	}
	if '"' == toReturn[len(toReturn)-1] {
		toReturn = toReturn[:len(toReturn)-1]
	}
	return toReturn
}

// String return the RFC3339 string representation of the time
func (j MarshallableTime) String() string {
	return j.Time.Format(time.RFC3339)
}
