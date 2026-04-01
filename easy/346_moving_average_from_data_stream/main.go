package main

type MovingAverage struct {
	vals   []int
	length int
	size   int
}

func (this *MovingAverage) IsEmpty() bool {
	return this.length == 0
}
func (this *MovingAverage) IsFull() bool {
	return this.length == len(this.vals)
}
func (this *MovingAverage) Add(v int) {
	if !this.IsFull() {
		this.vals[this.length] = v
		this.length++
	} else {
		this.vals = append(this.vals[1:], v)
	}
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		vals:   make([]int, size),
		length: 0,
		size:   size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.Add(val)

	sum := 0
	for _, v := range this.vals {
		sum += v
	}

	return float64(sum) / float64(this.length)
}

func main() {

}
