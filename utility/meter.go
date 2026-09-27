package utility

import (
	"errors"
	"fmt"
	"os"
)

type Meter struct {
	MeterNumber     string
	ConsumerName    string
	PreviousReading float64
	CurrentReading  float64
}

var meters []Meter

func RegisterMeter(m Meter) error {

	if m.ConsumerName == "" {
		return errors.New("consumer name cannot be empty")
	}

	if m.PreviousReading < 0 {
		return errors.New("previous reading cannot be negative")
	}

	if m.CurrentReading < m.PreviousReading {
		return errors.New("current reading cannot be less than previous reading")
	}

	meters = append(meters, m)

	SaveMeters()

	return nil
}

func GetMeter(meterNum string) (*Meter, error) {

	for i := range meters {

		if meters[i].MeterNumber == meterNum {
			return &meters[i], nil
		}
	}

	return nil, errors.New("meter not found")
}

func RecordNewReading(meterNum string, reading float64) error {

	meter, err := GetMeter(meterNum)

	if err != nil {
		return err
	}

	if reading < meter.CurrentReading {
		return errors.New("new reading cannot be less than current reading")
	}

	meter.PreviousReading = meter.CurrentReading
	meter.CurrentReading = reading

	SaveMeters()

	return nil
}

func RemoveMeter(meterNum string) error {

	for i := range meters {

		if meters[i].MeterNumber == meterNum {

			meters = append(
				meters[:i],
				meters[i+1:]...,
			)

			SaveMeters()

			return nil
		}
	}

	return errors.New("meter not found")
}

func SaveMeters() {

	file, err := os.Create("meters.txt")

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	for _, meter := range meters {

		fmt.Fprintln(
			file,
			meter.MeterNumber,
			meter.ConsumerName,
			meter.PreviousReading,
			meter.CurrentReading,
		)
	}

	file.Close()
}

func LoadMeters() {

	file, err := os.Open("meters.txt")

	if err != nil {
		return
	}

	for {

		var meter Meter

		_, err := fmt.Fscan(
			file,
			&meter.MeterNumber,
			&meter.ConsumerName,
			&meter.PreviousReading,
			&meter.CurrentReading,
		)

		if err != nil {
			break
		}

		meters = append(meters, meter)
	}

	file.Close()
}
