package entity

import vo "go-ddd/backend/domain/valueobject"

// OrganizationAttribute ... 「組織」データ定義
type OrganizationAttribute struct {
	// ユニークに特定するID
	id vo.UniqueID

	// 組織名称
	organizationName vo.OrganizationName
}
