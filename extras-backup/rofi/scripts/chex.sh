#!/usr/bin/bash
for i in *.sh
do
	chmod +x "$i"
done
# build rofi a(.out) script plugin
# this will appear in the -modi-combi "a" mode
# can add C routines just by adapting
# int main() of sokme other source
# making the command arrays use it renamed
# it's capable of compiling itself
# search "compile" in rofi search
gcc main.c
