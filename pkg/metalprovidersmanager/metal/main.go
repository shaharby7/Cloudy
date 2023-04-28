package metal

type Metal struct {
	provider string
	ip       string
	pem      string
}

func (m *Metal) Ssh() {
	// TODO
}

func (m *Metal) Shutdown() {
	// TODO
}
