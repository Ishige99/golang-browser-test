package main

import (
	"fmt"
	"github.com/sclevine/agouti"
)

func main() {
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

	// todo: コマンドラインから出勤と退勤を選択する

	if err := akashiAttendance(page); err != nil {
		fmt.Printf("attendance error: %s\n", err)
	}
	fmt.Println("success attendance.")

	//if err := akashiLeaving(page); err != nil {
	//	fmt.Printf("leaving error: %s\n", err)
	//}
	//fmt.Println("success leaving.")
}
