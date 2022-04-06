package worker

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"last9/schema"
	"last9/store"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

type sshWorker struct {
	basePath string
}

func NewSSH() *sshWorker {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(homedir)
	ssh := &sshWorker{
		basePath: homedir + "/.ssh/keys",
	}
	if _, err := os.Stat(ssh.basePath); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(ssh.basePath, os.ModePerm); err != nil {
			log.Println(err)
		}
	}

	return ssh
}

func (s *sshWorker) Do() {
	log.Println("SSH worker started")
	ec2Instns, err := store.Store.EC2Instances().All()
	if err != nil {
		log.Println(err)
		return
	}

	for _, ec2Inst := range ec2Instns {
		if ec2Inst.State == schema.EC2InstanceStateRunning {
			go s.generateSSHKeys(ec2Inst.InstanceID)
		}
	}
}

func (s *sshWorker) generateSSHKeys(ec2InstanceID string) {
	pubKeyPath := fmt.Sprintf("%s/%s.pub", s.basePath, ec2InstanceID)
	privKeyPath := fmt.Sprintf("%s/%s", s.basePath, ec2InstanceID)
	if _, err := os.Stat(pubKeyPath); !os.IsNotExist(err) {
		log.Println("SSH key already exists")
		return
	}

	log.Println("Generating SSH files for ", ec2InstanceID)
	pubKey, privKey, err := s.generateKeys()
	if err != nil {
		log.Println(err)
		return
	}
	// fmt.Println("my public key is...")
	// fmt.Println(string(pubKey))
	// fmt.Println("my private key is...")
	// fmt.Println(string(privKey))

	if err := s.writeToFile(pubKeyPath, pubKey); err != nil {
		log.Println(err)
		return
	}
	if err := s.writeToFile(privKeyPath, privKey); err != nil {
		log.Println(err)
		return
	}
}

func (s *sshWorker) generateKeys() ([]byte, []byte, error) {
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		return nil, nil, err
	}

	pub, err := ssh.NewPublicKey(key.Public())
	if err != nil {
		return nil, nil, err
	}
	pubKeyStr := ssh.MarshalAuthorizedKey(pub)
	privKeyStr := s.marshalRSAPrivate(key)

	return pubKeyStr, privKeyStr, nil
}

func (s *sshWorker) marshalRSAPrivate(priv *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv),
	})
}

// writePemToFile writes keys to a file
func (s *sshWorker) writeToFile(filePath string, keyBytes []byte) error {
	if err := ioutil.WriteFile(filePath, keyBytes, 0600); err != nil {
		return err
	}

	log.Printf("Key saved to: %s", filePath)
	return nil
}
