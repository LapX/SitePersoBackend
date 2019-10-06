package main

import (
	"github.com/LapX/SitePersoBackend/rest"
	_ "github.com/lib/pq"
)

func main() {
	rest.InitServer()
}
