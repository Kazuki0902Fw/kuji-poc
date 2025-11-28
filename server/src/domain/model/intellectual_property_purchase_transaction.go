package model

import (
	"time"
)

func NewIPCategoryPurchaseTransaction(ipCategory *IntellectualPropertyCategory, purchaseQuantity int) (*IntellectualPropertyPurchaseTransaction, error) {
	now := time.Now()
	purchasePrice := ipCategory.Price * float64(purchaseQuantity)
	transaction := &IntellectualPropertyPurchaseTransaction{
		ID: NewID(),
		// ユーザー機能作ってから追加
		// UserID: userID,
		IPCategoryID: ipCategory.ID,
		PurchaseQuantity: purchaseQuantity,
		PurchasePrice: purchasePrice,
		PaymentMethod: PurchaseTransactionPaymentMethodCreditCard,
		ProviderTransactionID: nil,
		Status: PurchaseTransactionStatusPaymentPending,
		PaidAt: nil,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return transaction, nil
}

func (t *IntellectualPropertyPurchaseTransaction) UpdateStatusPaid() {
	now := time.Now()
	t.Status = PurchaseTransactionStatusPaymentSuccess
	t.PaidAt = &now
	t.UpdatedAt = now
}

func (t *IntellectualPropertyPurchaseTransaction) UpdateStatusPaymentFailed() {
	now := time.Now()
	t.Status = PurchaseTransactionStatusPaymentFailed
	t.UpdatedAt = now
}

func (t *IntellectualPropertyPurchaseTransaction) UpdateStatusDrawSuccess() {
	now := time.Now()
	t.Status = PurchaseTransactionStatusDrawSuccess
	t.UpdatedAt = now
}

func (t *IntellectualPropertyPurchaseTransaction) UpdateStatusDrawFailed() {
	now := time.Now()
	t.Status = PurchaseTransactionStatusDrawFailed
	t.UpdatedAt = now
}
