package enum

const (
	// SystemAdministrator ... システム管理者
	SystemAdministrator Role = iota + 1

	// OrganizationManager ... 組織管理者
	OrganizationManager

	// AuthenticationUser ... 認証ユーザー
	AuthenticationUser

	// GuestUser ... ゲストユーザー
	GuestUser
)

// Role ... ユーザーの役割
type Role int8
