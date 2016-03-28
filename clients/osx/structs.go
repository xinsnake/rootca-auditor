package main

type OSXCert struct {
	Keychain   string
	Version    string
	Class      string
	Attributes struct {
		Alis string
		Cenc string
		Ctyp string
		Hpky string
		Issu string
		Labl string
		Skid string
		Snbr string
		Subj string
	}
}

type Cert struct {
	CertificateName string
	IssuedBy        string
	Type            string
	KeySize         string
	SigAlg          string
	SerialNumber    string
	Expires         string
	EVPolicy        string
}
