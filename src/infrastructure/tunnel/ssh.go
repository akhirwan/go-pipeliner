package tunnel

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type SSHConfig struct {
	Name               string // Tunnel Name for identification
	Host               string
	Port               int
	User               string
	Password           string
	PrivateKeyFile     string // Blank if it doesn't use PrivateKeyFile
	PrivateKeyPassword string // Blank If key is not encrypted with password
}

// CreateSSHConfig that will be used by MySQL Config
func (t *SSHConfig) CreateSSHTunnel() *ssh.Client {
	var authMethods []ssh.AuthMethod
	// AuthMethod using sshagent
	if netConn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		authMethods = append(authMethods, ssh.PublicKeysCallback(agent.NewClient(netConn).Signers))
	}

	// Authmethod using PrivateKeyFile
	if t.PrivateKeyFile != "" {
		// Read private file key
		pemBytes, err := ioutil.ReadFile(t.PrivateKeyFile)
		if err != nil {
			log.Fatalf("Reading private key file failed %v", err)
		}

		// create signer
		signer, err := ssh.ParsePrivateKey(pemBytes)
		// if key encrypted with password
		// signer, err := ssh.ParsePrivateKeyWithPassphrase(key, []byte(privateKeyPass))
		if err != nil {
			log.Fatal("Sign priv key failed", err)
		}

		// Add PublicKey to ssh auth when using PrivateKeyFile
		authMethods = append(authMethods, ssh.PublicKeys(signer))
	}

	// Auth Method 3 using password
	if t.Password != "" {
		authMethods = append(authMethods, ssh.PasswordCallback(func() (string, error) {
			return t.Password, nil // Fill prompt password
		}))
	}

	// The client configuration with configuration option to use the ssh-agent
	sshConfig := &ssh.ClientConfig{
		User: t.User,
		Auth: authMethods,
		// Ignore security
		//HostKeyCallback: ssh.InsecureIgnoreHostKey()
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", t.Host, t.Port), sshConfig)
	if err != nil {
		log.Printf("Dialing SSH failed: %s", err.Error())
		return nil
	}

	log.Println("SSH Connected")

	return sshClient
}
