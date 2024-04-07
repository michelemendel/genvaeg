package constants

// Env keys
const (
	ENV_APP_ENV_KEY         = "APP_ENV"
	ENV_WEB_SERVER_PORT_KEY = "WEB_SERVER_PORT"
	ENV_SESSION_KEY_KEY     = "SESSION_KEY"
	ENV_DB_DIR_KEY          = "DEV_DB_DIR"
	ENV_DB_NAME_KEY         = "DB_NAME"
	ENV_BYPASS_LOGIN        = "BYPASS_LOGIN"
)

// Auth keys
const (
	AUTH_SESSION_NAME = "session"
	AUTH_TOKEN_NAME   = "token"
)

// Routes
const (
	//
	ROUTE_SIGNUP = "/signup"
	ROUTE_LOGIN  = "/login"
	ROUTE_LOGOUT = "/logout"
	//
	ROUTE_CREATE_SHORT_URL = "/create"
	ROUTE_LIST_SHORT_URLS  = "/list"
	ROUTE_REDIRECT         = "/rd"
	//
	ROUTE_PING = "/ping"
)
