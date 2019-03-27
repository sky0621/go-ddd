package input

import "time"

// Notice ... ユースケースとしての「お知らせ」情報を定義
type Notice struct {
	// タイトル
	Title string

	// 詳細
	Detail string

	// お知らせの重要度
	NoticeSeverity int

	// 特定情報の公開設定
	PublishType int

	// 公開開始日時
	From *time.Time

	// 公開終了日時
	To *time.Time
}
