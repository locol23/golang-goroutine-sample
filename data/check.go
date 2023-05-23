package data

import (
	"fmt"
	"time"
)

func Check(actual []string, d time.Duration) error {
	err := checkTime(d)
	if err != nil {
		return err
	}

	err = checkData(actual)
	if err != nil {
		return err
	}

	return nil
}

func checkTime(actual time.Duration) error {
	const expect = time.Millisecond * 310

	if actual >= expect {
		return fmt.Errorf("TIME OVER!!!\nexpect: %v\nactual: %v", expect, actual)
	}

	return nil
}

func checkData(actual []string) error {
	expect := []string{
		"test01",
		"test02",
		"test03",
		"test04",
		"test05",
		"test06",
	}

	actualM := make(map[string]struct{})

	for _, item := range actual {
		actualM[item] = struct{}{}
	}

	if len(expect) != len(actual) {
		return fmt.Errorf("Could not match!!!\nexpect: %+v\nactual: %+v", expect, actual)
	}

	for _, item := range expect {
		if _, ok := actualM[item]; !ok {
			return fmt.Errorf("Could not match!!!\nexpect: %+v\nactual: %+v", expect, actual)
		}
	}

	return nil
}
