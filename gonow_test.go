package gonow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func JsonPrettyPrint(b []byte) []byte {
	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")
	return out.Bytes()
}

func init() {
	NewGoNow(time.Now()).FirstDayOfTheMonth()
	test1 := NewGoNow(time.Now()).BringInWeekly(4)
	t, _ := json.Marshal(test1)
	fmt.Printf("%s", JsonPrettyPrint(t))
}

func TestMain(t *testing.T) {

	goNow := GoNowInit(time.Now())

	fmt.Println(goNow.Time)

	fmt.Println(goNow.FirstDayOfTheMonth())
	fmt.Println(goNow.FirstDayOfTheWeek())
	fmt.Println(goNow.EndOfDay())

	goNow.BringInWeekly(4)
	/*
			[
		  {
		    "Range": "2022-04-16 to 2022-04-10",
		    "T1": "2022-04-16T00:00:00+03:00",
		    "T2": "2022-04-10T00:00:00+03:00"
		  },
		  {
		    "Range": "2022-04-10 to 2022-04-03",
		    "T1": "2022-04-10T00:00:00+03:00",
		    "T2": "2022-04-03T00:00:00+03:00"
		  },
		  {
		    "Range": "2022-04-03 to 2022-03-27",
		    "T1": "2022-04-03T00:00:00+03:00",
		    "T2": "2022-03-27T00:00:00+03:00"
		  },
		  {
		    "Range": "2022-03-27 to 2022-03-20",
		    "T1": "2022-03-27T00:00:00+03:00",
		    "T2": "2022-03-20T00:00:00+03:00"
		  }
		]
	*/

}
