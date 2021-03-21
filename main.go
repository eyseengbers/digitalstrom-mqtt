package main

import (
	"fmt"
	"github.com/gaetancollaud/digitalstrom-mqtt/config"
	"github.com/gaetancollaud/digitalstrom-mqtt/digitalstrom"
	"github.com/gaetancollaud/digitalstrom-mqtt/digitalstrom_mqtt"
	"time"
)

func main() {
	fmt.Println("String digitalstrom MQTT!")

	config := config.FromEnv()

	ds := digitalstrom.New(config)
	mqtt := digitalstrom_mqtt.New(config, ds)

	ds.Start()
	mqtt.Start()

	fmt.Println("Waiting forever")

	time.Sleep(100 * 365 * 24 * time.Hour)
}
