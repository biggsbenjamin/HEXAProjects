package cgocTest

// #include "cgoctest.h"
import "C"

import (
	"mind/core/framework/skill"
	"mind/core/framework/log"
)

//export intFromGo
func intFromGo(hello int) {
	log.Info.Println(hello)
}

//export stringFromGo
func stringFromGo() {
	hi := C.GoString(C.hello)
	log.Info.Println(hi)
}


type cgocTest struct {
	skill.Base
}

func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return &cgocTest{}
}

func (d *cgocTest) OnStart() {
	// Use this method to do something when this skill is starting.
	C.main2()
}

func (d *cgocTest) OnClose() {
	// Use this method to do something when this skill is closing.
}
