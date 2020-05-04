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
	TypeMX      Type = "MX"
	TypeSRV     Type = "SRV"
	TypeUnknown Type = "UNKNOWN"
)

// View represents an unmarshalled domain record body from API response.
type View struct {
	// ID is the identifier of the record.
	ID int `json:"id"`

	// CreateDate represents Unix timestamp when record has been created.
	CreateDate int `json:"create_date,omitempty"`

	// ChangeDate represents Unix timestamp when record has been modified.
	ChangeDate int `json:"change_date,omitempty"`

	// Name represents record name.
	Name string `json:"name"`

	// Type represents record's type.
	Type Type `json:"-"`

	// TTL represents record's time-to-live.
	TTL int `json:"ttl,omitempty"`

	// Emails represents email of domain's admin.
	// For SOA records only.
	Email string `json:"email,omitempty"`

	// Content represents record content
	// Absent for SRV.
	Content string `json:"content,omitempty"`
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
	default:
		result.Type = TypeUnknown
	}

	return err
}
