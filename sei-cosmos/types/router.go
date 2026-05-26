package types

import (
	"regexp"
)

var (
	// IsAlphaNumeric defines a regular expression for matching against alpha-numeric
	// values.
	IsAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

	// IsAlphaLower defines regular expression to check if the string has lowercase
	// alphabetic characters only.
	IsAlphaLower = regexp.MustCompile(`^[a-z]+$`).MatchString

	// IsAlphaUpper defines regular expression to check if the string has uppercase
	// alphabetic characters only.
	IsAlphaUpper = regexp.MustCompile(`^[A-Z]+$`).MatchString

	// IsAlpha defines regular expression to check if the string has alphabetic
	// characters only.
	IsAlpha = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	// IsNumeric defines regular expression to check if the string has numeric
	// characters only.
	IsNumeric = regexp.MustCompile(`^[0-9]+$`).MatchString
)

// Router provides handlers for each transaction type.
type Router interface {
	AddRoute(r Route) Router
	Route(ctx Context, path string) Handler
}

type Route struct {
	path    string
	handler Handler
}

// NewRoute returns an instance of Route.
func NewRoute(p string, h Handler) Route { _ = "STUB: not implemented"; return *new(Route) }

// Path returns the path the route has assigned.
func (r Route) Path() string {
	_ = "STUB: not implemented"

	// Handler returns the handler that handles the route.
	return ""
}

func (r Route) Handler() Handler {
	_ = "STUB: not implemented"

	// Empty returns true only if both handler and path are not empty.
	return *new(Handler)
}

func (r Route) Empty() bool { _ = "STUB: not implemented"; return false }

// QueryRouter provides queryables for each query path.
type QueryRouter interface {
	AddRoute(r string, h Querier) QueryRouter
	Route(path string) Querier
}
