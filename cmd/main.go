package main

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/codingconcepts/dynoapi"
	"github.com/codingconcepts/env"
)

var (
	buildVersion   string
	buildTimestamp string
)

func main() {
	config := struct {
		Host     string `env:"HOST" required:"true"`
		Port     int    `env:"PORT" required:"true"`
		SSL      bool   `env:"SSL" default:"true"`
		CertsDir string `env:"CERTS" default:"certs"`
	}{}
	err := env.Set(&config)
	if err != nil {
		log.Fatalf("loading environment configuration: %v", err)
	}

	routeConfig := flag.String("c", "", "route configuration file")
	flag.Parse()

	var configuration dynoapi.RouteConfigs
	if *routeConfig != "" {
		if configuration, err = loadRouteConfiguration(*routeConfig); err != nil {
			log.Fatalf("loading route configuration: %v", err)
		}
	}

	server := dynoapi.NewServer(
		config.Host,
		config.Port,
		dynoapi.SSL(config.SSL),
		dynoapi.CertsDir(config.CertsDir),
		dynoapi.Routes(configuration...),
		dynoapi.BuildInfo(buildVersion, buildTimestamp))

	server.Start()
}

func loadRouteConfiguration(path string) (routes dynoapi.RouteConfigs, err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	routes = dynoapi.RouteConfigs{}
	err = yaml.Unmarshal(bytes, &routes)
	return
}
