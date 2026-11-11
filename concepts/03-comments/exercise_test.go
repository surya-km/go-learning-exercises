package comments

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"strings"
	"testing"
)

// docComments parses exercise.go and returns the doc comment text for every
// exported top-level name (an empty string means it has no doc comment).
func docComments(t *testing.T) map[string]string {
	t.Helper()
	file, err := parser.ParseFile(token.NewFileSet(), "exercise.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatalf("could not parse exercise.go: %v", err)
	}

	docs := map[string]string{}
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Recv == nil && d.Name.IsExported() {
				docs[d.Name.Name] = d.Doc.Text()
			}
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				// A comment above an ungrouped declaration belongs to the GenDecl;
				// inside a ( ... ) group it belongs to each spec.
				var names []*ast.Ident
				var doc *ast.CommentGroup
				switch s := spec.(type) {
				case *ast.TypeSpec:
					names, doc = []*ast.Ident{s.Name}, s.Doc
				case *ast.ValueSpec:
					names, doc = s.Names, s.Doc
				}
				if doc == nil && len(d.Specs) == 1 {
					doc = d.Doc
				}
				for _, n := range names {
					if n.IsExported() {
						docs[n.Name] = doc.Text()
					}
				}
			}
		}
	}
	return docs
}

func TestDocComments(t *testing.T) {
	docs := docComments(t)

	for _, name := range []string{"CurrencySymbol", "Transaction", "FormatAmount", "Describe"} {
		if _, ok := docs[name]; !ok {
			docs[name] = "" // make sure we still check names that were deleted or renamed
		}
	}

	names := make([]string, 0, len(docs))
	for name := range docs {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		doc := docs[name]
		t.Run(name, func(t *testing.T) {
			if strings.TrimSpace(doc) == "" {
				t.Fatalf("%s has no doc comment; add one directly above it, starting with %q", name, name+" ")
			}
			if !strings.HasPrefix(doc, name+" ") {
				firstLine, _, _ := strings.Cut(doc, "\n")
				t.Errorf("doc comment for %s should start with %q, got %q", name, name+" ", firstLine)
			}
		})
	}
}

func TestFormatAmount(t *testing.T) {
	tests := []struct {
		cents int
		want  string
	}{
		{cents: 1234, want: "$12.34"},
		{cents: 5, want: "$0.05"},
		{cents: 50, want: "$0.50"},
		{cents: 0, want: "$0.00"},
		{cents: 100, want: "$1.00"},
		{cents: -350, want: "-$3.50"},
		{cents: -1, want: "-$0.01"},
		{cents: 123456789, want: "$1234567.89"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := FormatAmount(tt.cents); got != tt.want {
				t.Errorf("FormatAmount(%d) = %q, want %q", tt.cents, got, tt.want)
			}
		})
	}
}

func TestDescribe(t *testing.T) {
	tests := []struct {
		tx   Transaction
		want string
	}{
		{tx: Transaction{Description: "Coffee", AmountCents: 350}, want: "Coffee: $3.50"},
		{tx: Transaction{Description: "Salary", AmountCents: 300000}, want: "Salary: $3000.00"},
		{tx: Transaction{Description: "Refund", AmountCents: -1299}, want: "Refund: -$12.99"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := Describe(tt.tx); got != tt.want {
				t.Errorf("Describe(%+v) = %q, want %q", tt.tx, got, tt.want)
			}
		})
	}
}
