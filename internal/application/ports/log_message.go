package ports

type LogMessage interface {
	Info(msg string)
	Detail(msg string)
	Success(msg string)
	Error(msg string)
}
