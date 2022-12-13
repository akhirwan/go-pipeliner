package tunnel

type AuthMethod struct {
	PrivateKeyFile string
	PassPhrase     string
}

func NewAuthMethod(method string) {

}

func PrivateKeyFile(path string, passPhrase string) {}
func Password(password string)                      {}
func SSHAgent()                                     {}
