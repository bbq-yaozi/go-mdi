package test

import (
	"fmt"
	"testing"

	"github.com/bbq-yaozi/go-mdi/mdi"
)

func TestMdiConstant(t *testing.T) {
	t.Run("ab-testing1", func(t *testing.T) {
		a := mdi.AbTesting.Name()
		fmt.Println(a)
		a = "aaaaaa"
		fmt.Println(a)
		fmt.Println(mdi.AbTesting.Name())
	})

	t.Run("ab-testing2", func(t *testing.T) {
		a := mdi.AbTesting.Data()
		fmt.Println(a[0])
		a[0] = 88
		fmt.Println(a[0])
		fmt.Println(mdi.AbTesting.Data()[0])
	})

	t.Run("ab-testing3", func(t *testing.T) {
		a := mdi.Name("ab-testing").Data()
		fmt.Println(a[0])
		a[0] = 88
		fmt.Println(a[0])
		fmt.Println(mdi.AbTesting.Data()[0])
	})

	t.Run("ab-testing4", func(t *testing.T) {
		a := mdi.Icons()[0].Data()
		fmt.Println(a[0])
		a[0] = 88
		fmt.Println(a[0])
		fmt.Println(mdi.AbTesting.Data()[0])
	})
}
