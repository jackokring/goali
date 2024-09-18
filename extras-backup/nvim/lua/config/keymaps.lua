-- Keymaps are automatically loaded on the VeryLazy event
-- Default keymaps that are always set: https://github.com/LazyVim/LazyVim/blob/main/lua/lazyvim/config/keymaps.lua
-- Add any additional keymaps here

-- <cmd> or : is command mode in an action
-- Some various shifts possibly free
-- Alt sends an escape prefix

local function iwrap(func)
  if vim.mode(0) ~= "i" then
    vim.call(func, nil)
    return
  else
    -- stop insert mode for a bit
    local cur = vim.getcurpos()
    vim.cmd.stopinsert()
    vim.call(func, nil)
    vim.setpos(".", cur)
    -- restore, ah yes, the end of line
    vim.cmd.startinsert()
  end
end

local function nkey(seq, desc, action)
  vim.keymap.set({ "n", "v" }, seq, action, { desc = desc })
end

-- action must be a function for this one
local function nikey(seq, desc, action)
  vim.keymap.set({ "n", "v", "i" }, seq, function()
    iwrap(action)
  end, { desc = desc })
end

-- Bare Sparse Escape (Not in use)
-- abcdefghijklmnopqrstuvwxyz
-- ABCDEFGHIJKLMNOPQRSTUVWXYZ
nkey("\\a", "", "")

-- Leader Space (Many used, see use pressing <space> in normal mode)
-- adhijkmnoprtvyz
-- ABCFGHIJMNOPQRSTUVWXYZ
nkey("<Leader>a", "", "")

-- Control (Lowercase RESERVED for plugins with no control, uppercase free with no control but shifted)
-- Can be in insert mode sometimes as control and not <c-v> literal prefixed
-- ABCDEFHIJKLMOPQRSTUVWXYZ (with control as easiest to finger)
-- Perculiar shift combination needed
nikey("<C-_>", "", function()
  -- nil
end)
-- NOT <C-N> or <C-G> but rest of controls and not lowercase
nikey("<C-\\><C-A>", "", function()
  -- nil
end)
-- Uppercase with shift
nkey("<C-\\><S-A>", "", "")
