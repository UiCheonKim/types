package types

import "go-verifiable"

type VerifiableCredential struct {
	Verifiable *verifiable.Credential `json:"verifiable"`
	KeyName    string                 `json:"keyName"`
}
