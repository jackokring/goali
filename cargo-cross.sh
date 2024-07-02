#!/usr/bin/bash
echo "Maybe to Add 'arm64' Architecture:"
echo "  rustup target add aarch64-unknown-linux-gnu"
echo "  sudo apt install gcc-aarch64-linux-gnu"
echo "  sudo dpkg --add-architecture arm64"
echo "  sudo apt install lib\${name}-dev:arm64 ..."
# gnu_version crate_name
# ./cargo-cross.sh 13 just
# The recent tool kit must not exceed the aarch platform libs
ARC=aarch64-linux-gnu
export SYSROOT=/usr/${ARC}
export PKG_CONFIG_ALLOW_CROSS=1
export PKG_CONFIG_LIBDIR=/usr/lib/${ARC}/pkgconfig
export PKG_CONFIG_SYSROOT_DIR=${SYSROOT}
export PKG_CONFIG_SYSTEM_LIBRARY_PATH=/usr/lib/${ARC}
export PKG_CONFIG_SYSTEM_INCLUDE_PATH=/usr/${ARC}/include
RUSTFLAGS="-Clinker=${ARC}-ld -L /usr/lib/gcc-cross/${ARC}/${1}/" cargo install $2 --root ~/bin-arm64/ --target aarch64-unknown-linux-gnu
#Mint doesn't support arm64 arch files but ... debian does, maybe alt sources []
# TEST OK DOING ALSA
#sudo apt install libasound2-dev:arm64

# avoid error on example list to source multiarch
exit 0
# a multi source list
deb [arch=amd64,i386] https://www.mirrorservice.org/sites/packages.linuxmint.com/packages wilma main upstream import backport 

deb [arch=amd64,i386] http://archive.ubuntu.com/ubuntu noble main restricted universe multiverse
deb [arch=amd64,i386] http://archive.ubuntu.com/ubuntu noble-updates main restricted universe multiverse
deb [arch=amd64,i386] http://archive.ubuntu.com/ubuntu noble-backports main restricted universe multiverse

deb [arch=amd64,i386] http://security.ubuntu.com/ubuntu/ noble-security main restricted universe multiverse

deb [arch=arm64] http://ports.ubuntu.com/ubuntu-ports noble main restricted universe multiverse
deb [arch=arm64] http://ports.ubuntu.com/ubuntu-ports noble-updates main restricted universe multiverse
deb [arch=arm64] http://ports.ubuntu.com/ubuntu-ports noble-backports main restricted universe multiverse
deb [arch=arm64] http://ports.ubuntu.com/ubuntu-ports noble-security main restricted universe multiverse

