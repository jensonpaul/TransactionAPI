// This function iterates through each character in the PAN string and appends an asterisk for every character except the last four. The final masked PAN string is returned.
func MaskPAN(pan string) string {
    masked := ""
    for i, ch := range pan {
        if i < len(pan)-4 {
            masked += "*"
        } else {
            masked += string(ch)
        }
    }
    return masked
}
// This function takes a list of transactions as input, iterates over each transaction, and creates a modified version of each transaction where the PAN is replaced with asterisks, except for the last four digits. The modified transactions are added to a new list, which is returned at the end.
def anonymize_transactions(transactions):
    anonymized_transactions = []
    for transaction in transactions:
        anonymized_pan = '*' * 12 + transaction['pan'][-4:]
        anonymized_transaction = {
            'id': transaction['id'],
            'amount': transaction['amount'],
            'currency': transaction['currency'],
            'date': transaction['date'],
            'pan': anonymized_pan,
            'merchant': transaction['merchant']
        }
        anonymized_transactions.append(anonymized_transaction)
    return anonymized_transactions
// The maskedPAN function takes a PAN string and returns a masked version of it, with all characters except the last four replaced with asterisks.
func maskedPAN(pan string) string {
    masked := strings.Repeat("*", len(pan)-4) + pan[len(pan)-4:]
    return masked
}
// Inside the function, we create a new slice with the same length as the input transactions slice, and then loop through each transaction in the input slice. For each transaction, we extract the last four digits of the PAN using the len() and [] operators, and then create a new string with asterisks for the rest of the PAN digits.
func maskPAN(transactions []Transaction) []Transaction {
    maskedTransactions := make([]Transaction, len(transactions))

    for i, transaction := range transactions {
        maskedPAN := "**** **** **** " + transaction.PAN[len(transaction.PAN)-4:]
        maskedTransactions[i] = Transaction{
            PAN:      maskedPAN,
            Amount:   transaction.Amount,
            Currency: transaction.Currency,
        }
    }

    return maskedTransactions
}

func maskPan(pan string) string {
    maskedPan := "**** **** **** " + pan[len(pan)-4:]
    return maskedPan
}

// maskPAN replaces all but the last four digits of a PAN with asterisks
func maskPAN(pan string) string {
	if len(pan) <= 4 {
		return pan
	}
	masked := make([]byte, len(pan)-4)
	for i := range masked {
		masked[i] = '*'
	}
	return string(masked) + pan[len(pan)-4:]
}

func maskPAN() {
    // create a new slice to hold the masked transactions
    maskedTransactions := make([]Transaction, len(transactions))
    // loop through each transaction and mask the PAN number
    for i, transaction := range transactions {
        maskedTransaction := transaction
        panStr := fmt.Sprintf("%d", transaction.PAN)
        if len(panStr) > 4 {
            maskedTransaction.PAN = 0 // set to zero for security
            maskedTransaction.PAN = transaction.PAN % 10000 // get last four digits
            maskedTransaction.PAN += 10000 // add a prefix to ensure the number has four digits
            maskedTransaction.PAN *= 1000000 // shift the number to the left
            // maskedTransaction.PAN *= 10000 // shift the number to the left
            maskedTransactions[i] = maskedTransaction
        } else {
            maskedTransactions[i] = transaction
        }
    }
}

func maskPAN() {
    // create a new slice to hold the masked transactions
    maskedTransactions := make([]Transaction, len(transactions))
    
    // loop through each transaction and mask the PAN number
    for i, transaction := range transactions {
        maskedTransaction := transaction
        panStr := fmt.Sprintf("%d", transaction.PAN)
        if len(panStr) > 4 {
            lastFourDigits := transaction.PAN % 10000 // get last four digits
            maskedTransaction.PAN = lastFourDigits
            for j := 0; j < 12; j++ {
                maskedTransaction.PAN *= 10 // shift the number to the left
                maskedTransaction.PAN += 9 // add a digit to the number
            }
        }
        maskedTransactions[i] = maskedTransaction
    }
}

func maskPAN() {
    // create a new slice to hold the masked transactions
    maskedTransactions := make([]Transaction, len(transactions))
    
    // loop through each transaction and mask the PAN number
    for i, transaction := range transactions {
        maskedTransaction := transaction
        panStr := fmt.Sprintf("%d", transaction.PAN)
        if len(panStr) > 4 {
            lastFourDigits := panStr[len(panStr)-4:]
            maskedTransaction.PAN = "************" + lastFourDigits
        }
        maskedTransactions[i] = maskedTransaction
    }
}

func maskPAN() {
    // create a new slice to hold the masked transactions
    maskedTransactions := make([]Transaction, len(transactions))

    // loop through each transaction and mask the PAN number
    for i, transaction := range transactions {
        maskedTransaction := transaction
        if maskedTransaction.PAN > 999 {
            maskedTransaction.PAN = 1000000000000000 + maskedTransaction.PAN%10000
        }
        maskedTransactions[i] = maskedTransaction
    }
}

func maskPAN() {
    // create a new slice to hold the masked transactions
    maskedTransactions := make([]Transaction, len(transactions))

    // loop through each transaction and mask the PAN number
    for i, transaction := range transactions {
        maskedTransaction := transaction
        if maskedTransaction.PAN > 999 {
            panStr := fmt.Sprintf("%d", maskedTransaction.PAN)
            lastFourDigits := panStr[len(panStr)-4:]
            maskedTransaction.PAN = 0 // reset the PAN number to zero
            maskedTransaction.PAN, err = strconv.Atoi("000000000000" + lastFourDigits)
            if err != nil {
                log.Fatal(err)
            }
        }
        maskedTransactions[i] = maskedTransaction
    }
}