#include <stdio.h>
#include <string>
#include "_cgo_export.h" //needs to be included for c++ to access exported go code
//testing how data can be accessed by the instance of Go if it is initiallised in c++
std::string hi = "helloworld";

//reading data into hello variable
char *hello = &hi[0];

int main2() {
    //printf("Hi from C++\n");
	int x = 10;
	//using exported go code
	intFromGo(x);
	// ""
	stringFromGo();
	//reading new data into the shared c++/go variable
	std::string hi2 = "helloworld2";
	hello = &hi2[0];
	//accessing with go code
	stringFromGo();
    
    return 0;
}
