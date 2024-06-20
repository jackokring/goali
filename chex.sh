#!/usr/bin/bash
pushd ~/goali
for i in *.sh
do
	chmod +x "$i"
done
popd

