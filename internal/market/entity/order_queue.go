package entity

type OrderQueue struct {
	Orders []*Order
}

// less
func (oq *OrderQueue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

// swap
func (oq *OrderQueue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

// len
func (oq *OrderQueue) Len() int {
	return len(oq.Orders)
}

// push
func (oq *OrderQueue) Push(x interface{}) {
	order := x.(*Order)
	oq.Orders = append(oq.Orders, order)
}

// pop
func (oq *OrderQueue) Pop() interface{} {
	old := oq.Orders
	n := len(old)
	item := old[n-1]
	oq.Orders = old[:n-1]
	return item
}

// neworderqueue
func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}
