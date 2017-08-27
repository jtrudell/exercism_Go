package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	hours, min := convertToMinutes(hour, minute)
	return Clock{hours, min}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutes int) Clock {
	hours, min := convertToMinutes(clock.hour, clock.minute+minutes)

	clock.hour = hours
	clock.minute = min
	return clock
}

func convertToMinutes(hour, minute int) (hours, minutes int) {
	totalMinutes := (hour * 60) + minute
	hours = (totalMinutes / 60) % 24
	minutes = totalMinutes % 60

	for hours < 0 {
		hours += 24
	}

	for minutes < 0 {
		minutes += 60
		if hours == 0 {
			hours = 23
		} else {
			hours -= 1
		}
	}
	return hours, minutes
}
