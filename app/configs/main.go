package configs

import (
	. "api/pkg/configs"
	_ "api/pkg/dotenv"
)

var DATABASE_URL = Get("DATABASE_URL", "mongodb://fproot:xpto2318@localhost:27017/?readPreference=secondaryPreferred&retryWrites=false")
