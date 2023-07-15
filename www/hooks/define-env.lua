local json = require("json")
local strings = require("strings")

function interp(s, tab)
    return (s:gsub("\r?\n", " "):gsub('($%b{})', function(w) return tab[w:sub(3, -2)] or w end))
end

function file_exists(file)
    local f = io.open(file, "rb")
    if f then f:close() end
    return f ~= nil
end

function lines_from(file)
    if not file_exists(file) then return {} end
    local lines = {}
    for line in io.lines(file) do 
      lines[#lines + 1] = line
    end
    return lines
end


function Writer(filedata)
    local source_data = json.decode(filedata);

    local envLines = lines_from(".env")
    local envData = {}

    for k,v in pairs(envLines) do
        local envKeyVal = strings.split(v,"=")
        envData[envKeyVal[1]]=envKeyVal[2]
    end
    

    return json.encode({
		data = {
			env = envData
		},
	})
end