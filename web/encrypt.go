package web

import (
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"fmt"

	"github.com/aws/aws-sdk-go/service/kms"
	"golang.org/x/crypto/nacl/secretbox"
)

const (
	nonceLength = 24
	keyLength   = 32
)

var keySpec = "AES_256"

type payload struct {
	Key     []byte
	Nonce   *[nonceLength]byte
	Message []byte
}

func encrypt(kmscl *kms.KMS, keyID string, msg []byte) ([]byte, error) {
	keySpec := "AES_128"
	in := &kms.GenerateDataKeyInput{KeyId: &keyID, KeySpec: &keySpec}
	out, err := kmscl.GenerateDataKey(in)
	if err != nil {
		return nil, err
	}
	p := &payload{
		Key:   out.CiphertextBlob,
		Nonce: &[nonceLength]byte{},
	}

	// Set nonce
	if _, err = rand.Read(p.Nonce[:]); err != nil {
		return nil, err
	}

	// Create key
	key := &[keyLength]byte{}
	copy(key[:], out.Plaintext)

	// Encrypt message
	p.Message = secretbox.Seal(p.Message, msg, p.Nonce, key)

	// TODO: consider alternative that does not box them together, instead relying on the url containing the encrypted key
	buf := &bytes.Buffer{}
	if err := gob.NewEncoder(buf).Encode(p); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func decrypt(kmscl *kms.KMS, ciphertext []byte) ([]byte, error) {
	var p payload
	err := gob.NewDecoder(bytes.NewReader(ciphertext)).Decode(&p)
	if err != nil {
		return nil, err
	}

	out, err := kmscl.Decrypt(&kms.DecryptInput{CiphertextBlob: p.Key})
	if err != nil {
		return nil, err
	}

	key := &[keyLength]byte{}
	copy(key[:], out.Plaintext)

	// Decrypt message
	var decryptedMsg []byte
	decryptedMsg, ok := secretbox.Open(decryptedMsg, p.Message, p.Nonce, key)
	if !ok {
		return nil, fmt.Errorf("failed to open secretbox")
	}
	return decryptedMsg, nil
}
