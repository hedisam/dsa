package platestack

type DinnerPlates struct {
	capacity int 
    stacks []*Stack
	pq *MinHeap
	popPosition int
}


func Constructor(capacity int) DinnerPlates {
    return DinnerPlates{
		capacity: capacity,
		pq: NewMinHeap(10),
		popPosition: -1,
	}
}


func (this *DinnerPlates) Push(val int)  {
    pos, ok := this.pq.Min()
	if !ok {
		this.stacks = append(this.stacks, NewStack(this.capacity))
		pos = len(this.stacks) - 1
		this.popPosition = pos
		this.pq.InsertKey(pos)
	}
	
	stack := this.stacks[pos]
	stack.Push(val)
	if stack.IsFull() {
		this.pq.ExtractMin()
	}
}

func (this *DinnerPlates) Pop() int {
POP:
    if this.popPosition < 0 {
		return -1
	}
	stack := this.stacks[this.popPosition]
	if stack.IsFull() {
		this.pq.InsertKey(this.popPosition)
	}
	if stack.Size() == 1 {
		this.popPosition--
	} else if stack.Size() == 0 {
		this.popPosition--
		goto POP
	}
	return stack.Pop()
}


func (this *DinnerPlates) PopAtStack(index int) int {
	if index >= len(this.stacks) {return -1}

	if index == this.popPosition {
		// this is just like a Pop operation
		return this.Pop()
	}

    stack := this.stacks[index]
	if stack.IsFull() {
		this.pq.InsertKey(index)
	}
	if stack.IsEmpty() {return -1}
	return stack.Pop()
}