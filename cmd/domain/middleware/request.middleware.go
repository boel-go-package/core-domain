package middleware

var (
	active bool   = true
	mss    string = "Server is not available consume request"
)

func GetRequestMessage() string {
	return mss
}

func SetRequestMessage(msg string) {
	mss = msg
}

func GetRequestActive() bool {
	return active
}

func RequestActive() {
	active = true
}

func RequestInactive() {
	active = false
}
