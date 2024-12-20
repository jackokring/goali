-- custom completions
local source = {}

local function fetch_autojump_dirs()
  local scripts = vim.fn.system("awk -F '\t' '{print $2}' ~/.local/share/autojump/autojump.txt") -- input nil
  -- vim.notify(scripts, vim.log.levels.DEBUG)
  return vim.split(scripts, "\n")
end

-- function source:is_available()
-- return true
-- end

-- function source:get_debug_name()
-- return "projects"
-- end

-- function source:get_keyword_pattern()
-- return [[\k\+]]
-- end

-- Trigger completion (i.e. open up cmp) on these characters.
-- We can also trigger it manually, see `:help cmp.mapping.complete`.
function source:get_trigger_characters()
  return { "/" }
end

---Invoke completion (required).
---@param params cmp.SourceCompletionApiParams
---@param callback fun(response: lsp.CompletionResponse|nil)
function source:complete(params, callback)
  -- There's also the `cursor_after_line`, `cursor_line`, and a `cursor` fields on `context`.
  local cursor_before_line = params.context.cursor_before_line
  local lbl = {}
  -- Only complete if there's a `/` before the cursor.
  if cursor_before_line:sub(-1, -1) == "/" then
    local project_names = fetch_autojump_dirs()
    for _, v in ipairs(project_names) do
      table.insert(lbl, {
        -- in list
        label = v,
        -- icon
        kind = require("cmp.types.lsp").CompletionItemKind.Folder,
        -- but add quotes insert
        insertText = '"' .. v .. '"',
        -- text to filter on
        filterText = v,
      })
    end
    callback(lbl)
  else
    callback()
  end
end

function source:resolve(item, callback)
  -- provide documentation popup
  item.documentation = {
    kind = "markdown",
    value = "An autojump directory\n",
  }
  callback(item)
end

local function reg(name, src)
  local cmp = require("cmp")
  cmp.register_source(name, src)
end

reg("autojump", source)
