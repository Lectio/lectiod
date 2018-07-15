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

type IdentityPrincipal string
type IdentityPassword string
type IdentityKey string

type AuthenticatedSessionID uint
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
	graphql.MarshalInt(int(t)).MarshalGQL(w)
}

func (t *AuthenticatedSessionID) UnmarshalGQL(v interface{}) error {
	num, err := graphql.UnmarshalInt(v)
	if err == nil {
		*t = AuthenticatedSessionID(num)
	}
	return err
}

func (t AuthenticatedSessionTimeout) MarshalGQL(w io.Writer) {
	graphql.MarshalInt(int(t)).MarshalGQL(w)
}
