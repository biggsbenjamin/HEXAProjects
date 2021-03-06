#include <ros.h>
#include <embedded_linux_hardware.h>
#include <std_msgs/Float64.h>
#include <std_msgs/Float64MultiArray.h>
#include <std_msgs/MultiArrayDimension.h>
#include "_cgo_export.h"

//subscriber callback function
void messageCb(const std_msgs::Float64MultiArray& msg){
	//send data to go side
	for (int i=0;i<19;i++){
		for(int j=0;j<2;j++){
			CGoCallback(msg.data[
			(msg.layout.data_offset) + 
			(i*msg.layout.dim[0].stride) + j],
			i, j );
		}
	}
	//see go code for function
	runPose();
}

//wrapper class for the subscriber
class FloatSubscriberWrapper {
	private:
		//setting up the node
        	ros::NodeHandle nh;
		//setting up the types of messages the subscriber receives
        	ros::Subscriber<std_msgs::Float64MultiArray> sub;
	public:
		//subscriber class constructor called by Go code
        	FloatSubscriberWrapper(char *ip, char *topic): sub(topic, messageCb) {
            		nh.initNode(ip);
            		nh.subscribe(sub);
        	}
		//performs the spin function on the subscriber
        	void spinSub() {
            		nh.spinOnce();
        	}
};

/* Wrapper functions */
//these functions can be used by the go code and map to the subscriber class
//maps to the constructor
extern "C" FloatSubscriber* NewFloatSubscriber(char *ip, char *topic) {
    FloatSubscriber* flo_sub = new FloatSubscriber();
    flo_sub->wrapper = new FloatSubscriberWrapper(ip, topic);
    return flo_sub;
}
//maps to the spin function
extern "C" void SpinSub(FloatSubscriber* flo_sub) {
    ((FloatSubscriberWrapper *)flo_sub->wrapper)->spinSub();
}
//deletes the subscriber object
extern "C" void DeleteFloatSubscriber(FloatSubscriber* flo_sub) {
    FloatSubscriberWrapper* wrapper = (FloatSubscriberWrapper *)(flo_sub->wrapper);
    delete wrapper;
    delete flo_sub;
}
