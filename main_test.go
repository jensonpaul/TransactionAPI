package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTransactions(t *testing.T) {
	req, err := http.NewRequest("GET", "/transactions", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTransactions)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var transactions []Transaction
	err = json.Unmarshal(rr.Body.Bytes(), &transactions)
	if err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}

	if len(transactions) != 10 {
		t.Errorf("expected 10 transactions, got %d", len(transactions))
	}
}

func TestGetOrderedTransactions(t *testing.T) {
	req, err := http.NewRequest("GET", "/ordered", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOrderedTransactions)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var transactions []Transaction
	err = json.Unmarshal(rr.Body.Bytes(), &transactions)
	if err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}

	if len(transactions) != 10 {
		t.Errorf("expected 10 transactions, got %d", len(transactions))
	}

	if transactions[0].PostedTimeStamp < transactions[1].PostedTimeStamp {
		t.Errorf("transactions not sorted in descending order of PostedTimeStamp")
	}

	if transactions[0].PAN == 1234567890123456 {
		t.Errorf("PAN number was not masked")
	}
}

func TestMaskPAN(t *testing.T) {
	transaction := Transaction{
		ID:                  1,
		Amount:              1000,
		MessageType:         "Transaction",
		CreatedAt:           "2022-03-22T10:11:12Z",
		TransactionID:       123456789,
		PAN:                 1234567890123456,
		TransactionCategory: "Shopping",
		PostedTimeStamp:     "2022-03-22T10:11:12Z",
		TransactionType:     "Debit",
		SendingAccount:      12345678,
		ReceivingAccount:    98765432,
		TransactionNote:     "Test Transaction",
	}

	maskedTransaction := MaskPAN(transaction)

	if maskedTransaction.PAN == transaction.PAN {
		t.Errorf("PAN number was not masked")
	}
	// Verify that the original transaction's PAN number was not modified
	if transaction.PAN != 1234567890123456 {
		t.Errorf("MaskPAN function modified original transaction PAN number, expected %d but got %d", 1234567890123456, transaction.PAN)
	}

	// Verify that the masked transaction's PAN number is different from the original transaction's PAN number
	if maskedTransaction.PAN == 1234567890123456 {
		t.Errorf("MaskPAN function did not mask the PAN number, expected masked PAN number but got %d", maskedTransaction.PAN)
	}

	// Verify that the masked transaction's PAN number has the correct last 4 digits
	if maskedTransaction.PAN%10000 != 3456 {
		t.Errorf("MaskPAN function did not mask the PAN number correctly, expected masked PAN number to have last 4 digits 3456 but got %d", maskedTransaction.PAN%10000)
	}
}
