package rtda

type Frame struct {
	lower *Frame //用来 实现 链 表 数据 结构
	localVars LocalVars //保存局部变量表指针
	operandStack *OperandStack //保存 操 作数 栈 指针。

}

//执行 方法 所需 的 局部 变量 表 大小 和 操 作数 栈 深度 是由 编译器 预先 计算 好的，
// 存储 在 class 文件 method_ info 结构 的 Code 属性 中
func NewFrame(maxLocals,maxStack uint)*Frame  {
	return &Frame{
		localVars:newLocalVars(maxLocals),
		operandStack:newOperandStack(maxStack),
	}
}

func (this *Frame)LocalVars() LocalVars {
	return  this.localVars
}

func (this *Frame)OperandStack() *OperandStack  {
	return  this.operandStack
}