package main

import (
	"encoding/json"
	"fmt"
	"github.com/bienkma/monagent/collections"
	"github.com/olivere/elastic"
	"os"
	"time"
)

// Define the config file
type Configuration struct {
	HostName    string
	Address     string
	Interface   string
	Application string // Example orm, antispam, bigdata
	Location    string
	ElasticURL  []string
}

// Define struct for doc in elasticsearch
type host struct {
	HostName      string    `json: "hostname"`
	Address       string    `json: "address"`
	TimeStamp     time.Time `json: "timestamp"`
	Location      string    `json: "location"`
	CpuPercent    float64   `json: "cpu_usage"`
	MemoryTotal   float64   `json: "memory_total"`
	MemoryUsage   float64   `json: "memory_usage"`
	MemoryCache   float64   `json: "memory_usage"`
	BandWidthRec  float64   `json: "net_in"`
	BandwidthSent float64   `json: "net_out"`
}

// The main function
func main() {

	// Define configuration file
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}

	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// close define Configuration

	t := time.Now().UTC()

	doc := host{
		HostName:      configuration.HostName,
		Address:       configuration.Address,
		TimeStamp:     t,
		Location:      configuration.Location,
		CpuPercent:    collections.CPU(),
		MemoryTotal:   collections.Memory(),
		MemoryUsage:   512,
		MemoryCache:   256,
		BandWidthRec:  10,
		BandwidthSent: 20,
	}

	// Create connection
	client, err := elastic.NewClient(
		elastic.SetURL(configuration.ElasticURL[0]),
	)

	if err != nil {
		fmt.Printf("Connection fail!...")
	}

	// Add document to the index
	_, err = client.Index().
		Index(configuration.Application).
		Type(configuration.HostName).
		Id(strconv.FormatInt(t.UnixNano(), 10)).
		BodyJson(doc).
		Do()
	if err != nil {
		fmt.Printf("Error create index")
		panic(err)
	}

	fmt.Printf("Time: %v \n", strconv.FormatInt(t.UnixNano(), 10))

}
