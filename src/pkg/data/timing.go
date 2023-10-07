package data

import (
	"fmt"
	"time"
)

const (
	TIMING_TYPE_INCREASE = "increase"
	TIMING_TYPE_CONSUME  = "consume"
	timingFormatRange    = "15:04"
)

type timing struct {
	Id         int
	timingType string
	start      time.Time
	stop       time.Time
}

func (t timing) FormatRange() string {
	start := t.start.Format(timingFormatRange)

	if t.stop.IsZero() {
		return start
	}

	finish := t.stop.Format(timingFormatRange)
	return fmt.Sprintf("%s - %s", finish, start)
}

func (t timing) FormatDuration() string {
	prefix := ""
	if t.timingType == TIMING_TYPE_CONSUME {
		prefix = "-"
	}

	if t.stop.IsZero() || t.start.After(t.stop) {
		return fmt.Sprintf("%s0s", prefix)
	}

	duration := t.stop.Sub(t.start)

	return fmt.Sprintf("%s%s", prefix, getFormattedDuration(duration.Abs()))
}

func getFormattedDuration(duration time.Duration) string {
	d := duration.Round(time.Second)

	hour := int(d.Seconds() / 3600)
	minute := int(d.Seconds()/60) % 60
	second := int(d.Seconds()) % 60

	if hour > 0 {
		if minute > 0 && second > 0 {
			return fmt.Sprintf("%dh %dm %ds", hour, minute, second)
		} else if minute > 0 {
			return fmt.Sprintf("%dh %dm", hour, minute)
		} else if second > 0 {
			return fmt.Sprintf("%dh %ds", hour, second)
		} else {
			return fmt.Sprintf("%dh", hour)
		}

	} else if minute > 0 {
		if second > 0 {
			return fmt.Sprintf("%dm %ds", minute, second)
		} else {
			return fmt.Sprintf("%dm", minute)
		}

	} else if second > 0 {
		return fmt.Sprintf("%ds", second)

	} else {
		return "0s"
	}
}

var timings []timing = make([]timing, 0)
var timingType = ""
var id = 0

func getTimingType() string {
	return timingType
}

func setTimingType(newTimingType string) {
	timingType = newTimingType
}

func GetTimings() []timing {
	return timings
}

func AddTiming(timingType string) {
	if getTimingType() != "" {
		setLastTimingStop()
	}

	if getTimingType() == timingType {
		setTimingType("")
		return
	}

	setTimingType(timingType)
	timings = append(timings, timing{Id: id, timingType: timingType, start: time.Now()})
	id++
}

func setLastTimingStop() {
	if len(timings) > 0 {
		timings[len(timings)-1].stop = time.Now()
	}
}

func ClearTimings() {
	setTimingType("")
	timings = make([]timing, 0)
}

func GetFormattedTotalTiming() string {
	if len(timings) == 0 {
		return "0s"
	}

	var totalTiming time.Duration = 0

	for _, timing := range timings {
		if !timing.start.IsZero() && !timing.stop.IsZero() {
			duration := timing.stop.Sub(timing.start)

			if timing.timingType == TIMING_TYPE_INCREASE {
				totalTiming += duration
			} else {
				totalTiming -= duration
			}
		}
	}

	formattedDuration := getFormattedDuration(totalTiming.Abs())

	if totalTiming > 0 {
		return formattedDuration
	} else {
		return fmt.Sprintf("- %s", formattedDuration)
	}
}
