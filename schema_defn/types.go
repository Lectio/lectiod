package schema_defn

import (
	io "io"

	graphql "github.com/vektah/gqlgen/graphql"
)

type SmallText string
type MediumText string
type LargeText string
type ExtraLargeText string
type URLText string
type RegularExpression string
type ErrorMessage string
type ConfigurationName string

type AsymmetricCryptoPublicKey string
type IdentityPrincipal string
type IdentityPassword string
type IdentityKey string

type AuthenticatedSessionID string
type AuthenticatedSessionsCount uint
type AuthenticatedSessionTimeout uint

func (t SmallText) MarshalGQL(w io.Writer) {
	graphql.MarshalString(string(t)).MarshalGQL(w)
}

func (t MediumText) MarshalGQL(w io.Writer) {
	graphql.MarshalString(string(t)).MarshalGQL(w)
}

func (t LargeText) MarshalGQL(w io.Writer) {
	graphql.MarshalString(string(t)).MarshalGQL(w)
}

func (t ExtraLargeText) MarshalGQL(w io.Writer) {
	graphql.MarshalString(string(t)).MarshalGQL(w)
}

func (t URLText) MarshalGQL(w io.Writer) {
	graphql.MarshalString(string(t)).MarshalGQL(w)
}

func (t RegularExpression) MarshalGQL(w io.Writer) {
	graphql.MarshalString(string(t)).MarshalGQL(w)
}

func (t ErrorMessage) MarshalGQL(w io.Writer) {
	graphql.MarshalString(string(t)).MarshalGQL(w)
}

func (t ConfigurationName) MarshalGQL(w io.Writer) {
	graphql.MarshalString(string(t)).MarshalGQL(w)
}

func (t *ConfigurationName) UnmarshalGQL(v interface{}) error {
	str, err := graphql.UnmarshalString(v)
	if err == nil {
		*t = ConfigurationName(str)
	}
	return err
}

func (t IdentityPrincipal) MarshalGQL(w io.Writer) {
	graphql.MarshalString(string(t)).MarshalGQL(w)
}

func (t IdentityPassword) MarshalGQL(w io.Writer) {
	graphql.MarshalString(string(t)).MarshalGQL(w)
}

func (t IdentityKey) MarshalGQL(w io.Writer) {
	graphql.MarshalString(string(t)).MarshalGQL(w)
}

func (t AuthenticatedSessionID) MarshalGQL(w io.Writer) {
	graphql.MarshalString(string(t)).MarshalGQL(w)
}

func (t *AuthenticatedSessionID) UnmarshalGQL(v interface{}) error {
	str, err := graphql.UnmarshalString(v)
	if err == nil {
		*t = AuthenticatedSessionID(str)
	}
	return err
}

func (t AsymmetricCryptoPublicKey) MarshalGQL(w io.Writer) {
	graphql.MarshalString(string(t)).MarshalGQL(w)
}

func (t *AsymmetricCryptoPublicKey) UnmarshalGQL(v interface{}) error {
	str, err := graphql.UnmarshalString(v)
	if err == nil {
		*t = AsymmetricCryptoPublicKey(str)
	}
	return err
}

func (t AuthenticatedSessionTimeout) MarshalGQL(w io.Writer) {
	graphql.MarshalInt(int(t)).MarshalGQL(w)
}

func (t AuthenticatedSessionsCount) MarshalGQL(w io.Writer) {
	graphql.MarshalInt(int(t)).MarshalGQL(w)
}
