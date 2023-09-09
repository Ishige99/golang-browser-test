package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"time"
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

	// todo: コマンドラインから出勤と退勤を選択する

	// AKASHIログインページ移動
	if err := page.Navigate("https://atnd.ak4.jp/login?next=%2Fmypage%2Fpunch"); err != nil {
		fmt.Printf("Failed to navigate to AKASHI login page. %s\n", err)
		return
	}

	// フォームに内容を入力して、ログインを行う
	formCompanyId := page.FindByID("form_company_id")
	formLoginId := page.FindByID("form_login_id")
	formPassword := page.FindByID("form_password")

	formCompanyId.Fill("company")
	formLoginId.Fill("login_id")
	formPassword.Fill("password")

	// todo: 後で消す（テスト用スリープ、入力内容確認）
	time.Sleep(2 * time.Second)

	formSubmitButton := page.FindByID("submit-button")
	if err := formSubmitButton.Click(); err != nil {
		fmt.Printf("Failed to navigate to form button click. %s\n", err)
		return
	}

	// todo: 後で消す（テスト用スリープ、入力内容確認）
	time.Sleep(2 * time.Second)
}
