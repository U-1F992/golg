package golg

type LifeGame struct {
	field [64][64]float32
	decay float32
}

func toBits(n uint64) [64]float32 {
	var bits [64]float32
	for i := 0; i < 64; i++ {
		if n&(1<<i) != 0 {
			bits[63-i] = 1
		}
	}
	// for i := 0; i < 64; i++ {
	// 	print(int(bits[i]))
	// }
	// print("\n")
	return bits
}

func NewLifeGame(seed [64]uint64, decay float32) *LifeGame {
	var f [64][64]float32
	for i, line := range seed {
		f[i] = toBits(line)
	}

	d := decay
	switch {
	case d < 0:
		d = 0
	case 1 < d:
		d = 1
	}

	return &LifeGame{
		field: f,
		decay: d,
	}
}

func isAlive(f *[64][64]float32, r int, c int) bool {
	if r < 0 || len(*f)-1 < r || c < 0 || len((*f)[r])-1 < c {
		return false
	}
	return f[r][c] == 1
}

func countTrue(a *[8]bool) int {
	count := 0
	for _, v := range *a {
		if v {
			count++
		}
	}
	return count
}

func countLivingNeighbours(f *[64][64]float32, r int, c int) int {
	return countTrue(&[8]bool{
		isAlive(f, r-1, c-1),
		isAlive(f, r-1, c),
		isAlive(f, r-1, c+1),
		isAlive(f, r, c-1),
		// isAlive(&f,r,c),
		isAlive(f, r, c+1),
		isAlive(f, r+1, c-1),
		isAlive(f, r+1, c),
		isAlive(f, r+1, c+1),
	})
}

func (lg *LifeGame) Next() {
	f := &lg.field
	next := lg.field

	for r := 0; r < len(f); r++ {
		for c := 0; c < len(f[r]); c++ {
			alive := f[r][c] == 1
			count := countLivingNeighbours(f, r, c)
			switch {
			case (!alive && count == 3) || (alive && (count == 2 || count == 3)):
				next[r][c] = 1
			default:
				//case (alive && count < 2) || (alive && 3 < count):
				val := f[r][c] * lg.decay
				if val < 0 {
					val = 0
				}
				next[r][c] = val
			}
		}
	}
	lg.field = next
}
