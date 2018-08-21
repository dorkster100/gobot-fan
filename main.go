package main

import (
        "time"
	"fmt"

        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/gpio"
        "gobot.io/x/gobot/platforms/raspi"
)

func main() {
        adaptor := raspi.NewAdaptor()
        sensor := gpio.NewAnalogSensorDriver(adaptor, "14")

        work := func() {
                sensor.On(aio.Data, func(data interface{}) {
			fmt.Println("sensor", data)
                })
        }

        robot := gobot.NewRobot("snapbot",
                []gobot.Connection{adaptor},
                []gobot.Device{sensor},
                work,
        )

        robot.Start()
}
