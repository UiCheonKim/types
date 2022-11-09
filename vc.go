package types

import ""github.com/UiCheonKim/go-verifiable"

type VerifiableCredential struct {
	Verifiable *verifiable.Credential `json:"verifiable"`
	KeyName    string                 `json:"keyName"`
}
