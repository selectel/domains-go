package record

import "encoding/json"

// Type represents custom type for various records' statuses.
type Type string

const (
	TypeA       Type = "A"
	TypeAAAA    Type = "AAAA"
	TypeTXT     Type = "TXT"
	TypeCNAME   Type = "CNAME"
	TypeNS      Type = "NS"
	TypeSOA     Type = "SOA"
	TypeMX      Type = "MX"
	TypeSRV     Type = "SRV"
	TypeCAA     Type = "CAA"
	TypeSSHFP   Type = "SSHFP"
	TypeALIAS   Type = "ALIAS"
	TypeUnknown Type = "UNKNOWN"
)

// View represents an unmarshalled domain record body from API response.
type View struct {
	// ID is the identifier of the record.
	ID int `json:"id"`

	// Name represents record name.
	Name string `json:"name"`

	// Type represents record's type.
	Type Type `json:"-"`

	// TTL represents record's time-to-live.
	TTL int `json:"ttl,omitempty"`

	// Content represents record content.
	// Absent for SRV.
	Content string `json:"content,omitempty"`

	// ChangeDate represents Unix timestamp when record has been modified.
	// For SOA records only.
	ChangeDate *int `json:"change_date,omitempty"`

	// Emails represents email of domain's admin.
	// For SOA records only.
	Email string `json:"email,omitempty"`

	// Priority represents records preferences.
	// Lower value means more preferred.
	// For MX/SRV records only.
	Priority *int `json:"priority,omitempty"`

	// Weight represents a relative weight for records with the same priority,
	// higher value means higher chance of getting picked.
	// For SRV records only.
	Weight *int `json:"weight,omitempty"`

	// Port represents the TCP or UDP port on which the service is to be found.
	// For SRV records only.
	Port *int `json:"port,omitempty"`

	// Target represents the canonical hostname of the machine providing the service.
	// For SRV records only.
	Target string `json:"target,omitempty"`

	// Tag rrepresents the identifier of the property represented by the record.
	// For CAA records only.
	Tag string `json:"tag,omitempty"`

	// Flag represents the critical flag, that has a specific meaning per RFC.
	// For CAA records only.
	Flag *int `json:"flag,omitempty"`

	// The value associated with the tag.
	// For CAA records only.
	Value string `json:"value,omitempty"`

	// Algorithm.
	// For SSHFP records only.
	Algorithm *int `json:"algorithm,omitempty"`

	// Algorithm used to hash the public key
	// For SSHFP records only.
	FingerprintType *int `json:"fingerprint_type,omitempty"`

	// Hexadecimal representation of the hash result, as text.
	// For SSHFP records only.
	Fingerprint string `json:"fingerprint,omitempty"`
}

func (result *View) UnmarshalJSON(b []byte) error {
	type tmp View
	var v struct {
		tmp
		Type Type `json:"type"`
	}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	*result = View(v.tmp)

	// Check record types.
	switch v.Type {
	case TypeA:
		result.Type = TypeA
	case TypeAAAA:
		result.Type = TypeAAAA
	case TypeTXT:
		result.Type = TypeTXT
	case TypeCNAME:
		result.Type = TypeCNAME
	case TypeNS:
		result.Type = TypeNS
	case TypeMX:
		result.Type = TypeMX
	case TypeSRV:
		result.Type = TypeSRV
	case TypeSOA:
		result.Type = TypeSOA
	case TypeCAA:
		result.Type = TypeCAA
	case TypeALIAS:
		result.Type = TypeALIAS
	case TypeSSHFP:
		result.Type = TypeSSHFP
	default:
		result.Type = TypeUnknown
	}

	return err
}
