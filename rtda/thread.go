package rtda

//java命令提供了-Xms 和-Xmx 两个非标准选项， 用来调整堆的初始大小和最大大小。
type Thread struct {
	pc int
	stack *Stack
}

func NewThread() *Thread  {
	return &Thread{
		stack:newStack(1024),
	}
}

func (this *Thread)PC()int{
	return this.pc
}

func (this *Thread)SetPC(pc int)  {
	this.pc = pc
}

func (this *Thread)PushFrame(frame *Frame)  {
	this.stack.push(frame)
}

func (this *Thread)PopFrame( ) *Frame {
	return this.stack.pop()
}

func (this *Thread)CurrentFrame() *Frame {
	return this.stack.top()
}

func (this *Thread)NewFrame(maxLocals,maxStack uint)*Frame {
	return newFrame(this,maxLocals,maxStack)
}