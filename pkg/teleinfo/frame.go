package teleinfo

import "strconv"

const (
	TotalEnergy      = "EAST"
	InstantIntensity = "SINSTS"
	MaxIntensity     = "PCOUP"
)

type Frame struct {
	InstantIntensity int // SINSTS
	MaxIntensity     int // PCOUP
	TotalEnergy      int // EAST (kwh)
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

	if subscribedIntensity, ok := r.GetIntField(TotalEnergy); ok {
		f.TotalEnergy = subscribedIntensity
	}

	return f
}
