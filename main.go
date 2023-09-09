package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"log"
)

func main() {
	// chromeドライバを使って、Google Chromeを使用できるようにする。
	agoutiDriver := agouti.ChromeDriver()
	if err := agoutiDriver.Start(); err != nil {
		fmt.Printf("Failed to start driver. %s\n", err)
		return
	}
	defer agoutiDriver.Stop()

	// agoutiを使用して新規ブラウザを立ち上げる。このブラウザを使って色々操作します。
	page, err := agoutiDriver.NewPage()
	if err != nil {
		fmt.Printf("Failed to open a new page. %s\n", err)
		return
	}
}
