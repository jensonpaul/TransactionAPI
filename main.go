package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/gorilla/mux"
)

type Transaction struct {
	ID                  int    `json:"id"`
	Amount              int    `json:"amount"`
	MessageType         string `json:"conversation_type"`
	CreatedAt           string `json:"created_at"`
	TransactionID       int    `json:"created_at"`
	PAN                 int    `json:"pan"`
	TransactionCategory string `json:"transaction_category"`
	PostedTimeStamp     string `json:"posted_timestamp"`
	TransactionType     string `json:"transaction_type"`
	SendingAccount      int    `json:"sending_account"`
	ReceivingAccount    int    `json:"receiving_account"`
	TransactionNote     string `json:"transaction_note"`
}

var transactions []Transaction

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	//mock data
	transactions, err := readTransactionsFromFile("transactions.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\nPrinting Current Transactions\n")

	// create a new slice to hold the masked transactions
	maskedTransactions := make([]Transaction, len(transactions))

	// loop through each transaction and mask the PAN number
	for i, transaction := range transactions {
		maskedTransaction := MaskPAN(transaction)
		maskedTransactions[i] = maskedTransaction
	}

	json.NewEncoder(w).Encode(maskedTransactions)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// Route to get all transactions
	router.HandleFunc("/transactions", GetTransactions).Methods("GET")
	// Route to get all transactions by PostedTimeStamp order
	router.HandleFunc("/ordered", GetOrderedTransactions).Methods("GET")
	// Route to get transaction by ID
	router.HandleFunc("/transactions/{id}", GetTransactionByID).Methods("GET")

	fmt.Printf("Serving transactions on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func readTransactionsFromFile(filename string) ([]Transaction, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var transactions []Transaction
	err = json.NewDecoder(file).Decode(&transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Get any parameters passed through the URL
	transactions, err := readTransactionsFromFile("transactions.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\nPrinting Specific Transactions\n")
	for _, transaction := range transactions {
		if fmt.Sprintf("%d", transaction.ID) == params["id"] {
			// Mask the PAN number
			maskedTransaction := MaskPAN(transaction)
			json.NewEncoder(w).Encode(maskedTransaction)
			return
		}
	}
	json.NewEncoder(w).Encode(&Transaction{})
}

func GetOrderedTransactions(w http.ResponseWriter, r *http.Request) {
	// Read transactions from JSON file
	transactions, err := readTransactionsFromFile("transactions.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\nPrinting Transactions ordered by PostedTimeStamp\n")
	// create a new slice to hold the masked transactions
	maskedTransactions := make([]Transaction, len(transactions))

	// loop through each transaction and mask the PAN number
	for i, transaction := range transactions {
		maskedTransaction := MaskPAN(transaction)
		maskedTransactions[i] = maskedTransaction
	}

	// sort the masked transactions by PostedTimeStamp
	sort.Slice(maskedTransactions, func(i, j int) bool {
		return maskedTransactions[i].PostedTimeStamp > maskedTransactions[j].PostedTimeStamp
	})

	json.NewEncoder(w).Encode(maskedTransactions)
}

func MaskPAN(transaction Transaction) Transaction {
	maskedTransaction := transaction
	panStr := fmt.Sprintf("%d", transaction.PAN)
	if len(panStr) > 12 {
		maskedTransaction.PAN = 1000000000000000 + transaction.PAN%10000
	}
	return maskedTransaction
}
