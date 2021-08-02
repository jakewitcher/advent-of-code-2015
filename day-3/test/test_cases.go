package test

var santaDeliveryTestCases = []struct {
	route  string
	houses int
}{
	{
		route:  ">",
		houses: 2,
	},
	{
		route:  "^>v<",
		houses: 4,
	},
	{
		route:  "^v^v^v^v^v",
		houses: 2,
	},
}

var robotSantaDeliveryTestCases = []struct {
	route  string
	houses int
}{
	{
		route:  "^v",
		houses: 3,
	},
	{
		route:  "^>v<",
		houses: 3,
	},
	{
		route:  "^v^v^v^v^v",
		houses: 11,
	},
}
