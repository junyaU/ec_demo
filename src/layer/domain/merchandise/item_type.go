package merchandise

type ItemType int

const (
	T_SHIRT ItemType = iota + 1
	SWEAT_SHIRT
	HOODIE
	JACKET
	CAP
	BUCKET_HAT
	SANDALS
	PHONE_CASE
	TOTE_BAG
	// ...etc
)

func (t ItemType) String() string {
	v := [...]string{
		"T_SHIRT",
		"SWEAT_SHIRT",
		"HOODIE",
		"JACKET",
		"CAP",
		"BUCKET_HAT",
		"SANDALS",
		"PHONE_CASE",
		"TOTE_BAG",
	}
	return v[t-1]
}
