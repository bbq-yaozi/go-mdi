package test

import (
	"testing"

	"github.com/bbq-yaozi/go-mdi/internal/generate"
)

func TestGenerator(t *testing.T) {
	t.Run("LoadMetaList", func(t *testing.T) {
		metaList, err := generate.LoadMetaList()
		if err != nil {
			t.Error(err)
		}
		if len(metaList) == 0 {
			t.Errorf("svg not found")
		}
	})
}
