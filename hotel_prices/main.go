package main

import (
	"context"
	"time"
)

// TODO: найти все цены для отелей из  за callTimeout
// Результат: список названий отелей с указанием цен от разных поставщиков и средним значением цены, например
// marriott [42, 34, 98] 58.0
// hilton [10, 20], 15.0

const callTimeout = 3 * time.Second

type HotelMatching struct {
	Hotel    string
	Supplier string
}

func main() {
	md := []HotelMatching{
		{
			Hotel:    "marriott",
			Supplier: "Supplier1",
		},
		{
			Hotel:    "hilton",
			Supplier: "Supplier1",
		},
		{
			Hotel:    "hilton",
			Supplier: "Supplier2",
		},
		{
			Hotel:    "holiday_inn",
			Supplier: "Supplier3",
		},
	}
}

type SearchResult struct {
	Hotel string
	Price float64
}

func SearchRPC(ctx context.Context, supplier string, hotels []string) ([]SearchResult, error) {
	// 0-2 seconds delay

	return []SearchResult{
		{Hotel: hotels[0], Price: 1200},
		{Hotel: hotels[1], Price: 1100},
		{Hotel: hotels[0], Price: 2300},
	}, nil
}

func ProcessAvailabilityForHotel(av []SearchResult) ([]float64, float64) {
	// calculation takes 1-2 second
	res := make([]float64, len(av))
	sum := 0.0
	for i := range av {
		res[i] = av[i].Price + 100
		sum += res[i]
	}

	return res, sum / float64(len(res))
}
