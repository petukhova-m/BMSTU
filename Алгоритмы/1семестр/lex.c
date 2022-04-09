#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(){
	char str[1000];
	scanf("%s", str);
	int i = 0;
	while(str[i] != "\0"){
		if (str[i] == "+")
			printf("SPEC 0\n");
		if (str[i] == "-")
			printf("SREC 1\n");
		if (str[i] == "*")
			printf("SPEC 2\n");
		if (str[i] == "/")
			printf("SREC 3\n");
		if (str[i] == "(")
			printf("SPEC 4\n");
		if (str[i] == ")")
			printf("SPEC 5\n");
		if (CONST(str[i])
			printf
	return 0;
}
