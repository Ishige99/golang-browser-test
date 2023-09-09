package main

import (
	"github.com/sclevine/agouti"
	"log"
)

func main() {
	agoutiDriver := agouti.ChromeDriver()
	agoutiDriver.Start()
	defer agoutiDriver.Stop()
	page, _ := agoutiDriver.NewPage()
}
