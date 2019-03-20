package entity

import "go-ddd/backend/domain/enum"

// Organization ... 「組織」データ定義
type Organization struct {
	// ユニークに特定するID
	id enum.UniqueID

	// 組織名称
	logicalName string
}
