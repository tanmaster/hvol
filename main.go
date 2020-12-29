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
	info := accessory.Info{Name: "Mac Volume"}
	ac := accessory.NewLightbulb(info)

	brightness := characteristic.NewBrightness().Characteristic
	brightness.UpdateValue(vol)
	ac.Lightbulb.AddCharacteristic(brightness)
	ac.Lightbulb.On.SetValue(true)
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
