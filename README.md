# Go Now


<div >
  <div style="display: flex; align-items: flex-start;">
	  GoNow is a fast and simple time tool kit 😃
  <img align="right" src="https://i.ibb.co/QYLZDj1/1200px-Go-gopher-favicon-svg.png" alt="drawing" style="width:200px;"/>
  </div>
</div>

## Install

```
go get -u github.com/yigitaltunay/gonow
```

## Features

- First Day Of The Month
- First Day Of The Week
- Bring In Weekly
- EndOfDay


## How to use ?

Calculating time based on current time


```go
package main

import (
	"fmt"
	"time"

	"github.com/yigitaltunay/gonow"
)

func main() {

    // Initialize GoNow with current time
	goNow := gonow.GoNowInit(time.Now()) 


	fmt.Println(goNow.Time) 
    // 2022-04-16 11:35:38.0766812 

	fmt.Println(goNow.FirstDayOfTheMonth()) 
    // 2022-04-01 00:00:00 

	fmt.Println(goNow.FirstDayOfTheWeek()) 
     // 2022-04-10 00:00:00 

	fmt.Println(goNow.EndOfDay())          
     // 2022-04-16 23:59:59

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

```

## License

MIT

