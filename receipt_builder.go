package iota

import "fmt"

// NewReceiptBuilder creates a new ReceiptBuilder.
func NewReceiptBuilder(migratedAt uint32) *ReceiptBuilder {
	return &ReceiptBuilder{
		r: &Receipt{
			MigratedAt:  migratedAt,
			Funds:       Serializables{},
			Transaction: nil,
		},
	}
}

// ReceiptBuilder is used to easily build up a Receipt.
type ReceiptBuilder struct {
	r *Receipt
}

// AddEntry adds the given MigratedFundsEntry to the receipt.
func (rb *ReceiptBuilder) AddEntry(entry *MigratedFundsEntry) *ReceiptBuilder {
	rb.r.Funds = append(rb.r.Funds, entry)
	return rb
}

// AddTreasuryTransaction adds the given TreasuryTransaction to the receipt.
// This function overrides the previously added TreasuryTransaction.
func (rb *ReceiptBuilder) AddTreasuryTransaction(tx *TreasuryTransaction) *ReceiptBuilder {
	rb.r.Transaction = tx
	return rb
}

// Build builds the Receipt.
func (rb *ReceiptBuilder) Build() (*Receipt, error) {
	if _, err := rb.r.Serialize(DeSeriModePerformValidation | DeSeriModePerformLexicalOrdering); err != nil {
		return nil, fmt.Errorf("unable to build receipt: %w", err)
	}
	return rb.r, nil
}
