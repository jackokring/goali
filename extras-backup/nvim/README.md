# 💤 LazyVim

A starter template for [LazyVim](https://github.com/LazyVim/LazyVim).
Refer to the [documentation](https://lazyvim.github.io/installation) to get started.

## Things I've Found Out that are Not Obvious

### Coding Language Install

So to install languages, it is better to reopen Neovim and select Lazy Extras
from the opening screen. This includes various language programming setups
where all the dependencies will install.

### Pressing "<leader><bs>" For `which-key`

This should bring up a list of `i` mode key bindings, and it does to some
extent. It however does not include them all, and suggests some key binds
are handled at a lower level. This is a little confusing for control key
combinations.

### Free Leader Keys

After quite a bit of fiddling around, I used `<leader>m` for seeing the message
history of Noice, and added in a `\` user leader for other things. Check with
`keymaps.lua` for more information.
