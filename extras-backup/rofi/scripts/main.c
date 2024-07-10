#include <stdio.h>
#include <stdlib.h>
#include <string.h>
// FILE* popen("...","w");//?
char* names[] = {
	"Compile Rofi Boot",//0
	"test"//1
};

void compile(int argc, char** argv) {
	system("bash -c \"pushd ~/.config/rofi/scripts && gcc main.c && popd\"");// remake a.out
}

void nop(int argc, char** argv) { // null function
	printf("Ah, and a Hello Galaxy from the test function.");
}

void (*(fn[]))(int, char**) = {
	compile,//0
	nop//1
};

int main(int argc, char** argv) {
	--argc;
	for(int i = 0; i < (sizeof(names)/sizeof(char*)); i++) {
		switch(argc) {
		case 0:// list
			printf("%s\n", names[i]);
			break;
		default:// process
			if(strcmp(argv[1], names[i])) break;
			fn[i](--argc, ++argv);// do it
			return 0; //ok done
		}
	}
	return argc; //ok or not
}