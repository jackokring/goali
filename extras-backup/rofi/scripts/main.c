#include <stdio.h>
#include <stdlib.h>
#include <string.h>
char* names[] = {
	"Compile Mode 'a' Using 'a'",//0
	"Help"//1
};

int back(char sys[]) {
	// must be static to prevent stack smashing
	static char bash[256] = "bash -c \"coproc (sleep 1 && rofi -e \\\"$(";
	strcat(bash, sys);
	strcat(bash, ")\\\")\"");
	return system(bash);
}

// N.B. Don't use " as escape \\ .. blah, blah ..
int compile(int argc, char** argv) { // remake a.out
	return back("cd ~/.config/rofi/scripts && gcc main.c && echo ok");
}

int help(int argc, char** argv) { // help function
	return back("echo 'Ah, and a Hello Galaxy from the help function.'");
}

int (*(fn[]))(int, char**) = {
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
			return fn[i](argc - 1, argv + 1);// do it
			//ok done
		}
	}
	return argc - 1; //ok or not
}