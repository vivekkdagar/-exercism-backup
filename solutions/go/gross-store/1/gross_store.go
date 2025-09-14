package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	measurement := map[string]int{}
    
    measurement["quarter_of_a_dozen"] = 3
    measurement["half_of_a_dozen"] = 6
    measurement["dozen"] = 12
    measurement["small_gross"] = 120
    measurement["gross"] = 144
    measurement["great_gross"] = 1728

    return measurement
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if val, exists := units[unit]; exists == false {
        return false
    } else {
		if _, exists := bill[item]; exists == false {
            bill[item] = val
        } else {
            bill[item] += val
        }
        return true
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, exists := bill[item]
    
    if exists == false {
        return exists
    }

	value, exists := units[unit]

    if exists == false || bill[item] - value < 0 {
        return false
    }
    
    bill[item] -= value
    
    if bill[item] == 0 {
        delete(bill, item)
    }
    
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (quant int, exist bool) {
	if val, exists := bill[item]; exists == false {
        quant = 0
        exist = false
    } else {
        quant = val
        exist = true
    }
    return
}
