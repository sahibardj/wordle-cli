package colours

const reset = "\033[0m"
const (
	Code_grey   = -2
	Code_red    = -1
	Code_yellow = 0
	Code_green  = 1
)

func Bold(str string) string {
	return "\033[1m" + str + reset
}

func Red(str string) string {
	return "\033[31m" + str + reset
}

func Green(str string) string {
	return "\033[32m" + str + reset
}

func Yellow(str string) string {
	return "\033[33m" + str + reset
}

func Grey(str string) string {
	return "\033[37m" + str + reset
}
