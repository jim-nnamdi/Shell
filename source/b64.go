package source

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

type B64 interface {
	DecodeB64(b64 string) (int, error)
	DecodeToString(b64 string) (string, error)
	EncodeB64(b64 string) (string, error)
}

var _ B64 = &B64T{}

type B64T struct {
}

func NewB64T() *B64T {
	return &B64T{}
}

func (b64x *B64T) DecodeB64(b64 string) (int, []byte, error) {
	b64i := strings.TrimSuffix(b64, "\n")
	args := strings.Split(b64i, " ")

	switch args[0] {
	case "decode":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "%s", errors.New("wrong no of arguments"))
			return syscall.Stderr, nil, errors.New("invalid args")
		}
		writeToNewFile := os.WriteFile("decode.go", []byte(args[1]), 0777)
		if writeToNewFile != nil {
			fmt.Fprintf(os.Stderr, "%s", writeToNewFile)
			return syscall.Stderr, nil, writeToNewFile
		}
		readNFBytes, err := os.ReadFile("decode.go")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			return syscall.Stderr, nil, err
		}
		encodeB64 := base64.StdEncoding.EncodeToString(readNFBytes)
		decodeB64, err := base64.StdEncoding.DecodeString(encodeB64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			return syscall.Stderr, nil, err
		}
		return syscall.Stdout, decodeB64, err
	}
	cmd := exec.Command(b64i)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return syscall.Stderr, cmd.Run()
}

func (b64x *B64T) DecodeToString(b64 string) (string, error) {
	return "", nil
}
func (b64x *B64T) EncodeB64(b64 string) (string, error) {
	return "", nil
}
