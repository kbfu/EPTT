package services

type Report struct {
	Average           int
	Median            int
	Min               int
	Max               int
	NinetyPercent     int
	NinetyFivePercent int
	NinetyNinePercent int
	Tps               map[int]int
}
