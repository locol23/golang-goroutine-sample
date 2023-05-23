package main

import (
	"log"
	"sync"
	"time"

	"github.com/locol23/golang-goroutine-sample/data"
)

func findElement(target string, targetArr []string) bool {
	for _, v := range targetArr {
		if target == v {
			return true
		}
	}

	return false
}

type args struct {
	getData   func() []string
	wg        *sync.WaitGroup
	mu        *sync.Mutex
	resultArr *[]string
}

// この関数を変更してください。
// data.GetDataA
// data.GetDataB
// data.GetDataC
// 上記3つの関数を呼び、データを受け取って、それを戻り値に設定してください。
// この時、データが重複しないように重複チェックを行ったうえで返却してください。
// go run main.goで確認できます。
// 成功した場合はSUCCESS!!!が表示されます。
// それ以外の時はエラーが表示されます。
func makeData() []string {
	resultArr := make([]string, 0)
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(3)

	fetch := func(args args) {
		resultArr := args.resultArr
		wg := args.wg
		mu := args.mu
		defer wg.Done()
		arr := args.getData()

		for _, v := range arr {
			mu.Lock()

			if !findElement(v, *resultArr) {
				*resultArr = append(*resultArr, v)
			}

			mu.Unlock()
		}
	}

	go fetch(args{data.GetDataA, &wg, &mu, &resultArr})
	go fetch(args{data.GetDataB, &wg, &mu, &resultArr})
	go fetch(args{data.GetDataC, &wg, &mu, &resultArr})

	wg.Wait()
	return resultArr
}

// main関数等はノータッチで
func main() {
	s := time.Now()
	res := makeData()
	d := time.Since(s)

	err := data.Check(res, d)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("SUCCESS!!!")
}
