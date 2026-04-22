package enums

type OrderStatus int //alias

const (
	Pending OrderStatus = iota
	Shipped
	Delivered
	Cancelled
)
