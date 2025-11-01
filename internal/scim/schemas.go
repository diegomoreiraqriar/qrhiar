package scim

const (
	SchemaUserCore = "urn:ietf:params:scim:schemas:core:2.0:User"
	SchemaQriar    = "urn:ietf:params:scim:schemas:extension:qriar:1.0:ThirdParty"
)

type ScimUser struct {
	Schemas []string      `json:"schemas"`
	ID      string        `json:"id,omitempty"`
	UserName string       `json:"userName"`
	Name    *struct {
		GivenName  string `json:"givenName,omitempty"`
		FamilyName string `json:"familyName,omitempty"`
	} `json:"name,omitempty"`
	Active  bool          `json:"active"`
	Emails  []struct {
		Value string `json:"value"`
	} `json:"emails,omitempty"`
	Ext     map[string]any `json:"urn:ietf:params:scim:schemas:extension:qriar:1.0:ThirdParty,omitempty"`
}
