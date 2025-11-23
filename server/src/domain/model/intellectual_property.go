package model

import (
	"time"
)

// ストックを減らすメソッドを定義
func (p *IntellectualProperty) decreaseStock() {
	now := time.Now()
	p.Stock--
	p.UpdatedAt = now
}
