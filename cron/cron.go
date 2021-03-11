package cron

import (
	"fmt"
	"strings"
)

const spaceByte = 32

type Cron struct {
	Minutes []int
	Hours []int
	Days []int
	Months []int
	Weekdays []int
}

// Parse is used to parse a cron string
func Parse(c string) (*Cron, error) {
	cron := &Cron{}
	if c[len(c) - 1] == spaceByte {
		c = c[:len(c) - 1]
	}
	cp := strings.Split(c, " ")
	if len(cp) > 5 {
		return  nil, fmt.Errorf("unknown cron format: %s", c)
	}

	err := cron.parseMinutes(cp[0])
	if err != nil {
		return  nil, err
	}

	err = cron.parseHours(cp[1])
	if err != nil {
		return  nil, err
	}

	err = cron.parseDays(cp[2])
	if err != nil {
		return  nil, err
	}

	err = cron.parseMonths(cp[3])
	if err != nil {
		return  nil, err
	}

	err = cron.parseWeekdays(cp[4])

	return cron, nil
}
