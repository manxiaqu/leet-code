package leetcode

type CarType int

const (
	_ CarType = iota
	Big
	Medium
	Small
)

type ParkingSystem struct {
	big, medium, small int
}

func Constructor6(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big:    big,
		medium: medium,
		small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch CarType(carType) {
	case Big:
		if this.big == 0 {
			return false
		}
		this.big--
		return true
	case Medium:
		if this.medium == 0 {
			return false
		}
		this.medium--
		return true
	case Small:
		if this.small == 0 {
			return false
		}
		this.small--
		return true
	default:
		panic("invalid type")
	}
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
