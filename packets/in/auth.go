package in

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"gpko/packets/utils"
)

type Auth struct {
	KeyLen        uint16
	Key           []byte
	Login         string
	PasswordLen   uint16
	Password      []byte
	MAC           string
	IsCheat       uint16
	ClientVersion uint16
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) Read(in []byte) error {
	buf := bytes.NewBuffer(in)

	err := utils.ReadPKO(buf, binary.BigEndian, &a.KeyLen)
	if err != nil {
		return err
	}
	a.Key = buf.Next(int(a.KeyLen))

	err = utils.ReadPKO(buf, binary.BigEndian, &a.Login)
	if err != nil {
		return err
	}

	err = utils.ReadPKO(buf, binary.BigEndian, &a.PasswordLen)
	if err != nil {
		return err
	}
	a.Password = buf.Next(int(a.PasswordLen))

	err = utils.ReadPKO(buf, binary.BigEndian, &a.MAC)
	if err != nil {
		return err
	}

	err = utils.ReadPKO(buf, binary.BigEndian, &a.IsCheat)
	if err != nil {
		return err
	}

	err = utils.ReadPKO(buf, binary.BigEndian, &a.ClientVersion)
	if err != nil {
		return err
	}

	return nil
}

func (a *Auth) Handle() {
	fmt.Println(a)
}
