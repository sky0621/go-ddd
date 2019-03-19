package usecase

import (
	"context"
	"go-ddd/backend/domain/entity"
	"go-ddd/backend/domain/error"
	"go-ddd/backend/domain/repository"
	"go-ddd/backend/usecase/input"
)

// NewNoticeUsecase ...
func NewNoticeUsecase(commandRepository repository.NoticeCommandRepository, queryRepository repository.NoticeQueryRepository) NoticeUsecase {
	return &noticeUsecase{}
}

// NoticeUsecase ... 「お知らせ」に対するユースケース定義
type NoticeUsecase interface {
	AddNotice(context.Context, *input.Notice) error.ApplicationError
}

type noticeUsecase struct {
	commandRepository repository.NoticeCommandRepository
	queryRepository   repository.NoticeQueryRepository
}

// AddNotice ...
func (u *noticeUsecase) AddNotice(ctx context.Context, input *input.Notice) error.ApplicationError {
	var m *entity.Notice
	// FIXME: input -> model.Notice へのコンバーターを実装！
	return u.commandRepository.Create(ctx, m)
}
