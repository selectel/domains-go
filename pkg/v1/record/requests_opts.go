package record

// CreateOpts represents requests options to create a domain record.
type CreateOpts struct {
	// Name represents record name.
	Name string `json:"name"`

	// Type represents record type.
	Type Type `json:"type"`

	// TTL represents record time-to-live.
	TTL int `json:"ttl"`

	// Content represents record content.
	// Absent for SRV.
	Content string `json:"content,omitempty"`

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

	// Tag represents the identifier of the property represented by the record.
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

// UpdateOpts represents requests options to update a domain record.
type UpdateOpts struct {
	// Name represents record name.
	Name string `json:"name"`

	// Type represents record type.
	Type Type `json:"type"`

	// TTL represents record time-to-live.
	TTL int `json:"ttl"`

	// Content represents record content.
	// Absent for SRV.
	Content string `json:"content,omitempty"`

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

	// Tag represents the identifier of the property represented by the record.
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
