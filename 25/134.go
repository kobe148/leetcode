package _25

func canCompleteCircuit(gas []int, cost []int) int {
	// totalGas 表示总油量
	// currGas 表示当前总油量
	totalGas, currGas, startStation := 0, 0, 0
	for i := range gas {
		totalGas += gas[i] - cost[i]
		currGas += gas[i] - cost[i]
		if currGas < 0 {
			startStation = i + 1
			currGas = 0
		}
	}

	if totalGas >= 0 {
		return startStation
	} else {
		return -1
	}
}
