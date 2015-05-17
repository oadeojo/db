package sqlgen

import (
	"testing"
)

func TestOrderBy(t *testing.T) {
	o := NewOrderBy(
		NewSortColumns(
			SortColumn{Column: Column{Name: "foo"}},
		),
	)

	s := o.Compile(defaultTemplate)
	e := `ORDER BY "foo"`

	if trim(s) != e {
		t.Fatalf("Got: %s, Expecting: %s", s, e)
	}
}

func TestOrderByDesc(t *testing.T) {
	o := NewOrderBy(
		NewSortColumns(
			SortColumn{Column: Column{Name: "foo"}, Order: SqlOrderDesc},
		),
	)

	s := o.Compile(defaultTemplate)
	e := `ORDER BY "foo" DESC`

	if trim(s) != e {
		t.Fatalf("Got: %s, Expecting: %s", s, e)
	}
}

func BenchmarkOrderBy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewOrderBy(
			NewSortColumns(
				SortColumn{Column: Column{Name: "foo"}},
			),
		)
	}
}

func BenchmarkOrderByHash(b *testing.B) {
	o := OrderBy{
		SortColumns: NewSortColumns(
			SortColumn{Column: Column{Name: "foo"}},
		),
	}
	for i := 0; i < b.N; i++ {
		o.Hash()
	}
}

func BenchmarkCompileOrderByCompile(b *testing.B) {
	o := OrderBy{
		SortColumns: NewSortColumns(
			SortColumn{Column: Column{Name: "foo"}},
		),
	}
	for i := 0; i < b.N; i++ {
		o.Compile(defaultTemplate)
	}
}

func BenchmarkCompileOrderByCompileNoCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		o := NewOrderBy(
			NewSortColumns(
				SortColumn{Column: Column{Name: "foo"}},
			),
		)
		o.Compile(defaultTemplate)
	}
}

func BenchmarkCompileOrderCompile(b *testing.B) {
	o := SqlOrderDesc
	for i := 0; i < b.N; i++ {
		o.Compile(defaultTemplate)
	}
}

func BenchmarkCompileOrderCompileNoCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		o := SqlOrderDesc
		o.Compile(defaultTemplate)
	}
}

func BenchmarkSortColumnHash(b *testing.B) {
	s := SortColumn{Column: Column{Name: "foo"}}
	for i := 0; i < b.N; i++ {
		s.Hash()
	}
}

func BenchmarkSortColumnCompile(b *testing.B) {
	s := SortColumn{Column: Column{Name: "foo"}}
	for i := 0; i < b.N; i++ {
		s.Compile(defaultTemplate)
	}
}

func BenchmarkSortColumnCompileNoCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := SortColumn{Column: Column{Name: "foo"}}
		s.Compile(defaultTemplate)
	}
}

func BenchmarkSortColumnsHash(b *testing.B) {
	s := NewSortColumns(
		SortColumn{Column: Column{Name: "foo"}},
		SortColumn{Column: Column{Name: "bar"}},
	)
	for i := 0; i < b.N; i++ {
		s.Hash()
	}
}

func BenchmarkSortColumnsCompile(b *testing.B) {
	s := NewSortColumns(
		SortColumn{Column: Column{Name: "foo"}},
		SortColumn{Column: Column{Name: "bar"}},
	)
	for i := 0; i < b.N; i++ {
		s.Compile(defaultTemplate)
	}
}

func BenchmarkSortColumnsCompileNoCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := NewSortColumns(
			SortColumn{Column: Column{Name: "foo"}},
			SortColumn{Column: Column{Name: "bar"}},
		)
		s.Compile(defaultTemplate)
	}
}
