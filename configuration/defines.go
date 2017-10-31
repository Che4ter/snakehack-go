package configuration

//Head Types
const (
	BENDR = 1 + iota
	DEAD
	FANG
	PIXELHEAD
	REGULARHEAD
	SAFE
	SANDWORM
	SHADES
	SMILE
	TONGUE
)

var headtypes = [...]string{
	"bendr",
	"dead",
	"fang",
	"pixel",
	"regular",
	"safe",
	"sandworm",
	"shades",
	"smile",
	"tongue",
}

type HeadType int

// String() function will return the english name
// that we want out constant State be recognized as
func (headType HeadType) String() string {
	return headtypes[headType-1]
}

//Tail Types
const (
	SMALLRATTLE = 1 + iota
	ROUNDBUM
	REGULARTAIL
	PIXELTAIL
	FRECKLED
	FATRATTLE
	CURLED
	BLOCKBUM
)

var tailtypes = [...]string{
	"small-rattle",
	"round-bum",
	"regular",
	"pixel",
	"freckled",
	"fat-rattle",
	"curled",
	"block-bum",
}

type TailType int

// String() function will return the english name
// that we want out constant State be recognized as
func (tailType TailType) String() string {
	return tailtypes[tailType-1]
}

//Move Directions
const (
	MOVE_LEFT = 1 + iota
	MOVE_RIGHT
	MOVE_UP
	MOVE_DOWN
)

var movesType = [...]string{
	"left",
	"right",
	"up",
	"down",
}

type MoveDirection int

// String() function will return the english name
// that we want out constant State be recognized as
func (moveDirection MoveDirection) String() string {
	return movesType[moveDirection-1]
}
