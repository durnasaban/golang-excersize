package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
    if !exists {
        return false
    }

    billValue, billExists := bill[item]
    if !billExists {
        bill[item] = value
    } else {
        bill[item] = billValue + value
    }

    return true    
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billValue, billExists := bill[item]
    if !billExists {
        return false
    }

	unitValue, unitExists := units[unit]
    if !unitExists {
        return false
    }    

    if billValue < unitValue {
        return false
    } else if billValue == unitValue {
        delete(bill, item)
        return true
    }
    
    bill[item] = billValue - unitValue
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	billValue, billExists := bill[item]
    if !billExists {
        return 0, false
    }

    return billValue, true
}
