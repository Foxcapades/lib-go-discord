package lib

import "fmt"

type AspectRatio struct { Width, Height uint16 }

func (a AspectRatio) String() string {
	return fmt.Sprintf("%d:%d", a.Width, a.Height)
}

func (a AspectRatio) FloatVal() float64 {
	return float64(a.Width) / float64(a.Height)
}
