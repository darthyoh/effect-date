package effectdate

import (
	"time"
)

var frenchHolidays []string = []string{
	"2021-01-01", "2021-04-05", "2021-05-01", "2021-05-08", "2021-05-13", "2021-05-24", "2021-07-14", "2021-08-15", "2021-11-01", "2021-11-11", "2021-12-25",
	"2022-01-01", "2022-04-18", "2022-05-01", "2022-05-08", "2022-05-26", "2022-06-06", "2022-07-14", "2022-08-15", "2022-11-01", "2022-11-11", "2022-12-25",
	"2023-01-01", "2023-04-10", "2023-05-01", "2023-05-08", "2023-05-18", "2023-05-29", "2023-07-14", "2023-08-15", "2023-11-01", "2023-11-11", "2023-12-25",
	"2024-01-01", "2024-04-01", "2024-05-01", "2024-05-08", "2024-05-09", "2024-05-20", "2024-07-14", "2024-08-15", "2024-11-01", "2024-11-11", "2024-12-25",
	"2025-01-01", "2025-04-25", "2025-05-01", "2025-05-08", "2025-05-29", "2025-06-09", "2025-07-14", "2025-08-15", "2025-11-01", "2025-11-11", "2025-12-25",
	"2026-01-01", "2026-04-06", "2026-05-01", "2026-05-08", "2026-05-14", "2026-05-25", "2026-07-14", "2026-08-15", "2026-11-01", "2026-11-11", "2026-12-25",
	"2027-01-01", "2027-03-29", "2027-05-01", "2027-05-06", "2027-05-08", "2027-05-17", "2027-07-14", "2027-08-15", "2027-11-01", "2027-11-11", "2027-12-25",
}

var holidaysMap map[time.Time]bool

func init() {
	holidaysMap = make(map[time.Time]bool)
	for _, day := range frenchHolidays {
		day, _ := time.Parse("2006-01-02", day)
		holidaysMap[day] = true
	}
}

func GetEffectDate(isoDate string, delay int, openDays bool) (string, error) {

	inputDate, err := time.Parse("2006-01-02", isoDate)
	if err != nil {
		return "", err
	}

	if delay == 0 {
		return isoDate, nil
	}

	forward := func() {
		delay--
	}
	increaseTime := func() {
		inputDate = inputDate.Add(time.Hour * 24)
	}

	if delay < 0 {
		forward = func() {
			delay++
		}
		increaseTime = func() {
			inputDate = inputDate.Add(-time.Hour * 24)
		}
	}

	for delay != 0 {
		increaseTime()

		if _, ok := holidaysMap[inputDate]; !openDays || (int(inputDate.Weekday()) != 0 && int(inputDate.Weekday()) != 6 && !ok) {
			forward()
		}

	}

	return inputDate.Format("2006-01-02"), nil

}
