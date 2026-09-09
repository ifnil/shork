package host

type Host struct {
	User       string
	HostName   string
	Port       string
	PrivateKey string
}

func NewHost() *Host {
	return &Host{}
}
