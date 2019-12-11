package color

import (
	"math/big"
	rand_math "math/rand"
	"time"
)

func ShuffleColors(arr []int) []int {
	rand_math.Seed(time.Now().UnixNano())
	rand_math.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })

	return arr
}

func PlanColorBit(c *big.Int, color int) *big.Int {
	if color == 0 {
		c = new(big.Int).And(c, new(big.Int).Not(new(big.Int).Lsh(big.NewInt(1), 0)))
		return new(big.Int).And(c, new(big.Int).Not(new(big.Int).Lsh(big.NewInt(1), 1)))
	}

	if color == 10 {
		c = new(big.Int).Or(c, new(big.Int).Lsh(big.NewInt(1), 1))
		return new(big.Int).And(c, new(big.Int).Not(new(big.Int).Lsh(big.NewInt(1), 0)))
	}

	c = new(big.Int).And(c, new(big.Int).Not(new(big.Int).Lsh(big.NewInt(1), 1)))

	return new(big.Int).Or(c, new(big.Int).Lsh(big.NewInt(1), 0))
}

func GetColorFromR(r *big.Int) int {
	checkTen := new(big.Int).And(r, new(big.Int).Lsh(big.NewInt(1), 1)).Uint64()
	checkOne := new(big.Int).And(r, new(big.Int).Lsh(big.NewInt(1), 0)).Uint64()

	if checkTen == 2 {
		return 10
	}

	if checkOne == 1 {
		return 1
	} else {
		return 0
	}
}
