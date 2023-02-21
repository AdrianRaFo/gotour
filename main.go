package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"strings"
	"time"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (math.Pow(z, 2)-x)/(2*z)
	}
	return z, nil
}

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		res[y] = make([]uint8, dx)
	}
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			res[y][x] = uint8((x + y) / 2)
		}
	}
	return res
}

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		res[word] += 1
	}
	return res
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	seq := make([]int, 0)
	return func() int {
		if len(seq) == 0 {
			seq = append(seq, 0)
			return 0
		}
		if len(seq) == 1 {
			seq = append(seq, 1)
			return 1
		}
		first, second := seq[len(seq)-2], seq[len(seq)-1]
		seq = append(seq, first+second)
		return first + second
	}
}

type MyReader struct{}

func (r MyReader) Read(buffer []byte) (int, error) {
	for i := 0; i < len(buffer); i++ {
		buffer[i] = 'A'
	}
	return len(buffer), nil
}

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		b = 'A' + (b-'A'+13)%26
	} else if b >= 'a' && b <= 'z' {
		b = 'a' + (b-'a'+13)%26
	}
	return b
}

func (r rot13Reader) Read(buffer []byte) (int, error) {
	read, err := r.r.Read(buffer)
	for i := 0; i < read; i++ {
		buffer[i] = rot13(buffer[i])
	}
	return read, err
}

type Image struct {
	width  int
	height int
	color  uint8
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.color + uint8(x/2), i.color + uint8(y/2), i.color + uint8((x+y)/2), 255}
}

type List[T any] struct {
	val  T
	next *List[T]
}

func count[T any](l List[T]) int {
	sum := 1
	for i := l.next; i != nil; i = i.next {
		sum += 1
	}
	return sum
}

func (l List[T]) toArray() []T {
	arr := []T{l.val}
	for i := l.next; i != nil; i = i.next {
		arr = append(arr, i.val)
	}
	return arr
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	//defer fmt.Println("world")
	//fmt.Println("By math", math.Sqrt(2))
	//fmt.Println("By function", Sqrt(2))
	//fmt.Println("hello")
	//fmt.Println("counting")
	//
	//for i := 0; i < 10; i++ {
	//	defer fmt.Println(i)
	//}
	//
	//fmt.Println("done")

	//pic.Show(Pic)
	//wc.Test(WordCount)
	//f := fibonacci()
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%d ", f())
	//}
	//fmt.Println(Sqrt(-2))
	//reader.Validate(MyReader{})
	//s := strings.NewReader("Lbh penpxrq gur pbqr!")
	//r := rot13Reader{s}
	//io.Copy(os.Stdout, &r)
	//m := Image{100, 100, 100}
	//pic.ShowImage(m)
	//intList := List[int]{1, &List[int]{2, &List[int]{3, nil}}}
	//fmt.Println(count(intList))
	//fmt.Println(intList.toArray())
	go say("world")
	say("hello")
}
