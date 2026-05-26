package types

// denomUnits contains a mapping of denomination mapped to their respective unit
// multipliers (e.g. 1atom = 10^-6uatom).
var denomUnits = map[string]Dec{}

// baseDenom is the denom of smallest unit registered
var baseDenom string

// RegisterDenom registers a denomination with a corresponding unit. If the
// denomination is already registered, an error will be returned.
func RegisterDenom(denom string, unit Dec) error { _ = "STUB: not implemented"; return nil }

// GetDenomUnit returns a unit for a given denomination if it exists. A boolean
// is returned if the denomination is registered.
func GetDenomUnit(denom string) (Dec, bool) { _ = "STUB: not implemented"; return *new(Dec), false }

// GetBaseDenom returns the denom of smallest unit registered
func GetBaseDenom() (string, error) { _ = "STUB: not implemented"; return "", nil }

func MustGetBaseDenom() string { _ = "STUB: not implemented"; return "" }

// ConvertCoin attempts to convert a coin to a given denomination. If the given
// denomination is invalid or if neither denomination is registered, an error
// is returned.
func ConvertCoin(coin Coin, denom string) (Coin, error) {
	_ = "STUB: not implemented"
	return *new(Coin), nil
}

// ConvertDecCoin attempts to convert a decimal coin to a given denomination. If the given
// denomination is invalid or if neither denomination is registered, an error
// is returned.
func ConvertDecCoin(coin DecCoin, denom string) (DecCoin, error) {
	_ = "STUB: not implemented"
	return *new(DecCoin), nil
}

// NormalizeCoin try to convert a coin to the smallest unit registered,
// returns original one if failed.
func NormalizeCoin(coin Coin) Coin { _ = "STUB: not implemented"; return *new(Coin) }

// NormalizeDecCoin try to convert a decimal coin to the smallest unit registered,
// returns original one if failed.
func NormalizeDecCoin(coin DecCoin) DecCoin { _ = "STUB: not implemented"; return *new(DecCoin) }

// NormalizeCoins normalize and truncate a list of decimal coins
func NormalizeCoins(coins []DecCoin) Coins { _ = "STUB: not implemented"; return *new(Coins) }
