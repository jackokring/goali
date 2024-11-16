-- Keymaps are automatically loaded on the VeryLazy event
-- Default keymaps that are always set:
-- https://github.com/LazyVim/LazyVim/blob/main/lua/lazyvim/config/keymaps.lua
-- Add any additional keymaps here

-- <cmd> or : is command mode in an action
-- Some various shifts possibly free
-- Alt sends an escape prefix but also seems to check modifier state
-- :map xxx<cr> looks up mapping of a binding to xxx

--==============================================================================
-- N.B. Builtins use vim.fn prefix
local f = vim.fn
local a = vim.api
local v = vim.v
local k = vim.keymap.set

-- for adding in groups for key prefixes
---@param name string
---@param desc string
local function wk(name, desc)
  require("which-key").add({ name, group = desc })
end

-- lua_ls LSP is slightly late binding on warning of unsed function name in dict

--==============================================================================
-- first letter of name must be UPPERCASE
-- this then allows ":Com args<cr>"
-- this was considered better than allowing functions in nikey and ninkey
-- as it also allow manual execution of such functions
---@param name string first letter UPPERCASE
---@param desc string
---@param func string | function
local function com(name, desc, func)
  a.nvim_create_user_command(name, func, { desc = desc })
end

-- for the func in the command registration com to get args
---from the command line
---@param opts { fargs?: string[] }
---@return string[]
local function args(opts)
  -- return table of arg strings from opts argument
  return opts.fargs
end

--==============================================================================
-- action is function or key string (maybe recursive, careful)
---@param seq string
---@param desc string
---@param action string | function
---@param remap? boolean
local function nkey(seq, desc, action, remap)
  remap = remap or false -- nil -> false
  k("n", seq, action, { desc = desc, remap = remap })
end

-- also defines for i but ends with n mode (no func use com)
---@param seq string
---@param desc string
---@param action string
local function ninkey(seq, desc, action)
  nkey(seq, desc, action)
  k("i", seq, "<esc>" .. action, { desc = desc })
end

-- normal mode action then back to insert
---@param seq string
---@param desc string
---@param action string
local function ikey(seq, desc, action)
  -- cursor immutable possiblities by append not insert
  k("i", seq, "<esc>" .. action .. "a", { desc = desc })
end

-- remains in mode i if in i
---@param seq string
---@param desc string
---@param action string
local function nikey(seq, desc, action)
  -- can't rely on control codes below being nothing in n mode
  nkey(seq, desc, action)
  -- escape for one action step to normal mode
  ikey(seq, desc, action)
end

--==============================================================================
---passes string to _G.function named "lua_name" depending on type
---The function is called with one String argument:
---"line"	{motion} was linewise
---"char"	{motion} was charwise
---"block"	{motion} was blockwise-visual
---normal mode only, and _G.function of lua type, no ":Cmd args"
---@param seq string
---@param desc string
---@param lua_name string
local function opkey(seq, desc, lua_name)
  k("n", seq, ":set opfunc=v:lua." .. lua_name .. "<cr>g@", { desc = desc })
end

---get operator range selection within a lua_name function
---@return { r1: integer, c1: integer, r2: integer, c2: integer}
local function opval()
  -- (1, 0) indexed tuple pair (row, col)
  local r1, c1 = a.nvim_buf_get_mark(0, "[")
  local r2, c2 = a.nvim_buf_get_mark(0, "]")
  return { r1 = r1, c1 = c1, r2 = r2, c2 = c2 }
end

-- and for registers
---return a command string for delyed execution from lua_func
---also for any indirect function delayed key binding
---@param seq string
---@param desc string
---@param lua_func fun(): string
local function regkey(seq, desc, lua_func)
  k("n", seq, lua_func, { expr = true, desc = desc })
end

-- obtain the active register name
local regref = v.register
local regval = f.getreg
local regset = f.setreg

-- some factor assits for ease
---make a normal mode command by :<cr> surround
---@param str string
---@return string
local function ncom(str)
  return ":" .. str .. "<cr>"
end

