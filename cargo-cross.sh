#!/usr/bin/bash
echo "Maybe:\n"
echo "rustup target add aarch64-unknown-linux-gnu\n"
echo "sudo apt install gcc-aarch64-linux-gnu\n"
# gnu_version crate_name
# ./cargo-cross.sh 11 just
RUSTFLAGS="-Clinker=aarch64-linux-gnu-ld -L /usr/lib/gcc-cross/aarch64-linux-gnu/${1}/" cargo install $2 --root ~/bin-arm64/ --target aarch64-unknown-linux-gnu
