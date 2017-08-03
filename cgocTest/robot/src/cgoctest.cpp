#include <stdio.h>
#include <string>
#include "_cgo_export.h"

std::string hi = "helloworld";

char *hello = &hi[0];

int main2() {
    //printf("Hi from C++\n");
	int x = 10;
	intFromGo(x);

	stringFromGo();

	std::string hi2 = "helloworld2";
	hello = &hi2[0];
	stringFromGo();
    
    return 0;
}
