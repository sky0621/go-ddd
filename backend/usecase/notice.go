package usecase

import (
	"context"
	"go-ddd/backend/domain/error"
	"go-ddd/backend/domain/factory"
	"go-ddd/backend/domain/repository"
	vo "go-ddd/backend/domain/valueobject"
	"go-ddd/backend/usecase/input"
)

// NewNotice ...
func NewNotice(command repository.NoticeCommandRepository, query repository.NoticeQueryRepository) NoticeUsecase {
	return &noticeUsecase{
		command: command,
		query:   query,
	}
}

// NoticeUsecase ... 「お知らせ」に対するユースケース定義
type NoticeUsecase interface {
	AddNotice(context.Context, *input.Notice) error.ApplicationError
}

type noticeUsecase struct {
	command repository.NoticeCommandRepository
	query   repository.NoticeQueryRepository
}

// AddNotice ...
func (u *noticeUsecase) AddNotice(ctx context.Context, in *input.Notice) error.ApplicationError {
	f := factory.NoticeFactory{uid: vo.NewUniqueID()}
	n, err := f.CreateNotice(in.Title, in.Detail, in.NoticeSeverity, in.PublishType, in.From, in.To)
	if err != nil {
		return err
	}
	return u.command.Create(ctx, n)
}
