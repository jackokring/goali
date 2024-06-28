#!/usr/bin/bash
echo "Maybe:"
echo "rustup target add aarch64-unknown-linux-gnu"
echo "sudo apt install gcc-aarch64-linux-gnu"
# gnu_version crate_name
# ./cargo-cross.sh 11 just
PKG_CONFIG_SYSROOT_DIR=/.
PKG_CONFIG_PATH=/usr/lib/aarch64-linux-gnu/pkgconfig/
PKG_CONFIG_ALLOW_CROSS=1
RUSTFLAGS="-Clinker=aarch64-linux-gnu-ld -L /usr/lib/gcc-cross/aarch64-linux-gnu/${1}/" cargo install $2 --root ~/bin-arm64/ --target aarch64-unknown-linux-gnu
#Mint doesn't support arm64 arch files but ... debian does, maybe alt sources []
#sudo dpkg --add-architecture arm64
# TEST STILL NOT DOING ALSA -- FAIL
#sudo apt install libasound2-dev:arm64

# avoid error on example list to source multiarch
exit 0
# a multi source list
deb [arch=amd64,i386] https://www.mirrorservice.org/sites/packages.linuxmint.com/packages virginia main upstream import backport 

deb [arch=amd64,i386] http://archive.ubuntu.com/ubuntu jammy main restricted universe multiverse
deb [arch=amd64,i386] http://archive.ubuntu.com/ubuntu jammy-updates main restricted universe multiverse
deb [arch=amd64,i386] http://archive.ubuntu.com/ubuntu jammy-backports main restricted universe multiverse

deb [arch=amd64,i386] http://security.ubuntu.com/ubuntu/ jammy-security main restricted universe multiverse

deb [arch=arm64] http://ports.ubuntu.com/ubuntu-ports jammy main restricted universe multiverse
deb [arch=arm64] http://ports.ubuntu.com/ubuntu-ports jammy-updates main restricted universe multiverse
deb [arch=arm64] http://ports.ubuntu.com/ubuntu-ports jammy-backports main restricted universe multiverse
deb [arch=arm64] http://ports.ubuntu.com/ubuntu-ports jammy-security main restricted universe multiverse
deb [arch=arm64] http://archive.canonical.com/ubuntu jammy partner
