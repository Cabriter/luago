package state

type LuaType = int
type ArithOp = int
type CompareOp = int

type luaState struct {
	stack *luaStack
}

type LuaState interface {
	/* basic stack manipulation */
	GetTop() int
	AbsIndex(idx int) int
	CheckStack(n int) bool
	Pop(n int)
	Copy(fromIdx,toIdx int)
	PushValue(idx int)
	Replace(idx int)
	Insert(idx int)
	Remove(idx int)
	Rotate(idx,n int)
	SetTop(idx int)
	/* access functions (stack -> Go) */
	TypeName(tp LuaType) string
	Type(idx int) LuaType
	IsNone(idx int)bool
	IsNil(idx int)bool
	IsNoneOrNil(idx int)bool
	IsBoolean(idx int)bool
	IsInteger(idx int)bool
	IsNumber(idx int)bool
	IsString(idx int)bool
	ToBoolean(idx int) bool
	ToInteger(idx int)int64
	ToIntegerX(idx int)(int64,bool)
	ToNumber(idx int)float64
	ToNumberX(idx int)(float64,bool)
	ToString(idx int)string
	ToStringX(idx int)(string,bool)
	/* push function (Go -> stack) */
	PushNil()
	PushBoolean(b bool)
	PushInteger(n int64)
	PushNumber(n float64)
	PushString(s string)
	Arith(op ArithOp)
	Compare(idx1,idx2 int,op CompareOp)bool
	Len(idx int)
	Concat(n int)
}

func New() *luaState {
	return &luaState{stack: newLuaStack(20)}
}