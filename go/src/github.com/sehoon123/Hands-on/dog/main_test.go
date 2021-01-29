package dog

import (
	"fmt"
	"testing"
)

func TestYearsTwo(t *testing.T) {
	v := YearsTwo(10)
	if v != 70 {
		fmt.Println("Got", v, "Wanted", "70")
	}
}

func TestYears(t *testing.T) {
	v := Years(10)
	if v != 70 {
		fmt.Println("Got", v, "Wanted", "70")
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	//Output:
	//70
}
func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	//Output:
	//70
}

func BenchmarkYearsTwo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}

}
func BenchmarkYears(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Years(10)
	}

}
