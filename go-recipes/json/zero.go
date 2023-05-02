package main

// LineItem is a line in receipt
type LineItem struct {
	SKU      string
	Price    float64
	Discount float64
	Quantity int
}

func parseLineItem([]byte) (LineItem, error) {
	return LineItem{}, nil
}

func handleZero() interface{} {
	firstLine := []byte(`{"sku": "x3xs", "price": 1.2}`)
	firstLineItem, err := parseLineItem(firstLine)
	if err != nil {
		return err
	}

	secondLine := []byte(`{"sku": "x3xs", "price": 1.2, "quantity": 0}`)
	secondLineItem, err := parseLineItem(secondLine)
	if err != nil {
		return err
	}

	return []LineItem{firstLineItem, secondLineItem}
}
