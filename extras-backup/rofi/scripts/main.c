#include <stdio.h>
#include <stdlib.h>
#include <string.h>
// FILE* popen("...","w");//?
char* names[] = {
	"Compile a.out via a Bootstrap",//0
	"Help"//1
};

void back(char sys[]) {
	// must be static to prevent stack smashing
	static char bash[256] = "bash -c \"";
	strcat(bash, sys);
	strcat(bash, "\" >/dev/null");
	system(bash);
}

void compile(int argc, char** argv) {
	back("pushd ~/.config/rofi/scripts && gcc main.c && popd");// remake a.out
}

void help(int argc, char** argv) { // help function
	back("rofi -e 'Ah, and a Hello Galaxy from the help function.'");
}

void (*(fn[]))(int, char**) = {
	compile,//0
	help//1
};

char* icon_names[] = {
	"terminal",//0
	"help"//1
};

int icons[] = {
	0,//0
	1//1
};

int main(int argc, char** argv) {
	for(int i = 0; i < (sizeof(names)/sizeof(char*)); i++) {
		switch(argc - 1) {
		case 0:// list
			printf("%s", names[i]);
			putc(0, stdout);
			printf("icon\x1f%s\n", icon_names[icons[i]]);
			break;
		default:// process
			if(strcmp(argv[1], names[i])) break;
			fn[i](argc - 1, argv + 1);// do it
			return 0; //ok done
		}
	}
	return argc - 1; //ok or not
}