package services

import (
	"fmt"
	pegasusHttp "git.jiayincloud.com/TestDev/pegasus.git/components/http"
	"git.jiayincloud.com/TestDev/pegasus.git/core"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"sort"
	"strconv"
	"sync"
)

type AmmoDepot struct {
	Ammos []pegasusHttp.RequestData
	Sync  sync.Mutex
}

var ammos AmmoDepot

func Load(c *gin.Context) {
	var data []pegasusHttp.RequestData
	err := c.BindJSON(&data)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
	ammos.Sync.Lock()
	for _, v := range data {
		ammos.Ammos = append(ammos.Ammos, v)
	}
	ammos.Sync.Unlock()
	c.String(http.StatusOK, "Lock & Load")
}

func GetAmmos(c *gin.Context) {
	c.String(http.StatusOK, strconv.Itoa(len(ammos.Ammos)))
}

func DropAmmos(c *gin.Context) {
	ammos.Ammos = ammos.Ammos[:0]
	c.String(http.StatusOK, "Cleared")
}

func Fire(c *gin.Context) {
	workers, _ := strconv.Atoi(c.Query("workers"))
	rate, _ := strconv.Atoi(c.Query("rate"))
	defer func() {
		ammos.Ammos = ammos.Ammos[:0]
	}()
	if len(ammos.Ammos) == 0 {
		c.String(http.StatusNotFound, "Need to load first")
		return
	}
	jobs := make(chan func(), workers)
	results := make(chan map[string]interface{}, len(ammos.Ammos))

	core.InitWorkerPool(jobs, workers)
	core.InitJobs(rate, jobs, ammos.Ammos, results)

	var elapsed []int
	var endTimes []int
	var statuses []int
	total := 0
	tps := make(map[int]int)
	status := make(map[int]int)
	for i := 0; i < len(ammos.Ammos); i++ {
		result := <-results
		elapsed = append(elapsed, result["elapsed"].(int))
		endTimes = append(endTimes, result["endTime"].(int))
		statuses = append(statuses, result["statusCode"].(int))
	}
	sort.Ints(elapsed)
	for _, respTime := range elapsed {
		total += respTime
	}
	for _, endTime := range endTimes {
		endTimeSecond := endTime / int(math.Pow10(3))
		if tps[endTimeSecond] == 0 {
			tps[endTimeSecond] = 1
		} else {
			tps[endTimeSecond] += 1
		}
	}
	for _, v := range statuses {
		if status[v] == 0 {
			status[v] = 1
		} else {
			status[v] += 1
		}
	}

	report := Report{
		All:    elapsed,
		Tps:    tps,
		Status: status,
	}
	c.JSON(http.StatusOK, report)
}
