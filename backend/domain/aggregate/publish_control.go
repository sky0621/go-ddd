package aggregate

import (
	"go-ddd/backend/domain/enum"
)

// Publish ...
func Publish(term *PublishTerm) *PublishControl {
	return &PublishControl{
		isPublish: enum.GeneralPublic,
		term:      term,
	}
}

// Private ...
func Private() *PublishControl {
	return &PublishControl{
		isPublish: enum.Private,
	}
}

// PublishControl ... 特定情報の公開制御
type PublishControl struct {
	// IsPublish ... 公開フラグ
	isPublish enum.PublishType
	// Term ... 公開期間
	term *PublishTerm
}
