package elevator

var elevatorRunTestCases = []struct {
	route string
	floor int
}{
	{
		route: "(())",
		floor: 0,
	},
	{
		route: "()()",
		floor: 0,
	},
	{
		route: "(((",
		floor: 3,
	},
	{
		route: "(()(()(",
		floor: 3,
	},
	{
		route: "))(((((",
		floor: 3,
	},
	{
		route: "())",
		floor: -1,
	},
	{
		route: "))(",
		floor: -1,
	},
	{
		route: ")))",
		floor: -3,
	},
	{
		route: ")())())",
		floor: -3,
	},
}

var elevatorPositionTestCases = []struct {
	route    string
	floor    int
	position int
}{
	{
		route:    ")",
		floor:    -1,
		position: 1,
	},
	{
		route:    "()())",
		floor:    -1,
		position: 5,
	},
	{
		route:    "())())",
		floor:    -1,
		position: 3,
	},
}
