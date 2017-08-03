#include <string>
#include <ros.h>
#include <embedded_linux_hardware.h>
#include <std_msgs/String.h>
#include "_cgo_export.h"



void messageCb(const std_msgs::String& received_msg){
	//send data to go side
	std::string data = received_msg.data;
	subArray = &data[0];
	stringFromC();
		
}


class StringSubscriberWrapper {
    private:

        ros::NodeHandle nh;
        ros::Subscriber<std_msgs::String> sub;
    public:
        StringSubscriberWrapper(char *ip, char *topic): sub(topic, messageCb) {
            nh.initNode(ip);
            nh.subscribe(sub);
        }
        void subscribeString() {
            nh.spinOnce();
        }
};

std::string test = "null";

char *subArray = &test[0];

class StringPublisherWrapper {
    private:

        ros::NodeHandle nh2;
        std_msgs::String str_msg;
        ros::Publisher publisher;
    public:
        StringPublisherWrapper(char *ip, char *topic): publisher(topic, &str_msg) {
            nh2.initNode(ip);
            nh2.advertise(publisher);
        }
        void publishString(char* data, int len) {
            str_msg.data = (char*)data;
            publisher.publish(&str_msg);
            nh2.spinOnce();
        }
};

/* Wrapper functions */

extern "C" StringSubscriber* NewStringSubscriber(char *ip, char *topic) {
    StringSubscriber* str_sub = new StringSubscriber();
    str_sub->wrapper = new StringSubscriberWrapper(ip, topic);
    return str_sub;
}

extern "C" void SubscribeString(StringSubscriber* str_sub) {
    ((StringSubscriberWrapper *)str_sub->wrapper)->subscribeString();
}

extern "C" void DeleteStringSubscriber(StringSubscriber* str_sub) {
    StringSubscriberWrapper* wrapper = (StringSubscriberWrapper *)(str_sub->wrapper);
    delete wrapper;
    delete str_sub;
}

extern "C" StringPublisher* NewStringPublisher(char *ip, char *topic) {
    StringPublisher* str_pub = new StringPublisher();
    str_pub->wrapper2 = new StringPublisherWrapper(ip, topic);
    return str_pub;
}

extern "C" void PublishString(StringPublisher* str_pub, char *data, int len) {
    ((StringPublisherWrapper *)str_pub->wrapper2)->publishString(data, len);
}
extern "C" void DeleteStringPublisher(StringPublisher* str_pub) {
    StringPublisherWrapper* wrapper2 = (StringPublisherWrapper *)(str_pub->wrapper2);
    delete wrapper2;
    delete str_pub;
}