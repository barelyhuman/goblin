local json = require("json")
local alvu = require("alvu")

function Writer(filedata)
    local envData = {}

    envData["GOBLIN_ORIGIN_URL"] = alvu.get_env(".env","GOBLIN_ORIGIN_URL")

    return json.encode({
		data = {
			env = envData
		},
	})
end