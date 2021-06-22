package lc_6000

// 5370. 设计地铁系统

// 一把通过，牛逼

type UndergroundSystem struct {
	InEvent  map[string]map[int]int
	OutEvent map[string]map[int]int
}

func aConstructor() UndergroundSystem {
	return UndergroundSystem{
		InEvent:  make(map[string]map[int]int),
		OutEvent: make(map[string]map[int]int),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	if this.InEvent[stationName] == nil {
		this.InEvent[stationName] = map[int]int{}
	}
	this.InEvent[stationName][id] = t
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	if this.OutEvent[stationName] == nil {
		this.OutEvent[stationName] = map[int]int{}
	}
	this.OutEvent[stationName][id] = t
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {

	in := this.InEvent[startStation]
	out := this.OutEvent[endStation]

	time := 0
	count := 0
	for k, v := range out {
		if t, ok := in[k]; ok {
			time += (v - t)
			count++
		}
	}
	// 除数为0的问题怎么解决呢？题目没给出不存在时返回啥，就这样了
	return float64(time) / float64(count)

}
