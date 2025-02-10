package env

func NewEnvService(typesPaths *[]string) IEnv {
	return &SEnv{
		typesPaths: typesPaths,
	}
}
