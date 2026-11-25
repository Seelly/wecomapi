package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"io"
	"math"

	"github.com/seelly/wecomapi/internal/lowlevel/pkcs7"
)

type Payload struct {
	Msg       []byte
	ReceiveID []byte
}

type Encryptor struct {
	aesKey        []byte
	entropySource io.Reader
}

type EncryptorOption interface {
	applyTo(x *Encryptor)
}

type customEntropySource struct {
	inner io.Reader
}

func WithEntropySource(e io.Reader) EncryptorOption {
	return &customEntropySource{inner: e}
}

func (o *customEntropySource) applyTo(x *Encryptor) {
	x.entropySource = o.inner
}

var errMalformedEncodingAESKey = errors.New("malformed EncodingAESKey")
var errPayloadTooBig = errors.New("payload too big")

func NewEncryptor(
	encodingAESKey string,
	opts ...EncryptorOption,
) (*Encryptor, error) {
	aesKey, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")
	if err != nil {
		return nil, err
	}

	if len(aesKey) != 32 {
		return nil, errMalformedEncodingAESKey
	}

	obj := Encryptor{
		aesKey:        aesKey,
		entropySource: rand.Reader,
	}
	for _, o := range opts {
		o.applyTo(&obj)
	}

	return &obj, nil
}

func (e *Encryptor) Decrypt(base64Msg []byte) (Payload, error) {
	// base64 decode
	buflen := base64.StdEncoding.DecodedLen(len(base64Msg))
	buf := make([]byte, buflen)
	n, err := base64.StdEncoding.Decode(buf, base64Msg)
	if err != nil {
		return Payload{}, err
	}
	buf = buf[:n]

	// init cipher
	block, err := aes.NewCipher(e.aesKey)
	if err != nil {
		return Payload{}, err
	}

	iv := e.aesKey[:16]
	state := cipher.NewCBCDecrypter(block, iv)

	// decrypt in-place in the allocated temp buffer
	state.CryptBlocks(buf, buf)
	buf = pkcs7.Unpad(buf)

	// assemble decrypted payload
	// drop the 16-byte random prefix
	msglen := binary.BigEndian.Uint32(buf[16:20])
	msg := buf[20 : 20+msglen]
	receiveID := buf[20+msglen:]

	return Payload{
		Msg:       msg,
		ReceiveID: receiveID,
	}, nil
}

func (e *Encryptor) prepareBufForEncryption(payload *Payload) ([]byte, error) {
	resultMsgLen := 16 + 4 + len(payload.Msg) + len(payload.ReceiveID)

	// allocate buffer
	buf := make([]byte, 16, resultMsgLen)

	// add random prefix
	_, err := io.ReadFull(e.entropySource, buf) // len(buf) == 16 at this moment
	if err != nil {
		return nil, err
	}

	buf = buf[:cap(buf)] // grow to full capacity
	msglen := uint64(len(payload.Msg))
	if msglen > uint64(math.MaxUint32) {
		return nil, errPayloadTooBig
	}
	binary.BigEndian.PutUint32(buf[16:], uint32(msglen))
	copy(buf[20:], payload.Msg)
	copy(buf[20+len(payload.Msg):], payload.ReceiveID)

	return pkcs7.Pad(buf), nil
}

func (e *Encryptor) Encrypt(payload *Payload) (string, error) {
	buf, err := e.prepareBufForEncryption(payload)
	if err != nil {
		return "", err
	}

	// init cipher
	block, err := aes.NewCipher(e.aesKey)
	if err != nil {
		return "", err
	}

	iv := e.aesKey[:16]
	state := cipher.NewCBCEncrypter(block, iv)

	// encrypt in-place as we own the buffer
	state.CryptBlocks(buf, buf)

	return base64.StdEncoding.EncodeToString(buf), nil
}
