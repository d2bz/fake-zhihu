package xcode

var (
	OK                = Define(0, "ok")
	NoLogin           = Define(101, "no login")
	RequestErr        = Define(400, "request error")
	Unauthorized      = Define(401, "unauthorized")
	AccessDenied      = Define(403, "access denied")
	NotFound          = Define(404, "not found")
	MethodNotAllowd   = Define(405, "method not allowed")
	Canceled          = Define(498, "canceled")
	ServerErr         = Define(500, "server error")
	SeviceUnavaliable = Define(503, "service unavailable")
	DeadlineExceeded  = Define(504, "deadline exceeded")
	LimitExceeded     = Define(509, "limit exceeded")
)
