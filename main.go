package main

import (
	"github.com/alkaidchen/invitation-card/app"
	"log"
)

// @title           Marston invitation card backend API
// @version         1.0
// @description     This is a Marston invitation card backend server.
// @termsOfService  http://megrez.run

// @contact.name   Marston invitation card
// @contact.url    http://megrez.run
// @contact.email  alkaidchen@qq.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

func main() {
	a := app.New()
	if err := a.Init(); err != nil {
		log.Fatal("application init failed, ", err)
	}
	if err := a.Run(); err != nil {
		log.Fatal("application run failed, ", err)
	}
}
