#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

// N.B. Don't use plain " as escape \^(1, 3, 7, 15 ...) ... blah, blah ...
#define apos "\'"
// inside back usage >> from C stdio proxy
#define quot "\\\\\\\""
// inside back usage >> from C stdio proxy

char* names[] = { // command descriptions
	// doesn't really need the escape here, but ...
	"Compile Mode " apos "a" apos " Using " apos "a" apos,//0
	"Help"//1
};

bool wrapio[] = { // decides if stdio is wrapped by a self proxy call
	false,//0
	false//1
};

int back_to(char sys[], char* argv) { // allow argv passing
	// must be static to prevent stack smashing
	// allocate about a page
	static char bash[4000] = "bash -c \"coproc (sleep 1 && rofi -e \\\"$(";
	strcat(bash, sys);
	if(argv) {
		strcat(bash, argv); // name
		strcat(bash, "'"); 
	}
	strcat(bash, ")\\\")\"");
	return system(bash);
}

int back(char sys[]) { // no argv call bash
	return back_to(sys, NULL);
}

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

char* icon_names[] = { // indexed icon names used in icons[]
	"terminal",//0
	"help"//1
};

int icons[] = { // icon numbers for commands
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
		case 1:// process 
			if(strcmp(argv[1], names[i])) break;
			// it's how th IO gets sent to the right place instead of messing up rofi
			// allows easy C stdio
			// don't emit ' or "
			if(wrapio[i]) return back_to("cd ~/.config/rofi/scripts && ./a.out -- '", names[i]);// proxy call self
			return fn[i](argc - 1, argv + 1);// do it
			//ok done
		default:// process proxy (2 args)
			if(strcmp(argv[2], names[i])) break;
			return fn[i](argc - 2, argv + 2);// do it
                        //ok done
		}
	}
	return argc - 1; //ok or not
}