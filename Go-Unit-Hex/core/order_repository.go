package core

type OrderRepository interface { //secondary port
	Save(order Order) error
}
