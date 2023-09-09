package main

import (
	"fmt"
	"github.com/sclevine/agouti"
)

func akashiAttendance(page *agouti.Page) error {
	// AKASHIログインページ移動
	if err := openAkashiLoginPage(page); err != nil {
		return err
	}

	// AKASHIログイン
	if err := loginAkashi(page); err != nil {
		return err
	}

	// 音声ミュート
	if err := muteAkashi(page); err != nil {
		return err
	}

	// 出勤
	attendanceButton := page.FindByXPath("/html/body/div[1]/div/section/form/div[2]/div/div[2]/ul[1]/li[1]/a[@data-punch-type='attendance']")
	if err := attendanceButton.Click(); err != nil {
		return fmt.Errorf("Failed to click attendance button. %s\n", err)
	}

	return nil
}

func openAkashiLoginPage(page *agouti.Page) error {
	// 次に勤怠ページを指定したAKASHIのページを開きます。
	if err := page.Navigate("https://atnd.ak4.jp/login?next=%2Fmypage%2Fpunch"); err != nil {
		return fmt.Errorf("Failed to navigate to AKASHI login page. %s\n", err)
	}
	return nil
}

func loginAkashi(page *agouti.Page) error {
	// AKASHIログインフォームにそれぞれログイン内容を入力。
	formCompanyId := page.FindByID("form_company_id")
	formLoginId := page.FindByID("form_login_id")
	formPassword := page.FindByID("form_password")

	formCompanyId.Fill("company")
	formLoginId.Fill("login_id")
	formPassword.Fill("password")

	// submit
	formSubmitButton := page.FindByID("submit-button")
	if err := formSubmitButton.Click(); err != nil {
		return fmt.Errorf("Failed to click form button. %s\n", err)
	}

	return nil
}

func muteAkashi(page *agouti.Page) error {
	// muteボタンのXPathを指定して音声ミュート
	muteButton := page.FindByXPath("/html/body/div[1]/div/section/form/div[1]/div[2]/ul/li[2]")
	if err := muteButton.Click(); err != nil {
		return fmt.Errorf("Failed to click mute button. %s\n", err)
	}
	return nil
}
