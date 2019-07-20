package main

import (
	"fmt"
	"time"
)

/*func main()  {

	birthDate,err  := time.Parse(time.RFC3339,"1996-04-30T00:00:00+00:00")
	if err != nil {
		log.Println(err)
	}
	nowDate ,err := time.Parse(time.RFC3339,"2019-07-20T00:00:00+00:00")
	if err != nil {
		log.Println(err)
	}

nowDate.



	fmt.Println(birthDate)
	fmt.Println(nowDate)
}
*/
func Cal(birthDateY int, birthDateM int, birthDateD int, nowDateY int, nowDateM int, nowDateD int) (int, int, int, int, int) {
	// The leap year 2016 had 366 days.
	birthDate := Date(birthDateY, birthDateM, birthDateD)
	nowDate := Date(nowDateY, nowDateM, nowDateD)
	//days := nowDate.Sub(birthDate).Hours() / 24
	//weeks := nowDate.Sub(birthDate).Hours()/ (24*7)
	seconds := nowDate.Sub(birthDate).Seconds()
	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24
	weeks := days / 7

	//fmt.Println("days = " ,days) // 366
	//fmt.Println(math.Ceil(days /7))
	fmt.Println("seconds ", int(seconds))
	fmt.Println("minutes ", int(minutes))
	fmt.Println("hours ", int(hours))
	fmt.Println("day ", int(days))
	fmt.Println(int(weeks), " weeks ", int(days)-int(weeks)*7, " days ")
	return int(seconds), int(minutes), int(hours), int(days), int(days) - int(weeks)*7
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
