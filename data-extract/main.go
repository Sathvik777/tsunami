package main

import (
	"fmt"
	"log"
	"net/http"
)

// URL is Prometheus api endpoint for query
const URL = "http://localhost:9090/api/v1/query?query=%s"

// PROMQL1 query for Memory Query
const PROMQL1 = "{'query':'(node_memory_MemTotal_bytes - (node_memory_MemFree_bytes + node_memory_Cached_bytes + node_memory_Buffers_bytes)) / node_memory_MemTotal_bytes * 100'}"

// PROMQL2 query for  CPU Query
const PROMQL2 = "{'query':'100 - avg(irate(node_cpu_seconds_total{job='node',mode='idle'}[5m])) by (instance) * 100'}"

func main() {

}

func metric(query string) {
	reqURL := fmt.Sprintf(URL, query)
	_, err := http.Get(reqURL)
	if err != nil {
		log.Fatal("Could not query Prometheus")
	}
}
