package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sclevine/agouti"
	"os"
)

const (
	attendanceOptionNumber = 1
	leavingOptionNumber    = 2
)

func selectOption() (int, error) {
	var optionMap = map[int]string{
		attendanceOptionNumber: "出勤",
		leavingOptionNumber:    "退勤",
	}

	// 実行したいオプションを入力してもらう
	var option int
	fmt.Printf("select options \n")
	for k, v := range optionMap {
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Printf(">")
	fmt.Scan(&option)

	// 入力されたオプションが正しいか判定
	if _, ok := optionMap[option]; !ok {
		return 0, fmt.Errorf("not the correct option value")
	}

	return option, nil
}

func executeAkashiTimeClock(page *agouti.Page, option int) error {
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

	// オプションに応じて勤怠処理の変更
	var xPath string
	switch option {
	case attendanceOptionNumber:
		xPath = "/html/body/div[1]/div/section/form/div[2]/div/div[2]/ul[1]/li[1]/a[@data-punch-type='attendance']"
	case leavingOptionNumber:
		xPath = "/html/body/div[1]/div/section/form/div[2]/div/div[2]/ul[1]/li[2]/a[@data-punch-type='leaving']"
	}

	button := page.FindByXPath(xPath)
	if err := button.Click(); err != nil {
		return fmt.Errorf("Failed to click button. %s\n", err)
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
	// フォーム情報を取得
	formCompanyId := page.FindByID("form_company_id")
	formLoginId := page.FindByID("form_login_id")
	formPassword := page.FindByID("form_password")

	// 環境変数セット
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Failed loading .env file: %s\n", err)
	}

	companyId := os.Getenv("FORM_COMPANY_ID")
	loginId := os.Getenv("FORM_LOGIN_ID")
	password := os.Getenv("FORM_PASSWORD")

	// フォームにログイン情報を入力
	formCompanyId.Fill(companyId)
	formLoginId.Fill(loginId)
	formPassword.Fill(password)

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
