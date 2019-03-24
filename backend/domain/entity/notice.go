package entity

import (
	"errors"
	ag "go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/error"
	vo "go-ddd/backend/domain/valueobject"
)

// NoticeAttribute ... 「お知らせ」データ定義
type NoticeAttribute interface {
}

type noticeAttribute struct {
	// ユニークに特定するID
	id vo.UniqueID
	// 概要を示すタイトル
	title string
	// 詳細
	detail string
}

// NewNotice ...
func NewNotice(uid vo.UniqueID, title, detail string, severity vo.NoticeSeverity, publishControl *ag.PublishControl) (*Notice, error.ApplicationError) {
	if uid.GetVal() == "" {
		return nil, error.CreateValidationError(error.Required, errors.New("id is required"))
	}
	// FIXME: その他バリデーション

	return &Notice{
		id:             uid,
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
