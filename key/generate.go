package key

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"strings"
)

type Key struct {
	Public  string
	Private string
	priKey  *ecdsa.PrivateKey
	puKey   *ecdsa.PublicKey
}

func GenKeyPairT() *Key {
	priKey, e := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if e != nil {
		return nil
	}

	return &Key{
		Public:  hex.EncodeToString(priKey.X.Bytes()),
		Private: hex.EncodeToString(priKey.Y.Bytes()),
		priKey:  priKey,
		puKey:   &priKey.PublicKey,
	}
}

func (k *Key) Sign(args ...string) (string, error) {
	data := strings.Join(args[:], "")
	test, _ := ecdsa.SignASN1(rand.Reader, k.priKey, []byte(data))

	return hex.EncodeToString([]byte(test)), nil
}

func (k *Key) Verify(public string, sign string, args ...string) bool {
	data := strings.Join(args[:], "")
	//signuture := hex.EncodeToString([]byte(data))
	signuture, _ := hex.DecodeString(sign)

	if public != k.Public {
		return false
	}

	return ecdsa.VerifyASN1(k.puKey, []byte(data), signuture)
}
