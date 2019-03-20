package entity

import (
	"errors"
	ag "go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/enum"
	"go-ddd/backend/domain/error"
)

// Notice ... 「お知らせ」データ定義
type Notice struct {
	// ユニークに特定するID
	id enum.UniqueID
	// 概要を示すタイトル
	title string
	// 詳細
	detail string
	// 重要度
	severity enum.NoticeSeverity
	// 公開設定
	publishControl *ag.PublishControl
}

// NewNotice ...
func NewNotice(id enum.UniqueID, title, detail string, severity enum.NoticeSeverity, publishControl *ag.PublishControl) (*Notice, error.ApplicationError) {
	if id == "" {
		return nil, error.CreateValidationError(error.Required, errors.New("id is required"))
	}
	// FIXME: その他バリデーション

	return &Notice{
		id:             id,
		title:          title,
		detail:         detail,
		severity:       severity,
		publishControl: publishControl,
	}, nil
}

// NoticeCommandCondition ... 条件に該当する「お知らせ」データを決定するために利用
type NoticeCommandCondition struct {
	// FIXME 複数の Notice を持つなり、ページング情報を持つなり、要件にあった取得条件を定義！
}

// NoticeReadCondition ... 条件に該当する「お知らせ」データを決定するために利用
type NoticeReadCondition struct {
	// FIXME 複数の Notice を持つなり、ページング情報を持つなり、要件にあった取得条件を定義！
}
