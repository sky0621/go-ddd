package enum

const (
	// GeneralPublic ... 一派公開
	GeneralPublic PublishType = iota + 1

	// InHousePublish ... 組織内公開
	InHousePublish

	// Private ... 非公開
	Private
)

// PublishType ... 特定情報の公開設定
type PublishType int8
