package merchandise

type PrintArea int

const (
	FRONT PrintArea = iota + 1
	BACK
)

func (p PrintArea) String() string {
	v := [...]string{"FRONT", "BACK"}
	return v[p-1]
}
