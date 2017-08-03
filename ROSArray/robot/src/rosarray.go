package ROSArray

// #include "rosarray.h"
import "C"

import (
	"mind/core/framework/log"
	"mind/core/framework/skill"
	"time"
)

//export floatFromC
func floatFromC(val float64, row int, col int) {
	//hi := C.GoString(C.subArray)
	poseArr[row][col] = val
	
}

//export printPose
func printPose() {
	log.Info.Println(poseArr)
}

const (
	rosMasterIP = "192.168.0.109" // ROS_MASTER_IP need to be modified manually
	rosSubTopic  = "floater"
	hello = "HelloWorld"
)

var poseArr [19][2]float64

type ROSArray struct {
	skill.Base
	FloatSubscriber *C.FloatSubscriber
	stop           chan bool
}

func NewSkill() skill.Interface {
	return &ROSArray{
		FloatSubscriber: C.NewFloatSubscriber(C.CString(rosMasterIP), C.CString(rosSubTopic)),
		stop:           make(chan bool),
	}
}

func (d *ROSArray) subPubFloats() {
	log.Info.Println("starting subscriber")
	for {
		select {
		case <-d.stop:
			break
		default:
		}
		C.SubscribeFloat(d.FloatSubscriber)

		//log.Info.Println("spun")
		time.Sleep(1000 * time.Millisecond)
	}
}

func (d *ROSArray) OnStart() {
	d.subPubFloats()
}

func (d *ROSArray) OnClose() {
	d.stop <- true
	C.DeleteFloatSubscriber(d.FloatSubscriber)
}

func (d *ROSArray) OnConnect() {
	// Use this method to do something when the remote connected.
}

func (d *ROSArray) OnDisconnect() {
	// Use this method to do something when the remote disconnected.
}

func (d *ROSArray) OnRecvJSON(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
}

func (d *ROSArray) OnRecvFloat(data string) {
	// Use this method to do something when skill receive string from remote client.
}
