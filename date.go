package bank

import "time"

type Date string

type TodaysDateProvider func() Date

func TodaysSystemDateProvider() Date {
	return Date(time.Now().Format("2006-01-02"))
}
