package entity

type OrderQueue struct {
	Orders []*Order
}

// Len implements heap.Interface.
func (*OrderQueue) Len() int {
	panic("unimplemented")
}

// Less implements heap.Interface.
func (*OrderQueue) Less(i int, j int) bool {
	panic("unimplemented")
}

func (oq OrderQueue) less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}
func (oq OrderQueue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

func (oq OrderQueue) len() int {
	return len(oq.Orders)
}

func (oq *OrderQueue) Push(x interface{}) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

func (oq OrderQueue) Pop() interface{} {
	old := oq.Orders
	n := len(old)
	item := old[n-1]
	oq.Orders = old[0 : n-1]
	return item
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}
