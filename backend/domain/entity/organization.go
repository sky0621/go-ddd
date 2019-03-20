package entity

import vo "go-ddd/backend/domain/valueobject"

// Organization ... 「組織」データ定義
type Organization struct {
	// ユニークに特定するID
	id vo.UniqueID

	// 組織名称
	logicalName string
}
