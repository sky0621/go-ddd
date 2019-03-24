package aggregate

// NewNoticeReadCondition ...
func NewNoticeReadCondition(notices []Notice) NoticeReadCondition {
	return &noticeReadCondition{notices: notices}
}

// NoticeReadCondition ... 条件に該当する「お知らせ」データを決定するために利用
type NoticeReadCondition interface {

	// FIXME 複数の Notice を持つなり、ページング情報を持つなり、要件にあった取得条件を定義！
}

type noticeReadCondition struct {
	notices []Notice
}
