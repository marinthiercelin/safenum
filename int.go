package safenum

// Int represents a signed integer of arbitrary size.
//
// Similarly to Nat, each Int comes along with an announced size, representing
// the number of bits need to represent its absolute value. This can be
// larger than its true size, the number of bits actually needed.
type Int struct {
	// This number is represented by (-1)^sign * abs, essentially

	// When 1, this is a negative number, when 0 a positive number.
	//
	// There's a bit of redundancy to note, because -0 and +0 represent the same
	// number. We need to be careful around this edge case.
	sign Choice
	// The absolute value.
	//
	// Not using a point is important, that way the zero value for Int is actually zero.
	abs Nat
}

// SetBytes interprets a number in big-endian form, stores it in z, and returns z.
//
// This number will be positive.
func (z *Int) SetBytes(data []byte) *Int {
	z.sign = 0
	z.abs.SetBytes(data)
	return z
}

// Eq checks if this Int has the same value as another Int.
//
// Note that negative zero and positive zero are the same number.
func (z *Int) Eq(x *Int) Choice {
	zero := z.abs.EqZero()
	// If this is zero, then any number as the same sign,
	// otherwise, check that the signs aren't different
	sameSign := zero | (1 ^ z.sign ^ x.sign)
	return sameSign & z.abs.Eq(&x.abs)
}

// Neg calculates z <- -z.
//
// The result has the same announced size.
func Neg(z *Int) *Int {
	z.sign ^= 1
	return z
}