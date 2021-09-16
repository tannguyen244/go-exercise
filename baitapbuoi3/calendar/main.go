package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	firstday := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	lastday := firstday.AddDate(0, 1, 0).Add(time.Nanosecond * -1)
	firtdayWeek := (firstday.Weekday())
	_, _, d := (lastday.Date())
	fmt.Printf("    %2v %2d\n", t.Month(), t.Year())
	fmt.Println("Su Mo Tu We Th Fr Sa")
	switch {
	case firtdayWeek == 0:
		for i := 1; i < d+1; i++ {
			if i == 1 {
				fmt.Printf("%2d ", i)
			}
			if day := time.Date(t.Year(), t.Month(), i, 0, 0, 0, 0, time.Local); day.Weekday() == 6 {
				fmt.Printf("")
			}
		}
	case firtdayWeek == 1:
		for i := 1; i < d+1; i++ {
			if i == 1 {
				fmt.Printf("%5d ", i)
			} else {
				fmt.Printf("%2d ", i)
			}
			if day := time.Date(t.Year(), t.Month(), i, 0, 0, 0, 0, time.Local); day.Weekday() == 6 {
				fmt.Println()
			}
		}
	case firtdayWeek == 2:
		for i := 1; i < d+1; i++ {
			if i == 1 {
				fmt.Printf("%8d ", i)
			} else {
				fmt.Printf("%2d ", i)
			}
			if day := time.Date(t.Year(), t.Month(), i, 0, 0, 0, 0, time.Local); day.Weekday() == 6 {
				fmt.Println()
			}
		}
	case firtdayWeek == 3:
		for i := 1; i < d+1; i++ {
			if i == 1 {
				fmt.Printf("%11d ", i)
			} else {
				fmt.Printf("%2d ", i)
			}
			if day := time.Date(t.Year(), t.Month(), i, 0, 0, 0, 0, time.Local); day.Weekday() == 6 {
				fmt.Println()
			}
		}
	case firtdayWeek == 4:
		for i := 1; i < d+1; i++ {
			if i == 1 {
				fmt.Printf("%14d ", i)
			} else {
				fmt.Printf("%2d ", i)
			}
			if day := time.Date(t.Year(), t.Month(), i, 0, 0, 0, 0, time.Local); day.Weekday() == 6 {
				fmt.Println()
			}
		}
	case firtdayWeek == 5:
		for i := 1; i < d+1; i++ {
			if i == 1 {
				fmt.Printf("%17d ", i)
			} else {
				fmt.Printf("%2d ", i)
			}
			if day := time.Date(t.Year(), t.Month(), i, 0, 0, 0, 0, time.Local); day.Weekday() == 6 {
				fmt.Println()
			}
		}
	default:
		for i := 1; i < d+1; i++ {
			if i == 1 {
				fmt.Printf("%20d ", i)
			} else {
				fmt.Printf("%2d ", i)
			}
			if day := time.Date(t.Year(), t.Month(), i, 0, 0, 0, 0, time.Local); day.Weekday() == 6 {
				fmt.Println()
			}
		}
	}
}
