-- doris plugin loader for nvim
return {
  "jackokring/doris.nvim",
  -- **local build**
  name = "doris.nvim",
  dev = {
    path = "~/projects",
    patterns = { "jackokring" },
    fallback = true,
  },
  -- **dependancies**
  dependencies = {
    "nvim-lua/plenary.nvim",
  },
  -- **option overrides, implies calling .setup(opts)**
  opts = {},
  -- **on load build command**
  -- build = "",
  -- **on event**
  -- event = { "BufEnter", "BufEnter *.lua" },
  -- **on command use**
  -- cmd = { "cmd" },
  -- **on filetype**
  -- ft = { "lua" },
  -- **on keys**
  -- keys = {
  --    -- key tables
  --    {
  --      "<leader>ft",
  --      -- "<cmd>Neotree toggle<cr>",
  --      -- desc = "NeoTree",
  --      -- mode = "n",
  --      -- ft = "lua"
  --    },
  --  },
}
