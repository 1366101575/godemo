package sync

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func TestOnce(t *testing.T) {
	t.Run("test0", func(t *testing.T) {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		wg.Add(3)

		go getDataByKey("bb")
		time.Sleep(1 * time.Second)
		go getDataByKey("cc")
		time.Sleep(1 * time.Second)
		go getDataByKey("ff")

		wg.Wait()
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	})
}

var mapData map[string]string

func initMapData() {
	fmt.Println("initMapData", time.Now().Format("2006-01-02 15:04:05"))

	//mapData = map[string]string{
	//	"aa": loadData("aa"),
	//	"bb": loadData("bb"),
	//	"cc": loadData("cc"),
	//	"dd": loadData("dd"),
	//	"ee": loadData("ee"),
	//	"ff": loadData("ff"),
	//}

	mapData = make(map[string]string)
	mapData["aa"] = loadData("aa")
	mapData["bb"] = loadData("bb")
	mapData["cc"] = loadData("cc")
	mapData["dd"] = loadData("dd")
	mapData["ee"] = loadData("ee")
	mapData["ff"] = loadData("ff")
}

func loadData(k string) string {
	time.Sleep(1 * time.Second)
	return strings.ToUpper(k)
}

var initMapDataOnce sync.Once

func getDataByKey(k string) string {
	fmt.Printf("%s start %s\n", k, time.Now().Format("2006-01-02 15:04:05"))
	//if mapData == nil {
	//	initMapData()
	//}

	initMapDataOnce.Do(initMapData)

	fmt.Printf("%s: %s  time:%s\n", k, mapData[k], time.Now().Format("2006-01-02 15:04:05"))
	wg.Done()
	return mapData[k]
}
