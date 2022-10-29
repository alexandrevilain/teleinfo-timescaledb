package teleinfo

import "strconv"

const (
	SubscribedIntensity = "ISOUSC"
	InstantIntensity    = "IINST"
	MaxIntensity        = "IMAX"
)

type Frame struct {
	SubscribedIntensity int // ISOUSC
	InstantIntensity    int // IINST
	MaxIntensity        int // IMAX
}

type RawFrame map[string]string

func (f RawFrame) GetStringField(name string) (string, bool) {
	v, ok := f[name]
	return v, ok
}

func (f RawFrame) GetIntField(name string) (int, bool) {
	s, ok := f[name]
	if !ok {
		return 0, ok
	}
	num, err := strconv.ParseInt(s, 10, 32)
	ok = err == nil
	return int(num), ok
}

func NewFrameFromRaw(r *RawFrame) *Frame {
	f := &Frame{}

	if instantIntensity, ok := r.GetIntField(InstantIntensity); ok {
		f.InstantIntensity = instantIntensity
	}

	if maxIntensity, ok := r.GetIntField(MaxIntensity); ok {
		f.MaxIntensity = maxIntensity
	}

	if subscribedIntensity, ok := r.GetIntField(SubscribedIntensity); ok {
		f.SubscribedIntensity = subscribedIntensity
	}

	return f
}
