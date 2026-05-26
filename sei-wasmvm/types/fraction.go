package types

type Fraction struct {
	Numerator   int64
	Denominator int64
}

func (f *Fraction) Mul(m int64) Fraction { _ = "STUB: not implemented"; return *new(Fraction) }

func (f Fraction) Floor() int64 { _ = "STUB: not implemented"; return 0 }

type UFraction struct {
	Numerator   uint64
	Denominator uint64
}

func (f *UFraction) Mul(m uint64) UFraction { _ = "STUB: not implemented"; return *new(UFraction) }

func (f UFraction) Floor() uint64 { _ = "STUB: not implemented"; return 0 }
