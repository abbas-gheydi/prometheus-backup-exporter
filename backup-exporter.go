package main

import (
	"backup_exporter/checkfiles"
	confs "backup_exporter/config"
	"backup_exporter/prommetrics"
	"log"
	"net/http"
	"sync"
	"fmt"
)

var (
	configuration confs.Configurations
)
//RunExporter expose metrics to web server
func RunExporter(res http.ResponseWriter, req *http.Request) {

	var wg sync.WaitGroup
	wg.Add(len(configuration.DailyBackups))
	metricPrefix := "backup"
	sizeMetricName := metricPrefix + "_size"
	dateMetricName := metricPrefix + "_date"
	go func() {

		writeM := func(mType string, m string) {
			fmt.Fprintf(res,mType)
			fmt.Fprintf(res,m)

		}
		for _, bkAddress := range configuration.DailyBackups {

			pathName, size, date := checkfiles.Run(bkAddress)
			writeM(prommetrics.Generate(sizeMetricName, metricPrefix,pathName, size))
			writeM(prommetrics.Generate(dateMetricName, metricPrefix,pathName, date))

			wg.Done()
		}

	}()

	wg.Wait()

}

func main() {
	confs.ReadConfs(&configuration,"config")
	mux := http.NewServeMux()
	mux.HandleFunc("/metrics", RunExporter)
	log.Print("Server is running on: ", configuration.Server.Port)
	log.Fatal(http.ListenAndServe(":"+configuration.Server.Port, mux))

}
