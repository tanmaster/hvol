package main

import (
	"fmt"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/itchyny/volume-go"
	"log"
	"net"
)

func main() {
	vol, err := volume.GetVolume()

	// create an accessory
	info := accessory.Info{Name: "PC Volume"}
	ac := accessory.NewLightbulb(info)

	brightness := characteristic.NewBrightness().Characteristic
	ac.Lightbulb.AddCharacteristic(brightness)
	brightness.UpdateValue(vol)

	brightness.OnValueUpdateFromConn(func(conn net.Conn, c *characteristic.Characteristic, newValue, oldValue interface{}) {
		err = volume.SetVolume(newValue.(int))
		if err != nil {
			log.Fatalf("set volume failed: %+v", err)
		}
		fmt.Printf("set volume success\n")
	})

	// configure the ip transport
	config := hc.Config{Pin: "12344321"}
	t, err := hc.NewIPTransport(config, ac.Accessory)
	if err != nil {
		log.Panic(err)
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
