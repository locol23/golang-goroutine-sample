package data

import (
	"time"
)

func GetDataA() []string {
	time.Sleep(time.Millisecond * 300)
	return []string{
		"test01",
		"test02",
		"test03",
	}
}

func GetDataB() []string {
	time.Sleep(time.Millisecond * 300)
	return []string{
		"test02",
		"test03",
		"test04",
	}
}

func GetDataC() []string {
	time.Sleep(time.Millisecond * 300)
	return []string{
		"test03",
		"test05",
		"test06",
	}
}
