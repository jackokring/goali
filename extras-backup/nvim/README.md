# 💤 LazyVim

A starter template for [LazyVim](https://github.com/LazyVim/LazyVim).
Refer to the [documentation](https://lazyvim.github.io/installation) to get started.

## Things I've Found Out that are Not Obvious

### Coding Language Install

So to install languages, it is better to reopen Neovim and select Lazy Extras
from the opening screen. This includes various language programming setups
where all the dependencies will install.

### Pressing `<leader><bs>` For `which-key`

This should bring up a list of `n` mode key bindings, and it does to some
extent. It however does not include them all, and suggests some key binds
are handled at a lower level. This is a little confusing for control key
combinations.

### Free Leader Keys

After quite a bit of fiddling around, I used `<leader>m` for seeing the message
history of Noice, and added in a `\` user leader for other things. Check with
`keymaps.lua` for more information. This is `<leader>snh` though, so I removed
it again.

### Module Development

The module `doris.nvim` is added to the config to show how local development
needs `return` values for `dev`. This doesn't enable local testing but the
github action does run, and version tagging by `git tag v*` and pushing
by `git push origin --tags` to run the luarocks submit if the API key is
configured.

### LuaSnips Looks Interesting

I added a load from `init.lua` by `require("config.snips")` to require, for
complex or simple insertion templates.
