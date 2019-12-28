/*
Copyright © 2019 sathvik.katam

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
  "net/http"
  "math/rand"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v2"
)

type serverConfig struct {
	Port int `yaml:"port"`
}

//Config is a yaml replication
type Config struct {
	Server serverConfig `yaml:"server"`
}

var (
	// Create a summary to track fictional interservice RPC latencies for three
	// distinct services with different latency distributions. These services are
	// differentiated via a "service" label.
	rpcDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "rpc_durations_seconds",
			Help:       "RPC latency distributions.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"service"},
	)
)

func setUpConfig() Config {
	var config Config
	configYaml, err := ioutil.ReadFile(".config.yaml")
	if err != nil {
		log.Fatalf("failed to find config file : %s", err)
	}
	if err = yaml.Unmarshal(configYaml, &config); err != nil {
		log.Fatalf("failed to unmarshal config file : %s", err)
	}
	return config
}


func init() {
	// Register the summary and the histogram with Prometheus's default registry.
	prometheus.MustRegister(rpcDurations)
	// Add Go module build info.
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
}


func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func loadListen(w http.ResponseWriter, r *http.Request) {
  v := rand.Float64()
  rpcDurations.WithLabelValues("uniform").Observe(v)
	w.WriteHeader(http.StatusOK)
}

func main() {
	config := setUpConfig()
	http.HandleFunc("/health", health)
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), nil); err != nil {
		log.Fatal("Serving exited with error")
	}
}
