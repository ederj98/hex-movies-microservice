package exception

type Validator interface {
	Error() string
	IsBusinessLogic() bool
}
