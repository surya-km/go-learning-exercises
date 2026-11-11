package methods

// Bill is a monthly bill that is due on a given day of the month (1-31).
type Bill struct {
	Name        string
	AmountCents int64
	DueDay      int
}

// Bills is a list of bills. Because it is a named type, it can have methods.
type Bills []Bill
