package valueobject

import (
	"strings"

	"github.com/google/uuid"
)

// NewUniqueID ...
func NewUniqueID() UniqueID {
	return &uniqueID{val: CreateUniqueID()}
}

// NewUniqueIDByParam ...
func NewUniqueIDByParam(val string) UniqueID {
	return &uniqueID{val: val}
}

// UniqueID ... 各エンティティをユニークに識別するIDとして利用
type UniqueID interface {
	GetVal() string
	Equals(comparison UniqueID) bool
}

type uniqueID struct {
	val string
}

// GetVal ...
func (id *uniqueID) GetVal() string {
	return id.val
}

// Equals ...
func (id *uniqueID) Equals(comparison UniqueID) bool {
	if id == nil || comparison == nil {
		return false
	}
	return id.GetVal() == comparison.GetVal()
}

// CreateUniqueID ...
func CreateUniqueID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
