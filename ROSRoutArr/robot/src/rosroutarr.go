package ROSRoutArr

// #include "rosroutarr.h"
import "C"

import (
	"mind/core/framework/log"
	"mind/core/framework/skill"
	"mind/core/framework/drivers/hexabody"	
)

//export CGoCallback
func CGoCallback(val float64, row int, col int) {
	arrayUpdate <- false
	poseArr[row][col] = val	
}

//export runPose
func runPose() {
	log.Info.Println(poseArr)
	log.Info.Println("pose received")
	arrayUpdate <- true

}

const (
	rosMasterIP = "192.168.0.234" // ROS_MASTER_IP need to be modified manually
	rosSubTopic  = "floater"
)

var poseArr [19][2]float64
var arrayUpdate = make(chan bool)

type ROSRoutArr struct {
	skill.Base
	FloatSubscriber *C.FloatSubscriber
	stop chan bool
}

func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return &ROSRoutArr{
		FloatSubscriber: C.NewFloatSubscriber(C.CString(rosMasterIP), C.CString(rosSubTopic)),
		stop: make(chan bool),
	}
}

//setting up the 19 go routines, setting up the done channel, use global nature of poseArr

func (d *ROSRoutArr) setUpRout() {
	for leg := 0; leg < 6; leg++ {
		for jnt := 0; jnt < 3; jnt++ {	
			go jointPosRout(leg,jnt)	
		}
	}
	go headPosRout()
	return
}

func (d *ROSRoutArr) subPubFloats() {
	log.Info.Println("starting subscriber")
	for {
		select {
		case <-d.stop:
			break
		default:
		}
		//C.SubscribeFloat(d.FloatSubscriber)
		C.SpinSub(d.FloatSubscriber)
	}
}

func jointPosRout(legNum int, jointNum int) {
	for {
		 
		select {
		case <- arrayUpdate:
			go hexabody.MoveJoint(legNum, jointNum, 
			poseArr[(legNum*3)+jointNum][0], int(poseArr[(legNum*3)+jointNum][1]))
		}
	}
}

func headPosRout() {
	for {
		 
		select {
		case <- arrayUpdate:
			go hexabody.MoveHead(poseArr[18][0], int(poseArr[18][1]))
		}
	}
}

func (d *ROSRoutArr) OnStart() {
	// Use this method to do something when this skill is starting.
	hexabody.Start()
	d.setUpRout()
	d.subPubFloats()
}

func (d *ROSRoutArr) OnClose() {
	// Use this method to do something when this skill is closing.
	hexabody.Close()
	d.stop <- true
	C.DeleteFloatSubscriber(d.FloatSubscriber)
}
