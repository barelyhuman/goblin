require("dotenv").config();

module.exports = function () {
	return {
		originUrl: process.env.GOBLIN_ORIGIN_URL || "http://goblin.run",
	};
};
