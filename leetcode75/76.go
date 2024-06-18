package leetcode75

type StockSpanner struct {
	monoStack [][2]int
}

func Constructor76() StockSpanner {
	return StockSpanner{
		monoStack: make([][2]int, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	ret := 1
	for len(this.monoStack) > 0 && this.monoStack[len(this.monoStack)-1][0] <= price {
		ret += this.monoStack[len(this.monoStack)-1][1]
		this.monoStack = this.monoStack[:len(this.monoStack)-1]
	}
	this.monoStack = append(this.monoStack, [2]int{price, ret})
	return ret
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */