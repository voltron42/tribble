package tribble

type Tribble interface {
	Byte() byte
	Float32() float32
	Float64() float64
	Int() int
	Int8() int8
	Int16() int16
	Int32() int32
	Int64() int64
	UInt() uint
	UInt16() uint16
	UInt32() uint32
	UInt64() uint64
	UIntPtr() uintptr
}
type byteTribble byte

func (b *byteTribble) Byte() byte       { return byte(*b) }
func (b *byteTribble) Float32() float32 { return float32(byte(*b)) }
func (b *byteTribble) Float64() float64 { return float64(byte(*b)) }
func (b *byteTribble) Int() int         { return int(byte(*b)) }
func (b *byteTribble) Int8() int8       { return int8(byte(*b)) }
func (b *byteTribble) Int16() int16     { return int16(byte(*b)) }
func (b *byteTribble) Int32() int32     { return int32(byte(*b)) }
func (b *byteTribble) Int64() int64     { return int64(byte(*b)) }
func (b *byteTribble) UInt() uint       { return uint(byte(*b)) }
func (b *byteTribble) UInt16() uint16   { return uint16(byte(*b)) }
func (b *byteTribble) UInt32() uint32   { return uint32(byte(*b)) }
func (b *byteTribble) UInt64() uint64   { return uint64(byte(*b)) }
func (b *byteTribble) UIntPtr() uintptr { return uintptr(byte(*b)) }

type float32Tribble float32

func (f *float32Tribble) Byte() byte       { return byte(float32(*f)) }
func (f *float32Tribble) Float32() float32 { return float32(*f) }
func (f *float32Tribble) Float64() float64 { return float64(float32(*f)) }
func (f *float32Tribble) Int() int         { return int(float32(*f)) }
func (f *float32Tribble) Int8() int8       { return int8(float32(*f)) }
func (f *float32Tribble) Int16() int16     { return int16(float32(*f)) }
func (f *float32Tribble) Int32() int32     { return int32(float32(*f)) }
func (f *float32Tribble) Int64() int64     { return int64(float32(*f)) }
func (f *float32Tribble) UInt() uint       { return uint(float32(*f)) }
func (f *float32Tribble) UInt16() uint16   { return uint16(float32(*f)) }
func (f *float32Tribble) UInt32() uint32   { return uint32(float32(*f)) }
func (f *float32Tribble) UInt64() uint64   { return uint64(float32(*f)) }
func (f *float32Tribble) UIntPtr() uintptr { return uintptr(float32(*f)) }

type float64Tribble float64

func (f *float64Tribble) Byte() byte       { return byte(float64(*f)) }
func (f *float64Tribble) Float32() float32 { return float32(float64(*f)) }
func (f *float64Tribble) Float64() float64 { return float64(*f) }
func (f *float64Tribble) Int() int         { return int(float64(*f)) }
func (f *float64Tribble) Int8() int8       { return int8(float64(*f)) }
func (f *float64Tribble) Int16() int16     { return int16(float64(*f)) }
func (f *float64Tribble) Int32() int32     { return int32(float64(*f)) }
func (f *float64Tribble) Int64() int64     { return int64(float64(*f)) }
func (f *float64Tribble) UInt() uint       { return uint(float64(*f)) }
func (f *float64Tribble) UInt16() uint16   { return uint16(float64(*f)) }
func (f *float64Tribble) UInt32() uint32   { return uint32(float64(*f)) }
func (f *float64Tribble) UInt64() uint64   { return uint64(float64(*f)) }
func (f *float64Tribble) UIntPtr() uintptr { return uintptr(float64(*f)) }

type intTribble int

func (i *intTribble) Byte() byte       { return byte(int(*i)) }
func (i *intTribble) Float32() float32 { return float32(int(*i)) }
func (i *intTribble) Float64() float64 { return float64(int(*i)) }
func (i *intTribble) Int() int         { return int(*i) }
func (i *intTribble) Int8() int8       { return int8(int(*i)) }
func (i *intTribble) Int16() int16     { return int16(int(*i)) }
func (i *intTribble) Int32() int32     { return int32(int(*i)) }
func (i *intTribble) Int64() int64     { return int64(int(*i)) }
func (i *intTribble) UInt() uint       { return uint(int(*i)) }
func (i *intTribble) UInt16() uint16   { return uint16(int(*i)) }
func (i *intTribble) UInt32() uint32   { return uint32(int(*i)) }
func (i *intTribble) UInt64() uint64   { return uint64(int(*i)) }
func (i *intTribble) UIntPtr() uintptr { return uintptr(int(*i)) }

type int8Tribble int8

func (i *int8Tribble) Byte() byte       { return byte(int8(*i)) }
func (i *int8Tribble) Float32() float32 { return float32(int8(*i)) }
func (i *int8Tribble) Float64() float64 { return float64(int8(*i)) }
func (i *int8Tribble) Int() int         { return int(int8(*i)) }
func (i *int8Tribble) Int8() int8       { return int8(*i) }
func (i *int8Tribble) Int16() int16     { return int16(int8(*i)) }
func (i *int8Tribble) Int32() int32     { return int32(int8(*i)) }
func (i *int8Tribble) Int64() int64     { return int64(int8(*i)) }
func (i *int8Tribble) UInt() uint       { return uint(int8(*i)) }
func (i *int8Tribble) UInt16() uint16   { return uint16(int8(*i)) }
func (i *int8Tribble) UInt32() uint32   { return uint32(int8(*i)) }
func (i *int8Tribble) UInt64() uint64   { return uint64(int8(*i)) }
func (i *int8Tribble) UIntPtr() uintptr { return uintptr(int8(*i)) }

type int16Tribble int16

func (i *int16Tribble) Byte() byte       { return byte(int16(*i)) }
func (i *int16Tribble) Float32() float32 { return float32(int16(*i)) }
func (i *int16Tribble) Float64() float64 { return float64(int16(*i)) }
func (i *int16Tribble) Int() int         { return int(int16(*i)) }
func (i *int16Tribble) Int8() int8       { return int8(int16(*i)) }
func (i *int16Tribble) Int16() int16     { return int16(*i) }
func (i *int16Tribble) Int32() int32     { return int32(int16(*i)) }
func (i *int16Tribble) Int64() int64     { return int64(int16(*i)) }
func (i *int16Tribble) UInt() uint       { return uint(int16(*i)) }
func (i *int16Tribble) UInt16() uint16   { return uint16(int16(*i)) }
func (i *int16Tribble) UInt32() uint32   { return uint32(int16(*i)) }
func (i *int16Tribble) UInt64() uint64   { return uint64(int16(*i)) }
func (i *int16Tribble) UIntPtr() uintptr { return uintptr(int16(*i)) }

type int32Tribble int32

func (i *int32Tribble) Byte() byte       { return byte(int32(*i)) }
func (i *int32Tribble) Float32() float32 { return float32(int32(*i)) }
func (i *int32Tribble) Float64() float64 { return float64(int32(*i)) }
func (i *int32Tribble) Int() int         { return int(int32(*i)) }
func (i *int32Tribble) Int8() int8       { return int8(int32(*i)) }
func (i *int32Tribble) Int16() int16     { return int16(int32(*i)) }
func (i *int32Tribble) Int32() int32     { return int32(*i) }
func (i *int32Tribble) Int64() int64     { return int64(int32(*i)) }
func (i *int32Tribble) UInt() uint       { return uint(int32(*i)) }
func (i *int32Tribble) UInt16() uint16   { return uint16(int32(*i)) }
func (i *int32Tribble) UInt32() uint32   { return uint32(int32(*i)) }
func (i *int32Tribble) UInt64() uint64   { return uint64(int32(*i)) }
func (i *int32Tribble) UIntPtr() uintptr { return uintptr(int32(*i)) }

type int64Tribble int64

func (i *int64Tribble) Byte() byte       { return byte(int64(*i)) }
func (i *int64Tribble) Float32() float32 { return float32(int64(*i)) }
func (i *int64Tribble) Float64() float64 { return float64(int64(*i)) }
func (i *int64Tribble) Int() int         { return int(int64(*i)) }
func (i *int64Tribble) Int8() int8       { return int8(int64(*i)) }
func (i *int64Tribble) Int16() int16     { return int16(int64(*i)) }
func (i *int64Tribble) Int32() int32     { return int32(int64(*i)) }
func (i *int64Tribble) Int64() int64     { return int64(*i) }
func (i *int64Tribble) UInt() uint       { return uint(int64(*i)) }
func (i *int64Tribble) UInt16() uint16   { return uint16(int64(*i)) }
func (i *int64Tribble) UInt32() uint32   { return uint32(int64(*i)) }
func (i *int64Tribble) UInt64() uint64   { return uint64(int64(*i)) }
func (i *int64Tribble) UIntPtr() uintptr { return uintptr(int64(*i)) }

type uintTribble uint

func (u *uintTribble) Byte() byte       { return byte(uint(*u)) }
func (u *uintTribble) Float32() float32 { return float32(uint(*u)) }
func (u *uintTribble) Float64() float64 { return float64(uint(*u)) }
func (u *uintTribble) Int() int         { return int(uint(*u)) }
func (u *uintTribble) Int8() int8       { return int8(uint(*u)) }
func (u *uintTribble) Int16() int16     { return int16(uint(*u)) }
func (u *uintTribble) Int32() int32     { return int32(uint(*u)) }
func (u *uintTribble) Int64() int64     { return int64(uint(*u)) }
func (u *uintTribble) UInt() uint       { return uint(*u) }
func (u *uintTribble) UInt16() uint16   { return uint16(uint(*u)) }
func (u *uintTribble) UInt32() uint32   { return uint32(uint(*u)) }
func (u *uintTribble) UInt64() uint64   { return uint64(uint(*u)) }
func (u *uintTribble) UIntPtr() uintptr { return uintptr(uint(*u)) }

type uint16Tribble uint16

func (u *uint16Tribble) Byte() byte       { return byte(uint16(*u)) }
func (u *uint16Tribble) Float32() float32 { return float32(uint16(*u)) }
func (u *uint16Tribble) Float64() float64 { return float64(uint16(*u)) }
func (u *uint16Tribble) Int() int         { return int(uint16(*u)) }
func (u *uint16Tribble) Int8() int8       { return int8(uint16(*u)) }
func (u *uint16Tribble) Int16() int16     { return int16(uint16(*u)) }
func (u *uint16Tribble) Int32() int32     { return int32(uint16(*u)) }
func (u *uint16Tribble) Int64() int64     { return int64(uint16(*u)) }
func (u *uint16Tribble) UInt() uint       { return uint(uint16(*u)) }
func (u *uint16Tribble) UInt16() uint16   { return uint16(*u) }
func (u *uint16Tribble) UInt32() uint32   { return uint32(uint16(*u)) }
func (u *uint16Tribble) UInt64() uint64   { return uint64(uint16(*u)) }
func (u *uint16Tribble) UIntPtr() uintptr { return uintptr(uint16(*u)) }

type uint32Tribble uint32

func (u *uint32Tribble) Byte() byte       { return byte(uint32(*u)) }
func (u *uint32Tribble) Float32() float32 { return float32(uint32(*u)) }
func (u *uint32Tribble) Float64() float64 { return float64(uint32(*u)) }
func (u *uint32Tribble) Int() int         { return int(uint32(*u)) }
func (u *uint32Tribble) Int8() int8       { return int8(uint32(*u)) }
func (u *uint32Tribble) Int16() int16     { return int16(uint32(*u)) }
func (u *uint32Tribble) Int32() int32     { return int32(uint32(*u)) }
func (u *uint32Tribble) Int64() int64     { return int64(uint32(*u)) }
func (u *uint32Tribble) UInt() uint       { return uint(uint32(*u)) }
func (u *uint32Tribble) UInt16() uint16   { return uint16(uint32(*u)) }
func (u *uint32Tribble) UInt32() uint32   { return uint32(*u) }
func (u *uint32Tribble) UInt64() uint64   { return uint64(uint32(*u)) }
func (u *uint32Tribble) UIntPtr() uintptr { return uintptr(uint32(*u)) }

type uint64Tribble uint64