---make a short telescope command for  a sub-command
---@param str string
---@return string
local function tele(str)
  return ncom("Telescope " .. str)
end

-- marks are just for moving about
-- quite difficult to get a "seamless dynamic mark"" letter in functions
-- and there's the lower case a-z local file, A-Z global marks and ...

--==============================================================================
-- Bare Sparse Escape (Not in use)
-- benou
-- ABCDEFGHIJKLMNOPQRSTUVWXYZ
-- normal launch rofi, as <C-R> register recall in i mode, redo n mode
wk("\\", "user key")
nkey("\\\\", "Launch by Rofi-combi", ncom("!rofi -show combi"))
nkey("\\a", "Select all", "ggVG")
nkey("\\b", "Begin of line", "^")
nkey("\\c", "Commands", tele("commands"))
nkey("\\d", "Diagnostics", tele("diagnostics"))
nkey("\\e", "End of line", "$")
nkey("\\f", "Find in buffer", tele("current_buffer_fuzzy_find"))
nkey("\\g", "Search replace file", ":%s/\\v")
nkey("\\h", "History of commands", tele("command_history"))
nkey("\\i", "Insert ASCII art", ":r !figlet ")
nkey("\\j", "Line down", ":m .+1<cr>==")
nkey("\\k", "Line up", ":m .-2<cr>==")
nkey("\\l", "Run lua", ":lua ")
nkey("\\m", "Messages", ncom("Noice telescope"))
nkey("\\n", "Next warning", "]w", true) -- remap?
nkey("\\o", "Pure function", "vaf")
nkey("\\p", "Paste", tele("registers"))
nkey("\\q", "Quit save all", ncom("xa"))
nkey("\\r", "Reload package", tele("reloader"))
nkey("\\s", "Search replace line", ":s/\\v")
nkey("\\t", "Treesitter symbols", tele("treesitter"))
nkey("\\u", "Upper/lower case word", "viw~")
nkey("\\v", "Visual line", "V")
nkey("\\w", "Window cycle", "<C-w><C-w>")
nkey("\\x", "Make executable", ncom("!chmod +x %"))
nkey("\\y", "Yank line", "Vy")
nkey("\\z", "Undo/Redo Paragraph Comment", "vipgc")

--==============================================================================
-- Leader Space (Many used, see used by pressing <space> in normal mode)
-- aijkmnoprvyz
-- ABCFGIJMNOPQRSTUVWXYZ
wk("<leader>", "leader key")
-- what olde one eye said
-- snh nkey("<leader>m", "Message History", ":Noice<cr>")

--==============================================================================
-- Control (All used in some way, but just a few remaps)
-- Can be in insert mode as wrapped <esc> .. i by <C-\><C-O> or just <esc>
-- apparently terminal built in does not do terminal <C-/> works as <C-_>
-- GNO are used N for normal, O for temp normal, G for backward compatibility
-- after a <C-\> and it appears to be hard wired
-- save all <C-S> not just save one file and remain in mode
nikey("<C-S>", "Save all", ncom("wa"))
-- reload and place in n mode
ninkey("<C-Z>", "Revert to saved", ncom("e!"))
-- kill the LSP reach for effect (see y?)
-- just a consequence entering normal mode
-- I tend never to use obscure <C-\> combinations but get annoyed by
-- an open cmp dialog interfering with the cursor in insert mode
-- also a minor error 444 but sometimes close crashed pop-up good
-- i'm not a multi-window man, more tab based
ikey("<C-\\>", "Close LSP cmp, etc.", "") -- basic insert norm insert effect
nkey("<C-\\>", "Close pop-up", ncom("close"))

--==============================================================================
-- Alt (Very rare, only JKNP seem bound by default)
-- Use <M-?> for key ? input string, becomes <esc><?> CSI combination
-- can be both insert and normal mode ni/ninkey depending on mode on exit

--==============================================================================
-- Perculiar mode keys
-- for things like visual mode or visual line mode additions
-- see:
-- https://github.com/LazyVim/LazyVim/blob/main/lua/lazyvim/config/keymaps.lua
