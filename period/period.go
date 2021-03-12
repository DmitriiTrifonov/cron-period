package period

import (
	"cron-period/cron"
	"time"
)

const (
	day = time.Hour * 24
	week = day * 7
)

func ParsePeriod(c string) (period time.Duration, err error) {
	period = 0
	cronJob, err := cron.Parse(c)
	if err != nil {
		return 0, err
	}

	// Get a max period for minutes
	maxMinuteCycle := getMaxDuration(cronJob.Minutes, time.Minute, time.Hour)
	maxHourCycle := getMaxDuration(cronJob.Hours, time.Hour, day)

	period = maxMinuteCycle + (maxHourCycle - time.Hour)
	return period, err
}


func getMaxDuration(array []int, min, max time.Duration) time.Duration {
	arrayLen := len(array)

	last := time.Duration(array[arrayLen - 1]) * min
	first := time.Duration(array[0]) * min
	maxBetweenFirstAndLastPeriod := max - last + first

	maxCycle := time.Duration(0)
	for i := arrayLen - 1; i > 0; i-- {
		durationLastMinute := time.Duration(array[i]) * min
		durationFirstMinute := time.Duration(array[i - 1]) * min
		duration := durationLastMinute - durationFirstMinute
		if duration > maxCycle {
			maxCycle = duration
		}
	}

	if maxBetweenFirstAndLastPeriod > maxCycle {
		maxCycle = maxBetweenFirstAndLastPeriod
	}

	return maxCycle
}
