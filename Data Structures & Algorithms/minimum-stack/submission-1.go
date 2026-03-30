// type MinStack struct {
// 	stack []int
// 	mins []int
// }

// func Constructor() MinStack {
// 	return MinStack{
// 		stack: []int{},
// 		mins: []int{},
// 	}
// }

// func (this *MinStack) Push(val int) {
// 	this.stack = append(this.stack, val)
// 	if len(this.mins) == 0 {
// 		this.mins = []int{val}
// 	} else {
// 		topMinStack := this.mins[len(this.mins)-1]
// 		this.mins = append(this.mins, min(topMinStack, val))
// 	}
// }

// func (this *MinStack) Pop() {
// 	if len(this.stack) == 0 {
// 		return
// 	}
// 	this.stack = this.stack[:len(this.stack)-1]
// 	this.mins = this.mins[:len(this.mins)-1]
// }

// func (this *MinStack) Top() int {
// 	if len(this.stack) == 0 {
// 		return 0
// 	}
// 	return this.stack[len(this.stack) - 1]
// }

// func (this *MinStack) GetMin() int {
// 	if len(this.mins) == 0 {
// 		return 0
// 	}
// 	return this.mins[len(this.mins)-1]
// }


type MinStack struct {
	stack []int
	min int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min: 0,
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, 0)
		this.min = val
	} else {
		this.stack = append(this.stack, val - this.min)
		this.min = min(this.min, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	pop := this.stack[len(this.stack)-1]
	this.min = this.Top() - pop
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	top := this.stack[len(this.stack)-1]
	if top > 0 {
		return this.min + top
	}

	return this.min
}

func (this *MinStack) GetMin() int {
	return this.min
}

