package teleinfo

import (
	"fmt"
	"strconv"
	"time"
)

type RawFrame struct {
	TotalActiveEnergy            string `json:"EAST"`
	InstantApparentPower         string `json:"SINSTS"`
	InstantApparentInjectedPower string `json:"SINSTI"`
}

func (r *RawFrame) SetDefaults() {
	if r.TotalActiveEnergy == "" {
		r.TotalActiveEnergy = "0"
	}

	if r.InstantApparentPower == "" {
		r.InstantApparentPower = "0"
	}

	if r.InstantApparentInjectedPower == "" {
		r.InstantApparentInjectedPower = "0"
	}
}

func (r *RawFrame) ToFrame() (*Frame, error) {
	result := &Frame{}
	var err error
	result.TotalActiveEnergy, err = strconv.Atoi(r.TotalActiveEnergy)
	if err != nil {
		return nil, fmt.Errorf("can't cast TotalActiveEnergy to int: %w", err)
	}

	result.InstantApparentPower, err = strconv.Atoi(r.InstantApparentPower)
	if err != nil {
		return nil, fmt.Errorf("can't cast InstantApparentPower to int: %w", err)
	}

	result.InstantApparentInjectedPower, err = strconv.Atoi(r.InstantApparentInjectedPower)
	if err != nil {
		return nil, fmt.Errorf("can't cast InstantApparentInjectedPower to int: %w", err)
	}

	return result, nil
}

type Frame struct {
	Time time.Time `gorm:"time"`
	// TotalActiveEnergy in watt per hour
	TotalActiveEnergy int `gorm:"total_active_energy"`
	// InstantApparentPower in VA
	InstantApparentPower int `gorm:"instant_apparent_power"`
	// InstantApparentInjectedPower in VA
	InstantApparentInjectedPower int `gorm:"instant_apparent_injected_power"`
}

func (f Frame) TableName() string {
	return "teleinfo_frames"
}
