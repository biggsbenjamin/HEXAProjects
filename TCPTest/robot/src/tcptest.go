//testing the tcp capabilities of the hexa in preparation for using rosserial
package TCPTest

import (
	"mind/core/framework/skill"
	"mind/core/framework/log"
	"mind/core/framework/drivers/hexabody"

	"net"
	"os"
	"strings"
	"strconv"
)

type TCPTest struct {
	skill.Base
}
//taken from a tcp golang server example
const (
   	CONN_HOST = "localhost"
   	CONN_PORT = ":3333"
   	CONN_TYPE = "tcp"	
)

func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return &TCPTest{}
}

func (d *TCPTest) OnStart() {
	hexabody.Start()
	// Use this method to do something when this skill is starting.

	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_PORT)
	if err != nil {
        	log.Info.Println("Error listening:", err.Error())
        	os.Exit(1)
    	}
    	// Close the listener when the application closes.
    	defer l.Close()
    	log.Info.Println("Listening on " + CONN_PORT)
    	for {
        	// Listen for an incoming connection.
        	conn, err := l.Accept()
        	if err != nil {
            		log.Info.Println("Error accepting: ", err.Error())
            		os.Exit(1)
        	}
        	// Handle connections in a new goroutine.        	
		handleRequest(conn)
    	}
	//debug
	log.Info.Println("skipped")
}

func handleRequest(conn net.Conn) {
	for {
	  	// Make a buffer to hold incoming data.
	  	buf := make([]byte, 16)
	  	// Read the incoming connection into the buffer.
	  	reqLen, err := conn.Read(buf)
	  	if err != nil {
	    		log.Info.Println("Error reading:", err.Error())
			return
	  	}
		data := string(buf[:reqLen])
		log.Info.Println(data, reqLen)
	  	// Send a response back to person contacting us.
	  	conn.Write([]byte(data))
	  	// Close the connection when you're done with it.
	  	if reqLen != 0 {
			legTest(data)
		} else {
			conn.Close()		
		}
	}	
}
//legtest function from MoveLegSkill
func legTest(data string) {
	
	pos_slice := strings.Split(data, ":")

	for joint, angle := range pos_slice {
		ang, _ := strconv.ParseFloat(angle, 64)
		log.Info.Println(joint, ang)
		if checkJoint(joint, ang) {
			hexabody.MoveJoint(0, joint, ang, 1)
		}
	}
}

//error catching function that checks the ranges of the joint angles to see if they are within the hardware limits, legs can still collide with each other
func checkJoint(jointNum int, angle float64) bool {
	good_degree := false
	if jointNum == 0 {
		if angle > 35.0 && angle < 145.0 {
			good_degree = true
		}
	} else if jointNum == 1 {
		if angle > 10.0 && angle < 170.0 {
			good_degree = true
		}
	} else if jointNum == 2 {
		if angle > 10.0 && angle < 160.0 {
			good_degree = true
		}
	} else {
		good_degree = false
	}
	return good_degree
}

func (d *TCPTest) OnClose() {
	hexabody.Close()
}

func (d *TCPTest) OnDisconnect() {
	os.Exit(0) // Closes the process when remote disconnects
}



