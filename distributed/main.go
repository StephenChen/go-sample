package main

import (
	"distributed/loadbalance"
	"distributed/lock"
	"distributed/snowflake"
)

func main() {
	snowflake.Snowflake()

	snowflake.Sonyflake()

	lock.InProcess()
	lock.TryLock()

	loadbalance.LoadBalance()
}
