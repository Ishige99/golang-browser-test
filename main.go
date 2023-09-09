package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"log"
)

func main() {
	// オプションの選択
	option, err := selectOption()
	if err != nil {
		fmt.Printf("select option error: %s\n", err)
		return
	}

	// chromeドライバを使って、Google Chromeを使用できるようにする。
	agoutiDriver := agouti.ChromeDriver()
	if err := agoutiDriver.Start(); err != nil {
		fmt.Printf("Failed to start driver. %s\n", err)
		return
	}
	defer agoutiDriver.Stop()

	// agoutiを使用して新規ブラウザを立ち上げる。開いたブラウザを使って色々操作します。
	page, err := agoutiDriver.NewPage()
	if err != nil {
		fmt.Printf("Failed to open a new page. %s\n", err)
		return
	}

	// オプションに応じて勤怠処理の実行
	if err := executeAkashiTimeClock(page, option); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("success")
}
