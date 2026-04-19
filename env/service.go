package env

func New(typesPaths *[]string) Env {
	return &envImpl{
		typesPaths: typesPaths,
	}
}
