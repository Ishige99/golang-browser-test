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

	// オプションの選択
	akashiOption, err := selectOption()
	if err != nil {
		fmt.Printf("select option error: %s\n", err)
	}

	var successMessage string

	switch akashiOption {
	case attendanceOptionNumber:
		if err := akashiAttendance(page); err != nil {
			fmt.Printf("attendance error: %s\n", err)
		}
		successMessage = "success attendance."
	case leavingOptionNumber:
		if err := akashiLeaving(page); err != nil {
			fmt.Printf("leaving error: %s\n", err)
		}
		successMessage = "success leaving."
	}

	fmt.Println(successMessage)
}
