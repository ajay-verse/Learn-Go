package orders

import (
	// Go Internal Packages
	"time"

	// Local Packages
	"learn-go/errors"
)

type Order struct {
	OrderID     string      `json:"order_id"`
	UserID      string      `json:"user_id"`
	LineItems   []LineItems `json:"line_items"`
	OrderStatus string      `json:"order_status"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	ShippedAt   string      `json:"shipped_at"`
	DeliveredAt string      `json:"delivered_at"`
}

type LineItems struct {
	ItemID   string `json:"item_id"`
	Quantity int    `json:"quantity"`
	Price    string `json:"price"`
}

func (o *Order) ValidateCreation() error {
	ve := errors.ValidationErrs()

	if o.OrderID != "" {
		ve.Add("order_id", "need to be empty")
	}
	if o.UserID == "" {
		ve.Add("user_id", "cannot be empty")
	}
	if o.LineItems == nil {
		ve.Add("line_items", "cannot be empty")
	}
	if o.OrderStatus == "" {
		ve.Add("order_status", "cannot be empty")
	}
	if o.ShippedAt == "" {
		o.ShippedAt = "Will Be Shipped Soon"
	}
	if o.DeliveredAt == "" {
		o.DeliveredAt = "Will Be Delivered Soon"
	}

	return ve.Err()
}

func (o *Order) ValidateUpdate(orderID string) error {
	ve := errors.ValidationErrs()

	if o.OrderID != orderID {
		ve.Add("order_id", "not matched")
	}
	if o.UserID == "" {
		ve.Add("user_id", "cannot be empty")
	}
	if o.LineItems == nil {
		ve.Add("line_items", "cannot be empty")
	}
	if o.OrderStatus == "" {
		ve.Add("order_status", "cannot be empty")
	}
	if o.CreatedAt.IsZero() {
		ve.Add("created_at", "cannot be empty")
	}
	if o.ShippedAt == "" {
		o.ShippedAt = "Will Be Shipped Soon"
	}
	if o.DeliveredAt == "" {
		o.DeliveredAt = "Will Be Delivered Soon"
	}

	return ve.Err()
}
