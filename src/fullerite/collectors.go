package main

import (
	"fullerite/collector"
	"fullerite/metric"
	"log"
	"time"
)

func startCollectors(c Config) (collectors []collector.Collector) {
	log.Println("Starting collectors...")
	for _, name := range c.Collectors {
		collectors = append(collectors, startCollector(name))
	}
	return collectors
}

func startCollector(name string) collector.Collector {
	log.Println("Starting collector", name)
	collector := collector.New(name)
	readCollectorConfig(collector)
	go runCollector(collector)
	return collector
}

func readCollectorConfig(collector collector.Collector) {
	// TODO: actually read from configuration file.
	collector.SetInterval(10)
}

func runCollector(collector collector.Collector) {
	for {
		log.Println("Collecting from", collector)
		collector.Collect()
		time.Sleep(time.Duration(collector.Interval()) * time.Second)
	}
}

func readFromCollectors(collectors []collector.Collector, metrics chan metric.Metric) {
	for _, collector := range collectors {
		go readFromCollector(collector, metrics)
	}
}

func readFromCollector(collector collector.Collector, metrics chan metric.Metric) {
	for metric := range collector.Channel() {
		metrics <- metric
	}
}
