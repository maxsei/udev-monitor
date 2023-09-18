package main

import (
	serialdevicemonitor "github.com/maxsei/serial-device-monitor"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	ctx, err := serialdevicemonitor.NewContext()
	if err != nil {
		log.Panic(err)
	}
	defer ctx.Deinit()
	e, err := serialdevicemonitor.NewEnumerator(ctx)
	if err != nil {
		log.Panic(err)
	}
	defer e.Deinit()
	for device := e.Next(); device != nil; device = e.Next() {
		log.Infof("device enumerated path = %s", device.DeviceNode())
		device.Deinit()
	}

	m, err := serialdevicemonitor.NewMonitor(ctx)
	if err != nil {
		log.Panic(err)
	}
	defer m.Deinit()
	device, err := m.Receive()
	if err != nil {
		log.Panic(err)
	}
	defer device.Deinit()
	devpath := device.DeviceNode()
	log.Infof("got new terminal device at: %s", devpath)
	props := device.Properties()
	for k, v := range props {
		log.Infof("%s: %s", k, v)
	}
	defer m.Deinit()
}
