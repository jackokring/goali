#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

// Maybe XDG factor, but depends on system
#define config "cd ~/.config/rofi/scripts&&"

// N.B. Don't use plain " as escape \^(1, 3, 7, 15 ...) ... blah, blah ...
// inside names usage maybe for proxy '...'
#define quot "\\\\\\\""
// inside back or back_to usage as literal

// luckily, any proxy C stdout is not literal
// and so is inserted in the "$(...)"
// without needing replacement in context

// useful for feedback in bash
#define okfail "&&echo ok||echo failed"

char* names[] = { // command names
	// doesn't really need the ' escape here, but ...
	// if it were a proxy stdout C routine it would be passed
	// it's really just the printing which removes the escape
	// sequences, the match works ...
	"Compile Mode 'a' Using 'a'",//0
	"Help"//1
};

bool wrapio[] = { // decides if stdio is wrapped by a self proxy call
	// rofi does not like direct stdout
	// so any C using stdout can be wrapped (= true)
	// to work as input to a rofi error message box
	// stderr is kind of silent when called by rofi
	// try "rofi -show combi -normal-window" for debug
	// of stderr stream
	false,//0
	false//1
};

bool reload[] = { // reload rofi after command by showing list again
	// this can't be used with proxy commands
	false,//0
	false//1
};

char* icon_names[] = { // indexed icon names used in icons[]
	// these do not need trailing MIME or file type
	"terminal",//0
	"help"//1
};

int icons[] = { // icon numbers for commands
	// slightly more efficient when large numbers of commands share icons
	0,//0
	1//1
};

void item(char* name, int icon, bool data) { // just in case you want to make items
	printf("%s", name);
	putc(0, stdout);
	// maybe some feedback information items
	if(data) printf("nonselectable\x1ftrue\x1fpermanent\x1ftrue\x1f");
	printf("icon\x1f%s\n", icon_names[icon]);
}

int back_to(char sys[], char* argv) { // allow argv passing
	// must be static to prevent stack smashing
	// allocate about a page
	// technically best in one combining buffer so all strcat in one place
	// so wait a bit for rofi close, and open error message box of rofi
	static char bash[4000] = "bash -c \"coproc (sleep 1 && rofi -e \\\"$(";
	strcat(bash, sys);
	if(argv) {
		strcat(bash, argv); // name
		strcat(bash, "\""); // yep, it's a proxy name needing a close
	}
	strcat(bash, ")\\\")\"");
	return system(bash);
}

int back(char sys[]) { // no argv call bash
	// often you'd not expect to set a proxy up yourself
	return back_to(sys, NULL);
}

int compile(int argc, char** argv) { // remake a.out
	// stderr for C errors does not go to rofi message error box
	// nice okfail macro does inline string combine of exit state to error box
	return back(config "gcc main.c" okfail);
}

int help(int argc, char** argv) { // help function
	// ok, so not fabulous help, but I'm not here to do it all for you
	return back("echo 'Ah, and a Hello Galaxy from the help function.'");
}

int (*(fn[]))(int, char**) = { // command function pointers
	// it's almost as if you'd just map a main to another name
	compile,//0
	help//1
};

int main(int argc, char** argv) {
	// loop over possible commands
	for(int i = 0; i < (sizeof(names)/sizeof(char*)); i++) {
		switch(argc - 1) {
		case 0:// list
			item(names[i], icons[i], false);// not an information item
			break;
		case 1:// process 
			if(strcmp(argv[1], names[i])) break;
			// it's how th IO gets sent to the right place instead of messing up rofi
			// allows easy C stdio
			// stderr is fine, but might no appear anywhere
			// don't emit ' or " to stdout, use apos and quot macros
			// spaces maybe contained inbetween apostrophies
			if(wrapio[i]) return back_to(config "./a.out -- \"", names[i]);// proxy call self
			int j = fn[i](argc - 1, argv + 1);// do it
			// possibly add items by item in fn data=true
			if(reload[i]) main(argc - 1, argv);// make a nested listing call
			return j;
			//ok done
		default:// process proxy (2 args)
			if(strcmp(argv[2], names[i])) break;
			return fn[i](argc - 2, argv + 2);// do it
                        //ok done
		}
	}
	// something wierd happened
	// are you calling this from rofi?
	return argc - 1; //ok or not
}