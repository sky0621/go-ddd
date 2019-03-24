package entity

import (
	"errors"
	"go-ddd/backend/domain/error"
	vo "go-ddd/backend/domain/valueobject"
)

// NewNoticeNoticeAttribute ...
func NewNoticeNoticeAttribute(uid vo.UniqueID, title, detail string) (NoticeAttribute, error.ApplicationError) {
	if uid.GetVal() == "" {
		return nil, error.CreateValidationError(error.Required, errors.New("id is required"))
	}
	if title == "" {
		return nil, error.CreateValidationError(error.Required, errors.New("title is required"))
	}

	return &noticeAttribute{
		id:     uid,
		title:  title,
		detail: detail,
	}, nil
}

// NoticeAttribute ... 「お知らせ」属性
type NoticeAttribute interface {
	GetID() vo.UniqueID
	GetTitle() string
	GetDetail() string
}

type noticeAttribute struct {
	// ユニークに特定するID
	id vo.UniqueID
	// 概要を示すタイトル
	title string
	// 詳細
	detail string
}

// GetID ...
func (a *noticeAttribute) GetID() vo.UniqueID {
	if a == nil {
		return nil
	}
	return a.id
}

// GetTitle ...
func (a *noticeAttribute) GetTitle() string {
	if a == nil {
		return ""
	}
	return a.title
}

// GetDetail ...
func (a *noticeAttribute) GetDetail() string {
	if a == nil {
		return ""
	}
	return a.detail
}
