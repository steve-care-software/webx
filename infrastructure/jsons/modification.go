package jsons

import "time"

// Modification represents a modification object
type Modification struct {
	Name      string    `json:"name"`
	SigPK     string    `json:"sig_pk"`
	EncPK     []byte    `json:"enc_pk"`
	CreatedOn time.Time `json:"created_on"`
}
