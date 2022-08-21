package collector

import (
	"hwstatexporter/data"
	"math/rand"
	"time"
)

func Collector(store *data.ExportData) {
	timer := 1
	rand.Seed(time.Now().UnixNano())
	for timer != 0 {
		time.Sleep(5 * time.Second)
		store.GpuPowerWatts = rand.Float64()
		store.GpuTempC = rand.Float64()
	}
}
