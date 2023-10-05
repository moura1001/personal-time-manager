package data

import "time"

const (
	TIMING_TYPE_INCREASE = "increase"
	TIMING_TYPE_CONSUME  = "consume"
)

type timing struct {
	Id         int
	timingType string
	start      time.Time
	stop       time.Time
}

var timings []timing = make([]timing, 0)
var timingType = ""
var id = 0

func GetTimingType() string {
	return timingType
}

func setTimingType(newTimingType string) {
	timingType = newTimingType
}

func GetTimings() []timing {
	return timings
}

func AddTiming(timingType string) {
	setTimingType(timingType)
	timings = append(timings, timing{Id: id, timingType: timingType, start: time.Now()})
	id++
}

func SetLastTimingStop() {
	if len(timings) > 0 {
		timings[len(timings)-1].stop = time.Now()
	}
}

func ClearTimings() {
	setTimingType("")
	timings = make([]timing, 0)
}
