package instrumenter

type Instrumenter interface {
	Instrument(file string) ([]byte, error)
}
