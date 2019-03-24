package aggregate

import (
	ag "go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/entity"
	vo "go-ddd/backend/domain/valueobject"
)

type Notice interface {
}

type notice struct {
	noticeAttribute entity.NoticeAttribute
	// 重要度
	severity vo.NoticeSeverity
	// 公開設定
	publishControl *ag.PublishControl
}
