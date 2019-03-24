package factory

import (
	"go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/entity"
	"go-ddd/backend/domain/error"
	vo "go-ddd/backend/domain/valueobject"
	"time"
)

// CreateNoticeAggregate ...
func CreateNoticeAggregate(uid vo.UniqueID, title, detail string, noticeSeverity, publishType int, from, to *time.Time) (aggregate.Notice, error.ApplicationError) {
	attribute, err := entity.NewNoticeNoticeAttribute(uid, title, detail)
	if err != nil {
		return nil, err
	}
	severity := vo.NewNoticeSeverity(noticeSeverity)
	publishControl := vo.NewPublishControl(vo.NewPublishType(publishType), vo.NewPublishTerm(from, to))

	return aggregate.NewNotice(attribute, severity, publishControl), nil
}
