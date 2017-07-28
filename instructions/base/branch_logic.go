package base

import "github.com/WYL53/jvmgo/rtda"

func Branck(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
