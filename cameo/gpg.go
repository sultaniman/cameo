package cameo

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/openpgp"
	"io/ioutil"
)

func (g *GPG) Encrypt(message []byte) ([]byte, error) {
	hints := &openpgp.FileHints{IsBinary: true}
	encryptBuf := new(bytes.Buffer)

	// Here we are not signing the message because we only want to
	// encryptBuf it and recipient should have private key to decrypt.
	wc, err := openpgp.Encrypt(encryptBuf, g.Entities, nil, hints, nil)
	if err != nil {
		return nil, err
	}

	_, err = wc.Write(message)
	_ = wc.Close()
	if err != nil {
		return nil, err
	}

	encrypted, err := ioutil.ReadAll(encryptBuf)
	if err != nil {
		return nil, err
	}
	fmt.Println(encrypted)
	return encrypted, nil
}
