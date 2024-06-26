package hyper

import "github.com/Luzifer/go-openssl"

func AESEncrypt(plaintext string, secret string) (string, error) {
	o := openssl.New()

	salt, err := o.GenerateSalt()
	if err != nil {
		return "", err
	}

	enc, err := o.EncryptBytesWithSaltAndDigestFunc(secret, salt, []byte(plaintext), openssl.DigestMD5Sum)
	if err != nil {
		return "", err
	}

	return string(enc), nil
}

func AESDecrypt(plaintext string, secret string) (string, error) {
	o := openssl.New()

	dec, err := o.DecryptBytes(secret, []byte(plaintext))
	if err != nil {
		return "", err
	}

	return string(dec), nil
}