func (u *uint64Tribble) Byte() byte       { return byte(uint64(*u)) }
func (u *uint64Tribble) Float32() float32 { return float32(uint64(*u)) }
func (u *uint64Tribble) Float64() float64 { return float64(uint64(*u)) }
func (u *uint64Tribble) Int() int         { return int(uint64(*u)) }
func (u *uint64Tribble) Int8() int8       { return int8(uint64(*u)) }
func (u *uint64Tribble) Int16() int16     { return int16(uint64(*u)) }
func (u *uint64Tribble) Int32() int32     { return int32(uint64(*u)) }
func (u *uint64Tribble) Int64() int64     { return int64(uint64(*u)) }
func (u *uint64Tribble) UInt() uint       { return uint(uint64(*u)) }
func (u *uint64Tribble) UInt16() uint16   { return uint16(uint64(*u)) }
func (u *uint64Tribble) UInt32() uint32   { return uint32(uint64(*u)) }
func (u *uint64Tribble) UInt64() uint64   { return uint64(*u) }
func (u *uint64Tribble) UIntPtr() uintptr { return uintptr(uint64(*u)) }

type uintptrTribble uintptr

func (u *uintptrTribble) Byte() byte       { return byte(uintptr(*u)) }
func (u *uintptrTribble) Float32() float32 { return float32(uintptr(*u)) }
func (u *uintptrTribble) Float64() float64 { return float64(uintptr(*u)) }
func (u *uintptrTribble) Int() int         { return int(uintptr(*u)) }
func (u *uintptrTribble) Int8() int8       { return int8(uintptr(*u)) }
func (u *uintptrTribble) Int16() int16     { return int16(uintptr(*u)) }
func (u *uintptrTribble) Int32() int32     { return int32(uintptr(*u)) }
func (u *uintptrTribble) Int64() int64     { return int64(uintptr(*u)) }
func (u *uintptrTribble) UInt() uint       { return uint(uintptr(*u)) }
func (u *uintptrTribble) UInt16() uint16   { return uint16(uintptr(*u)) }
func (u *uintptrTribble) UInt32() uint32   { return uint32(uintptr(*u)) }
func (u *uintptrTribble) UInt64() uint64   { return uint64(uintptr(*u)) }
func (u *uintptrTribble) UIntPtr() uintptr { return uintptr(*u) }

func NewTribble(value interface{}) Tribble {
	switch value.(type) {
	case byte:
		b, ok := value.(byte)
		if !ok {
			panic("not a byte")
		}
		out := byteTribble(b)
		return &out
	case float32:
		f, ok := value.(float32)
		if !ok {
			panic("not a float32")
		}
		out := float32Tribble(f)
		return &out
	case float64:
		f, ok := value.(float64)
		if !ok {
			panic("not a float64")
		}
		out := float64Tribble(f)
		return &out
	case int:
		i, ok := value.(int)
		if !ok {
			panic("not a int")
		}
		out := intTribble(i)
		return &out
	case int8:
		i, ok := value.(int8)
		if !ok {
			panic("not a int8")
		}
		out := int8Tribble(i)
		return &out
	case int16:
		i, ok := value.(int16)
		if !ok {
			panic("not a int16")
		}
		out := int16Tribble(i)
		return &out
	case int32:
		i, ok := value.(int32)
		if !ok {
			panic("not a int32")
		}
		out := int32Tribble(i)
		return &out
	case int64:
		i, ok := value.(int64)
		if !ok {
			panic("not a int64")
		}
		out := int64Tribble(i)
		return &out
	case uint:
		u, ok := value.(uint)
		if !ok {
			panic("not a uint")
		}
		out := uintTribble(u)
		return &out
	case uint16:
		u, ok := value.(uint16)
		if !ok {
			panic("not a uint16")
		}
		out := uint16Tribble(u)
		return &out
	case uint32:
		u, ok := value.(uint32)
		if !ok {
			panic("not a uint32")
		}
		out := uint32Tribble(u)
		return &out
	case uint64:
		u, ok := value.(uint64)
		if !ok {
			panic("not a uint64")
		}
		out := uint64Tribble(u)
		return &out
	case uintptr:
		u, ok := value.(uintptr)
		if !ok {
			panic("not a uintptr")
		}
		out := uintptrTribble(u)
		return &out
	default:
		panic("not a number type")
	}

}
