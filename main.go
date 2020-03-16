package main

import (
	"github.com/kiki-ki/cal-app-go/api/router"
	"github.com/kiki-ki/cal-app-go/src/database"
)

func main() {
	database.New()
	r := router.New()
	r.Run()
}
