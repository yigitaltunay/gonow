package gonow

import "time"

// default week start day Sunday
var WeekStartDay = time.Sunday

var DefaultConfig *Config

type Config struct {
	WeekStartDay time.Weekday
	TimeLocation *time.Location
}

// embed time.Time
type GoNow struct {
	time.Time
	*Config
}

func (config *Config) WithConfig(t time.Time) *GoNow {
	return &GoNow{Time: t, Config: config}
}

// with default initialize Now with time
func GoNowInit(t time.Time) *GoNow {
	config := DefaultConfig
	if config == nil {
		config = &Config{
			WeekStartDay: WeekStartDay,
		}
	}

	return &GoNow{Time: t, Config: config}
}

// New initialize Now with time
func NewGoNow(t time.Time) *GoNow {
	return GoNowInit(t)
}
