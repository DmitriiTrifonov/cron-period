package cron

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	jan = iota + 1
	feb
	mar
	apt
	may
	jun
	jul
	aug
	sep
	oct
	nov
	dec
)

var monthsMap = map[string]int{
	"JAN": jan,
	"FEB": feb,
	"MAR": mar,
	"APR": apt,
	"MAY": may,
	"JUN": jun,
	"JUL": jul,
	"AUG": aug,
	"SEP": sep,
	"OCT": oct,
	"NOV": nov,
	"DEC": dec,
}

var monthsFullYear = []int{
	jan,
	feb,
	mar,
	apt,
	may,
	jun,
	jul,
	aug,
	sep,
	oct,
	nov,
	dec,
}

func (c *Cron) parseMonths(s string) (err error) {
	m := make([]int, 0)
	divider := 1
	if s == "*" {
		m = monthsFullYear
		c.Months = m
		return
	}

	if strings.Contains(s, "/") {
		divider, err = strconv.Atoi(strings.Split(s,"/")[1])
		if err != nil{
			return
		}
	}
	num := 0
	for _, str := range strings.Split(strings.Split(s,"/")[0], ",") {
		if str == "*" {
			for i := 1; i <= (len(monthsMap) + 1) - divider; i += divider {
				m = append(m, i)
			}
			c.Months = m
			return
		}
		mon := monthsMap[str]
		if mon != 0 {
			m = append(m, mon)
			continue
		}
		num, err = strconv.Atoi(str)
		if err != nil {
			return fmt.Errorf("cannot parse month: %s", err)
		}

		if num < 1 || num > 12 {
			return fmt.Errorf("month is out of range: %d", num)
		}

		m = append(m, num)
	}
	c.Months = m
	return
}

