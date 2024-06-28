---
layout: post
title:  "Cargo Cross Compiling"
date:   2024-06-28 19:00:00 +0000
categories: rust compile
---
So I made a `cargo-cross.sh` file and mashed up Linux Mint sources a little to add in `arm64` from `jammy`. Running an alternate architecture? No, I'm just making a cross compile system as my Chromebook (arm64) does not have enough "Rust Gust Oompf". I could say the only type of memory errors I get with rust was out of memory ones. But transferring the build process onto another box (amd64) was possible, and even in hind sight easier than imagined.

This allows me to use `cargo` as a package manager. Build remote, and then load up the binaries for use. Buiklding against a slightly older LTS libc and friends does help. Go doesn't need any of that stuff, as it's a compact development footprint. For later the server includes `nix`, `flatpak`, `cargo` and of course the usual `apt`. If you believe the `$PATH` it still has some vestigial `/snap`, but I'm not going there. I might not even go `.AppImage`.

It's still all linux, and depending on the architecture "port" package servers, and setting `$ARC` in the script, it will work on multiarch, but might need persuasion to do `MinGW-w64` or `macOS`. But I have no need myself for other cross compiles at the moment.

I also added customizations of some `extras-backup` using `./extras-install.sh` to opt in to various customization files. Today I added `neofetch` customization. I know the C written `fastfetch` is faster, but I don't use it often enough to have a need for speed at the expense of a larger software footprint. The `emacs` config is getting better. It now has default `(tab-bar-mode)`. Emacs is kind of special in that a tab is not a unique open file buffer, but more of a "panel group" (emacs window set), but all buffers can be viewed or rotated through in any tab. Cool.

I think next I should setup and improve the language modes in emacs, and perhaps add some spell checking as I type, as that's something I'd perhaps like.
