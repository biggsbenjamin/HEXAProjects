package ROSNode

// #include "rosnode.h"
import "C"

import (
	"mind/core/framework/log"
	"mind/core/framework/skill"
	"time"
)

//export stringFromC
func stringFromC() {
	hi := C.GoString(C.subArray)
	log.Info.Println(hi)
}

const (
	rosMasterIP = "192.168.0.234" // ROS_MASTER_IP need to be modified manually
	rosSubTopic  = "subTopic"
	rosPubTopic  = "pubTopic"
	hello = "HelloWorld"
)

type ROSNode struct {
	skill.Base
	StringNode *C.StringNode
	stop chan bool
}

func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return &ROSNode{StringNode: C.NewStringNode(C.CString(rosMasterIP), C.CString(rosSubTopic), C.CString(rosPubTopic)),
		stop: make(chan bool),}
}

func (d *ROSNode) OnStart() {
	// Use this method to do something when this skill is starting.
	d.pubStrings()
}

func (d *ROSNode) OnClose() {
	// Use this method to do something when this skill is closing.
	d.stop <- true
	C.DeleteStringNode(d.StringNode)
}

func (d *ROSNode) pubStrings() {
	//go d.subStrings()
	for {
		select {
		case <-d.stop:
			break
		default:
		}
		C.PublishString(
			d.StringNode,
			C.CString(hello),
			C.int(len(hello)),
		)
		log.Info.Println("Sent String with length:", len(hello))
		time.Sleep(1000 * time.Millisecond)
	}
}

func (d *ROSNode) subStrings(){
	for {
		C.SubscribeString(d.StringNode)

		//log.Info.Println("spun")
		time.Sleep(1000 * time.Millisecond)
	}
}
