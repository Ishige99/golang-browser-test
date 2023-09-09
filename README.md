## 概要

AKASHIの勤怠ボタンを自動で押してくれるやつ

## 開発環境

- macOS Monterey(12.2.1)
- Goland
- go version go1.20 darwin/arm64

## 環境変数

`.env`

```
# AKASHIログイン情報
FORM_COMPANY_ID=""
FORM_LOGIN_ID=""
FORM_PASSWORD=""
```

## 使用方法

```bash
% go build golang-browser-test
% go run golang-browser-test
select options
2: 退勤
1: 出勤
> // 実行したいオプションを入力
```

実行後にブラウザが自動で立ち上がって勤怠のボタンを押します。