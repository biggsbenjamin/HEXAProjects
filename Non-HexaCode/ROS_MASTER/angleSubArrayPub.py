#!/usr/bin/env python

import rospy
from std_msgs.msg import Float64, Float64MultiArray, MultiArrayLayout, MultiArrayDimension

#setting up a data array using the standard ros messages
poseArr =Float64MultiArray()
#adding dimensions to the array
poseArr.layout.dim.append(MultiArrayDimension())
poseArr.layout.dim.append(MultiArrayDimension())
#labelling the dimensions
poseArr.layout.dim[0].label = "angle"
poseArr.layout.dim[1].label = "millis"
#size of the dimensions
poseArr.layout.dim[0].size = 19
poseArr.layout.dim[1].size = 19
#not required by this code but throws error if removed
poseArr.layout.dim[0].stride = 2
poseArr.layout.dim[1].stride = 19
poseArr.layout.data_offset = 0
#sets all values to 90 degrees
poseArr.data = [90.0]*38

#posArray needs to be a list of angles of each joint followed time to movement in millis
def arrayPub(posArray):
	#prepare the publisher with topic name, message type 
	pub = rospy.Publisher('floater', Float64MultiArray, queue_size=10)
	rospy.init_node('talker', anonymous=True)
	rate = rospy.Rate(1) #hz
	#each row contains the values for one array which contains all joint positions and movement times for one pose
	for row in posArray:
		i = 0
		#error handling
		if not rospy.is_shutdown():
			#iterate through each cell of the csv and assign it to the data of the pose array
			for val in row:
				if val != "":
					poseArr.data[i] = float(val)
				i+=1
			pub.publish(poseArr)
			print poseArr.data
			#sleeps so that loop occurs at the rospy rate 
			rate.sleep()
		else:
			break
def callBack(data):
#read adc
	convList = []
	i=0
	for leg in range(6):
	#convert adc to angles with maths
	#current conversion based on leg 4, need to repeat for all legs
		convList.append((data.data[i*3]-904.5909)/(-4.6455))
		convList.append((data.data[i*3 + 1]-84.0625)/(4.4938))
		convList.append((data.data[i*3 + 2]-133.7333)/(4.6267))
	print(convList)

def angleSub():
	rospy.init_node('adc_read', anonymous=True)
	rospy.Subscriber("adc_feed", Float64MultiArray, callBack)
	rospy.spin()

if __name__ == '__main__':

	try:
		angleSub()
		#arrayPub()
#error handling
	except rospy.ROSInterruptException:
	        pass
