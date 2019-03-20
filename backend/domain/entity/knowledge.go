package entity

import (
	"errors"
	ag "go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/error"
	vo "go-ddd/backend/domain/valueobject"
)

// Knowledge ... 「ナレッジ」データ定義
type Knowledge interface {
	GetID() vo.UniqueID
	GetTitle() string
	GetDetail() string
	GetPublishControl() *ag.PublishControl
}

type knowledge struct {
	// ユニークに特定するID
	id vo.UniqueID
	// 概要を示すタイトル
	title string
	// 詳細
	detail string
	// 公開設定
	publishControl *ag.PublishControl
}

// NewKnowledge ...
func NewKnowledge(uid vo.UniqueID, title, detail string, publishControl *ag.PublishControl) (Knowledge, error.ApplicationError) {
	if uid.GetVal() == "" {
		return nil, error.CreateValidationError(error.Required, errors.New("id is required"))
	}
	// FIXME: その他バリデーション

	return &knowledge{
		id:             uid,
		title:          title,
		detail:         detail,
		publishControl: publishControl,
	}, nil
}

// GetID ...
func (k *knowledge) GetID() vo.UniqueID {
	if k == nil {
		return nil
	}
	return k.id
}

// GetTitle ...
func (k *knowledge) GetTitle() string {
	if k == nil {
		return ""
	}
	return k.title
}

// GetDetail ...
func (k *knowledge) GetDetail() string {
	if k == nil {
		return ""
	}
	return k.detail
}

// GetPublishControl ...
func (k *knowledge) GetPublishControl() *ag.PublishControl {
	if k == nil {
		return nil
	}
	return k.publishControl
}
