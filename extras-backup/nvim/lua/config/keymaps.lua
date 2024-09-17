-- Keymaps are automatically loaded on the VeryLazy event
-- Default keymaps that are always set: https://github.com/LazyVim/LazyVim/blob/main/lua/lazyvim/config/keymaps.lua
-- Add any additional keymaps here

-- <cmd> or : is command mode in an action
-- Some various shifts possibly free
-- Alt sends an escape prefix

-- Bare Sparse Escape (Not in use)
-- abcdefghijklmnopqrstuvwxyz
-- ABCDEFGHIJKLMNOPQRSTUVWXYZ
vim.keymap.set('n', '\\a', '')

-- Leader Space (Many used, see use pressing <space> in normal mode)
-- adhijkmnoprtvyz
-- ABCFGHIJMNOPQRSTUVWXYZ
vim.keymap.set('n', '<Leader>a', '')

-- Control (Lowercase RESERVED for plugins with no control, uppercasw free with no control but shifted)
-- ABCDEFHIJKLMOPQRSTUVWXYZ (with control as easiest to finger)
-- Perculiar shift combination needed
vim.keymap.set('n', '<C-_>', '')
-- NOT <C-N> or <C-G> but rest of controls and not lowercase
vim.keymap.set('n', '<C-\\><C-A>', '')
-- Control+space <C-@>
vim.keymap.set('n', '<C-@>', '')
