package valueobject

import "strings"

// NewOrganizationLogicalName ...
func NewOrganizationLogicalName(logicalName string) OrganizationName {
	return NewOrganizationName(logicalName, "", "")
}

// NewOrganizationName ...
func NewOrganizationName(logicalName, logicalKanaName, physicalName string) OrganizationName {
	// FIXME: logicalName に対する形式チェック！（桁チェック。）
	// FIXME: logicalKanaName に対する形式チェック！（桁チェック。カタカナのみチェック。）
	// FIXME: physicalName に対する形式チェック！（桁チェック。半角英数字記号のみチェック。）
	return &organizationName{
		logicalName:     logicalName,
		logicalKanaName: logicalKanaName,
		physicalName:    physicalName,
	}
}

// OrganizationName ... 組織名
type OrganizationName interface {
	GetOrganizationLogicalName() string
	GetOrganizationLogicalKanaName() string
	GetOrganizationPhysicalName() string
	Equals(comparison OrganizationName) bool
}

type organizationName struct {
	// 論理名
	logicalName string
	// 論理名（カナ）
	logicalKanaName string
	// 物理名
	physicalName string
}

// GetOrganizationLogicalName ...
func (on *organizationName) GetOrganizationLogicalName() string {
	return on.logicalName
}

// GetOrganizationLogicalKanaName ...
func (on *organizationName) GetOrganizationLogicalKanaName() string {
	return on.logicalKanaName
}

// GetOrganizationPhysicalName ...
func (on *organizationName) GetOrganizationPhysicalName() string {
	return on.physicalName
}

// Equals ...
func (on *organizationName) Equals(comparison OrganizationName) bool {
	if on == nil || comparison == nil {
		return false
	}
	return (strings.Compare(on.GetOrganizationLogicalName(), comparison.GetOrganizationLogicalName()) != -1) &&
		(strings.Compare(on.GetOrganizationLogicalKanaName(), comparison.GetOrganizationLogicalKanaName()) != -1) &&
		(strings.Compare(on.GetOrganizationPhysicalName(), comparison.GetOrganizationPhysicalName()) != -1)
}
