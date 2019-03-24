package aggregate

import (
	ag "go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/entity"
	vo "go-ddd/backend/domain/valueobject"
)

type Notice interface {
	GetNoticeAttribute() entity.NoticeAttribute
	GetNoticeSeverity() vo.NoticeSeverity
}

type notice struct {
	// 「お知らせ」属性
	noticeAttribute entity.NoticeAttribute
	// 重要度
	severity vo.NoticeSeverity
	// 公開設定
	publishControl *ag.PublishControl
}
