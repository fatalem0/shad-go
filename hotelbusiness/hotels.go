//go:build !solution

package hotelbusiness

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func ComputeLoad(guests []Guest) []Load {
	if len(guests) == 0 {
		return nil
	}

	maxDate := guests[0].CheckOutDate
	for _, guest := range guests {
		maxDate = max(maxDate, guest.CheckOutDate)
	}

	dateGuests := make([]int, maxDate+1)
	for _, v := range guests {
		dateGuests[v.CheckInDate]++
		dateGuests[v.CheckOutDate]--
	}

	res := make([]Load, 0)

	guestsCount := 0
	for date, guests := range dateGuests {
		guestsCount += guests

		if guests == 0 {
			continue
		}

		res = append(res, Load{StartDate: date, GuestCount: guestsCount})
	}

	return res
}
