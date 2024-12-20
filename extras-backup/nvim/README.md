# 💤 LazyVim

A starter template for [LazyVim](https://github.com/LazyVim/LazyVim).
Refer to the [documentation](https://lazyvim.github.io/installation) to get started.

## Layout of Files

- `after/ftplugin/<lang>.lua` for `vim.opt_local.<setting>` per language.
- `snips/<lang>.lua` for LuaSnip definitions (auto completions).
- `lua/config/<file>.lua` for running to setup things (no return).
- `lua/plugins/<file>.lua` return a spec for loading a plugin (always return spec).

## Installing Languages

Use the `install extras` from the home screen in preference to `mason`. It
might need quit then restart as it does not appear available after opening or
restoring a session. Occasionally, open `mason` by `<leader>cm` to do any
updates, as it has its own updating procedure. Most of the time the `extras` from
`LazyVim` selects the best plugins from `mason` for particular languages but
you might just want a very specific language tool.

## Plugin Development

The module `doris.nvim` is added to the config to show how local development
works for a plugin spec.
