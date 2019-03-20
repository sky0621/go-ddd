package entity

import (
	ag "go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/enum"
)

// Knowledge ... 「ナレッジ」データ定義
type Knowledge struct {
	// ユニークに特定するID
	id enum.UniqueID
	// 概要を示すタイトル
	title string
	// 詳細
	detail string
	// 公開設定
	publishControl *ag.PublishControl
}
