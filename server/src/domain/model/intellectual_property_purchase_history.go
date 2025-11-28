package model

import "time"

func NewIntellectualPropertyPurchaseHistory(purchaseTransactionID ID, intellectualPropertyID ID) *IntellectualPropertyPurchaseHistory {
	now := time.Now()
	return &IntellectualPropertyPurchaseHistory{
		ID: NewID(),
		PurchaseTransactionID: purchaseTransactionID,
		IntellectualPropertyID: intellectualPropertyID,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
