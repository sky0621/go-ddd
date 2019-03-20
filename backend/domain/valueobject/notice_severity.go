package valueobject

const (
	// NoticeSeverityInfo ... 通常レベルのお知らせ
	NoticeSeverityInfo = iota + 1

	// NoticeSeverityImportant ... 重要レベルのお知らせ
	NoticeSeverityImportant
)

// NewNoticeSeverity ...
func NewNoticeSeverity(val int) NoticeSeverity {
	// 事前条件
	assertion := func(v int) bool {
		if v == NoticeSeverityInfo {
			return true
		}
		if v == NoticeSeverityImportant {
			return true
		}
		return false
	}
	if !assertion(val) {
		return nil
	}

	return &noticeSeverity{val: val}
}

// NewNoticeSeverityInfo ... 通常レベルのお知らせを生成する
func NewNoticeSeverityInfo() NoticeSeverity {
	return &noticeSeverity{val: NoticeSeverityInfo}
}

// NewNoticeSeverityImportant ... 重要レベルのお知らせを生成する
func NewNoticeSeverityImportant() NoticeSeverity {
	return &noticeSeverity{val: NoticeSeverityImportant}
}

// NoticeSeverity ... お知らせの重要度
type NoticeSeverity interface {
	GetVal() int
	Equals(comparison NoticeSeverity) bool
	IsNoticeSeverityInfo() bool
	IsNoticeSeverityImportant() bool
}

type noticeSeverity struct {
	val int
}

// GetVal ...
func (ns *noticeSeverity) GetVal() int {
	return ns.val
}

// Equals ...
func (ns *noticeSeverity) Equals(comparison NoticeSeverity) bool {
	if ns == nil || comparison == nil {
		return false
	}
	return ns.GetVal() == comparison.GetVal()
}

// IsNoticeSeverityInfo ... 通常レベルのお知らせか否かを返す
func (ns *noticeSeverity) IsNoticeSeverityInfo() bool {
	if ns == nil {
		return false
	}
	return ns.val == NoticeSeverityInfo
}

// IsNoticeSeverityImportant ... 重要レベルのお知らせか否かを返す
func (ns *noticeSeverity) IsNoticeSeverityImportant() bool {
	if ns == nil {
		return false
	}
	return ns.val == NoticeSeverityImportant
}

// IsNoticeSeverityInfo ... 通常レベルのお知らせか否かを返す
func IsNoticeSeverityInfo(val int) bool {
	return val == NoticeSeverityInfo
}

// IsNoticeSeverityImportant ... 重要レベルのお知らせか否かを返す
func IsNoticeSeverityImportant(val int) bool {
	return val == NoticeSeverityImportant
}
