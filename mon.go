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
	for {
		// Define configuration file
		file, _ := os.Open("config.json")
		decoder := json.NewDecoder(file)
		configuration := Configuration{}

		err := decoder.Decode(&configuration)
		if err != nil {
			fmt.Println("Error:", err)
		}
		// close define Configuration


		t := time.Now().Format("2006-01-02T15:04:05.999999")

		mem_total, mem_usage, mem_cache := collections.Memory()
		rx, tx := collections.Bandwidth(configuration.Interface)

		doc := host{
			HostName:      configuration.HostName,
			Address:       configuration.Address,
			Timestamp:     t,
			Location:      configuration.Location,
			CpuPercent:    collections.CPU()[0],
			MemoryTotal:   mem_total,
			MemoryUsage:   mem_usage,
			MemoryCache:   mem_cache,
			BandWidthRec:  rx,
			BandwidthSent: tx,
		}
		fmt.Printf("doc:%v ", doc)
		// Create connection
		client, err := elastic.NewClient(
			elastic.SetURL(configuration.ElasticURL[0]),
		)

		if err != nil {
			fmt.Printf("Connection fail!...")
		}

		// Add document to the index
		t2 := time.Now()
		_, err = client.Index().
			Index(configuration.Application).
			Type(configuration.HostName).
			Id(strconv.FormatInt(t2.UnixNano(), 10)).
			BodyJson(doc).
			Do()
		if err != nil {
			fmt.Printf("Error create index")
			panic(err)
		}

		fmt.Printf("Time: %v \n", strconv.FormatInt(t2.UnixNano(), 10))

	}
}
