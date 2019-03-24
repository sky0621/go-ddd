package aggregate

import (
	"go-ddd/backend/domain/entity"
	vo "go-ddd/backend/domain/valueobject"
)

// NewNotice ...
func NewNotice(noticeAttribute entity.NoticeAttribute, severity vo.NoticeSeverity, publishControl vo.PublishControl) Notice {
	return &notice{
		noticeAttribute: noticeAttribute,
		severity:        severity,
		publishControl:  publishControl,
	}
}

// Notice ... 「お知らせ」集約情報
type Notice interface {
	GetNoticeAttribute() entity.NoticeAttribute
	GetNoticeSeverity() vo.NoticeSeverity
	GetPublishControl() vo.PublishControl
}

type notice struct {
	// 「お知らせ」属性
	noticeAttribute entity.NoticeAttribute
	// 重要度
	severity vo.NoticeSeverity
	// 公開設定
	publishControl vo.PublishControl
}

// GetNoticeAttribute ...
func (a *notice) GetNoticeAttribute() entity.NoticeAttribute {
	if a == nil {
		return nil
	}
	return a.noticeAttribute
}

// GetNoticeSeverity ...
func (a *notice) GetNoticeSeverity() vo.NoticeSeverity {
	if a == nil {
		return nil
	}
	return a.severity
}

// GetPublishControl ...
func (a *notice) GetPublishControl() vo.PublishControl {
	if a == nil {
		return nil
	}
	return a.publishControl
}
