package main

import (
	"encoding/json"
	"fmt"
	"github.com/bienkma/monagent/collections"
	"github.com/olivere/elastic"
	"os"
	"strconv"
	"time"
)

// Define the config file
type Configuration struct {
	HostName    string
	Address     string
	Interface   string
	Application string
	Location    string
	ElasticURL  []string
	Interval    uint
}

// Define struct for doc in elasticsearch
type host struct {
	HostName      string  `json: "hostname"`
	Address       string  `json: "address"`
	Timestamp     string  `json: "timestamp"`
	Location      string  `json: "location"`
	CpuPercent    float64 `json: "cpu_usage"`
	MemoryTotal   uint64  `json: "memory_total"`
	MemoryUsage   uint64  `json: "memory_usage"`
	MemoryCache   uint64  `json: "memory_usage"`
	BandWidthRec  float64 `json: "net_in"`
	BandwidthSent float64 `json: "net_out"`
}

// The main function
func main() {
	// Define configuration file
	file, err_opening := os.Open("config.json")
	if err_opening != nil {
		collections.Log("config.json file not found")
		panic(err_opening)
	}
	decoder := json.NewDecoder(file)
	configuration := Configuration{}

	err := decoder.Decode(&configuration)
	if err != nil {
		collections.Log("config.json file wrong format!..")
		panic(err)
	}
	// close define Configuration
	// Create connection
	client, err := elastic.NewClient(
		elastic.SetURL(configuration.ElasticURL[0]),
	)
	if err != nil {
		collections.Log("Can't connection elasticsearch!...")
		panic(err)
	}

	for {
		t := time.Now().Format(time.RFC3339Nano)
		mem_total, mem_usage, mem_cache := collections.Memory()
		rx, tx := collections.Bandwidth(configuration.Interval, configuration.Interface)
		// create doc for management host
		doc := host{
			HostName:      configuration.HostName,
			Address:       configuration.Address,
			Timestamp:     t,
			Location:      configuration.Location,
			CpuPercent:    collections.CPU(configuration.Interval)[0],
			MemoryTotal:   mem_total,
			MemoryUsage:   mem_usage,
			MemoryCache:   mem_cache,
			BandWidthRec:  rx,
			BandwidthSent: tx,
		}

		// Add document to the index
		t2 := strconv.FormatInt(time.Now().UnixNano(), 10)
		_, err = client.Index().
			Index(configuration.Application).
			Type(configuration.HostName).
			Id(t2).
			BodyJson(doc).
			Do()
		if err != nil {
			collections.Log("Can't create index in elasticsearch. Maybe duplicate index or problem connection elastic server!.. ")
			panic(err)
		}

		fmt.Printf("Time: %v \n", t2)
	}
}
