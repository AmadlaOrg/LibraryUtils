package desktop

// Generator 🧩 defines the interface for generating `.desktop` files.
type Generator interface {
	Generate(desktop *Desktop) (Generator, error)
	WriteToFile(path string) error
}
