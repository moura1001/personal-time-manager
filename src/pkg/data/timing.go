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
	finish := t.stop.Format(timingFormatRange)
	return fmt.Sprintf("%s - %s", finish, start)
}

func (t timing) FormatDuration() string {
	prefix := ""
	if t.timingType == TIMING_TYPE_CONSUME {
		prefix = "-"
	}

	duration := t.stop.Sub(t.start).Round(time.Second)
	hour := int(duration.Seconds() / 3600)
	minute := int(duration.Seconds()/60) % 60
	second := int(duration.Seconds()) % 60

	if hour > 0 {
		if minute > 0 && second > 0 {
			return fmt.Sprintf("%s%dh %dm %ds", prefix, hour, minute, second)
		} else if minute > 0 {
			return fmt.Sprintf("%s%dh %dm", prefix, hour, minute)
		} else if second > 0 {
			return fmt.Sprintf("%s%dh %ds", prefix, hour, second)
		} else {
			return fmt.Sprintf("%s%dh", prefix, hour)
		}

	} else if minute > 0 {
		if second > 0 {
			return fmt.Sprintf("%s%dm %ds", prefix, minute, second)
		} else {
			return fmt.Sprintf("%s%dm", prefix, minute)
		}

	} else if second > 0 {
		return fmt.Sprintf("%s%ds", prefix, second)

	} else {
		return fmt.Sprintf("%s0s", prefix)
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
