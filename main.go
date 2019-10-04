package main

import (
	"github.com/LapX/SitePersoBackend/resource"
	_ "github.com/lib/pq"
)

func main() {
	resource.InitServer()
}
