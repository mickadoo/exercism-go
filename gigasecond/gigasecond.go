// Package gigasecond calculates the moment when someone has lived for 10^9 seconds.
package gigasecond

// import path for the time package from the standard library
import (
	"math"
	"time"
)

// AddGigasecond takes a birth date and adds 10^9 seconds to it
func AddGigasecond(birth time.Time) time.Time {
	return birth.Add(time.Second * time.Duration(math.Pow(10, 9)))
}
