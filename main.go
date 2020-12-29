package main

import (
	"fmt"
	"log"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/itchyny/volume-go"
)

func main() {
	vol, err := volume.GetVolume()
	if err != nil {
		log.Fatalf("get volume failed: %+v", err)
	}
	fmt.Printf("current volume: %d\n", vol)

	err = volume.SetVolume(10)
	if err != nil {
		log.Fatalf("set volume failed: %+v", err)
	}
	fmt.Printf("set volume success\n")

	err = volume.Mute()
	if err != nil {
		log.Fatalf("mute failed: %+v", err)
	}

	err = volume.Unmute()
	if err != nil {
		log.Fatalf("unmute failed: %+v", err)
	}

	// create an accessory
	info := accessory.Info{Name: "Lamp"}
	ac := accessory.NewSwitch(info)

	// configure the ip transport
	config := hc.Config{Pin: "00102003"}
	t, err := hc.NewIPTransport(config, ac.Accessory)
	if err != nil {
		log.Panic(err)
	}

	hc.OnTermination(func(){
		<-t.Stop()
	})

	t.Start()
}
