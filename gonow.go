package gonow

import (
	"fmt"
	"time"
)

func (now *GoNow) StartToday() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, now.Time.Location())
}

func (now *GoNow) BringInWeekly(count int) (mdl []BringInWeekly) {
	t1 := now.StartToday()
	t2 := now.StartToday()
	d := ""
	for d == "" {
		if int(t2.Weekday()) == 0 {
			d = "Sunday"
			break
		}
		t2 = t2.Add(-1 * time.Hour * 24)
	}
	for i := 0; i < count; i++ {
		mdl = append(mdl, BringInWeekly{Range: fmt.Sprintf("%s to %s", t1.Format("2006-01-02"), t2.Format("2006-01-02")), T1: t1, T2: t2})
		if i == 0 {
			t1 = t2
			t2 = t2.Add(-7 * time.Hour * 24)
			continue
		}
		t1 = t1.Add(-7 * time.Hour * 24)
		t2 = t2.Add(-7 * time.Hour * 24)
	}
	return mdl
}

// find the first day of the month
func (now *GoNow) FirstDayOfTheMonth() time.Time {
	y, m, _ := now.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, now.Location())
}

// find the first day of the week
func (now *GoNow) FirstDayOfTheWeek() time.Time {
	t := now.StartToday()
	weekday := int(t.Weekday())
	if now.WeekStartDay != time.Sunday {
		weekStartDayInt := int(now.WeekStartDay)
		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}
	return t.AddDate(0, 0, -weekday)
}

func (now *GoNow) EndOfDay() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, 23, 59, 59, 0, now.Location())
}
