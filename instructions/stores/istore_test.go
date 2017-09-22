package stores

import (
	"testing"
	"github.com/WYL53/jvmgo/instructions/base"
)

func TestISTORE(t *testing.T) {
	is := &ISTORE{}
	var inst  base.Instruction
	inst = is
	t.Log(inst)
}

