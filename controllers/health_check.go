package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"

	"github.com/apinanyogaratnam/golang-vercel-service-template/models"
)

func TestHttp() error {
	client := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	resp, err := client.Get("http://example.com")

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func printSystemResources() {
	fmt.Printf("GOOS: %s\n", runtime.GOOS)        // prints the operating system name
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU()) // prints the number of logical CPUs usable by the current process
	fmt.Printf("GOROOT: %s\n", runtime.GOROOT()) // prints the root of the Go tree

	pwd, err := os.Getwd() // get the current working directory
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	fmt.Printf("PWD: %s\n", pwd)

	// Get memory stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("\nAlloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\nTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\nSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\nNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func GetSystemLoad() (float64, float64) {
	v, _ := mem.VirtualMemory()

	// Get CPU usage
	percent, _ := cpu.Percent(time.Second, false)

	cpuLoad := percent[0]
	memLoad := v.UsedPercent

	return cpuLoad, memLoad
}

func GetHealthStatus() models.HealthCheckResponse {
	healthStatus := models.HealthCheckResponse{
		Status: "OK",
		Load: models.Load{
			Cpu: 0.0,
			Mem: 0.0,
		},
	}

	err := TestHttp()
	if err != nil {
		healthStatus.Status = "ERROR"
	}

	cpuLoad, memLoad := GetSystemLoad()
	healthStatus.Load.Cpu = cpuLoad
	healthStatus.Load.Mem = memLoad

	return healthStatus
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	printSystemResources()

	healthStatus := GetHealthStatus()

	// Here you can add checks for your database, external services, etc.
	// If everything is okay, respond with a 200 status. Otherwise, respond with a 500 status and include more information.

	json.NewEncoder(w).Encode(healthStatus)
}
