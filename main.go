package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

var cpu_work, cpu_total int64

func getCpu() (load int64) {
	b, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		panic(fmt.Sprintf("Error opening /proc/stat", err))
	}

	var jif1, jif2, jif3, jif4, jif5, jif6, jif7, work, total int64

	_, err = fmt.Sscanf(string(b), "cpu %d %d %d %d %d %d %d 0 0 0\n", &jif1, &jif2, &jif3, &jif4, &jif5, &jif6, &jif7)
	if err != nil {
		panic(fmt.Sprintf("Error scanning /proc/stat", err))
	}

	work = jif1 + jif2 + jif3 + jif6 + jif7
	total = work + jif4 + jif5

	load = 100 * (work - cpu_work) / (total - cpu_total)

	cpu_work = work
	cpu_total = total

	return load
}

func getMem() (mem int64) {
	b, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		panic(fmt.Sprintf("Error opening /proc/meminfo", err))
	}

	var total, free, avail, buf, cache, used int64

	_, err = fmt.Sscanf(string(b), "MemTotal: %d kB\nMemFree: %d kB\nMemAvailable: %d kB\nBuffers: %d kB\nCached: %d kB\n", &total, &free, &avail, &buf, &cache)
	if err != nil {
		panic(fmt.Sprintf("Error scanning /proc/meminfo", err))
	}
	used = 100 * (total - free - buf - cache) / total
	return used
}

//func getWeather() (s string) {
//	b, err := ioutil.ReadFile("/tmp/weather.txt")
//	if err != nil {
//		panic(fmt.Sprintf("Error opening /tmp/weather.txt", err))
//	}
//	return strings.TrimSpace(fmt.Sprintf("%s", b))
//}

func getDate() (s string) {
	t := time.Now().Format("Mon Jan _2 15:04")
	return strings.TrimSpace(fmt.Sprintf("%s", t))
}

func main() {
	for {
		m := getMem()
		c := getCpu()
		//w := getWeather()
		t := getDate()
		t = strings.Replace(t, "  ", " ", -1)
		fmt.Printf("cpu:%2d | mem:%2d | %s\n", c, m, t)
		time.Sleep(2 * 1e9)
	}
}
