// Command r5import copies the FHIR R5 models produced by golang-fhir-models
// (https://github.com/samply/golang-fhir-models) into models/r5/fhir500.
//
// It only ever adds. A generated file is copied when none of the types it
// declares exists in the package yet, so the hand-written models stay exactly
// as they are and callers keep compiling. To replace a hand-written model,
// delete its types from the package first and run the import again. A file
// that keeps other hand-written code must not share the generated file's name,
// or the copy is refused.
//
// Each copied file is adapted to the package's conventions: the package is
// renamed to fhir500, the bson struct tags are dropped, and the generated Id
// field is renamed to ID. The licence header and the generated-file notice are
// kept.
//
// Usage:
//
//	go run ./tools/r5import -src ~/golang-fhir-models/fhir-models/fhir
package main

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// skipped are generated files that cannot be used as they are.
//
// auditEventOutcome.go declares the audit-event-outcome code list as
// AuditEventOutcome, the name R5 also gives the AuditEvent.outcome backbone
// element in auditEvent.go. R5 binds that element to a Coding, so the code list
// is unused and the backbone element wins.
var skipped = map[string]bool{
	"root.go":              true,
	"auditEventOutcome.go": true,
}

var (
	bsonTag     = regexp.MustCompile(`bson:"[^"]*" `)
	idField     = regexp.MustCompile(`(?m)^(\s+)Id(\s+)`)
	idSelector  = regexp.MustCompile(`\.Id\b`)
	packageLine = regexp.MustCompile(`(?m)^package fhir$`)
)

type generated struct {
	name     string
	path     string
	declares []string
	uses     []string
}

func main() {
	src := flag.String("src", "", "directory holding the generated R5 Go files")
	dst := flag.String("dst", "models/r5/fhir500", "package to copy the models into")
	dryRun := flag.Bool("dry-run", false, "report what would be copied without writing")
	flag.Parse()

	if *src == "" {
		log.Fatal("r5import: -src is required")
	}

	if err := run(*src, *dst, *dryRun); err != nil {
		log.Fatal(err)
	}
}

func run(src, dst string, dryRun bool) error {
	existing, err := declaredIn(dst)
	if err != nil {
		return err
	}

	files, err := readGenerated(src)
	if err != nil {
		return err
	}

	var added, clashing []string

	for _, f := range files {
		if clash := intersect(f.declares, existing); len(clash) > 0 {
			clashing = append(clashing, fmt.Sprintf("%s (%s)", f.name, strings.Join(clash, ", ")))

			continue
		}

		added = append(added, f.name)

		if dryRun {
			continue
		}

		if err := copyFile(f, dst); err != nil {
			return err
		}
	}

	fmt.Printf("added %d files, kept %d hand-written\n", len(added), len(clashing))

	for _, c := range clashing {
		fmt.Println("  kept:", c)
	}

	return nil
}

// declaredIn returns the names of the types declared in a package directory,
// ignoring tests.
func declaredIn(dir string) (map[string]bool, error) {
	paths, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		return nil, fmt.Errorf("r5import: list %s: %w", dir, err)
	}

	names := map[string]bool{}

	for _, p := range paths {
		if strings.HasSuffix(p, "_test.go") {
			continue
		}

		file, err := parser.ParseFile(token.NewFileSet(), p, nil, parser.SkipObjectResolution)
		if err != nil {
			return nil, fmt.Errorf("r5import: parse %s: %w", p, err)
		}

		for _, n := range typeNames(file) {
			names[n] = true
		}
	}

	return names, nil
}

func readGenerated(dir string) ([]generated, error) {
	paths, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		return nil, fmt.Errorf("r5import: list %s: %w", dir, err)
	}

	if len(paths) == 0 {
		return nil, errors.New("r5import: no Go files in " + dir)
	}

	sort.Strings(paths)

	files := make([]generated, 0, len(paths))

	for _, p := range paths {
		name := filepath.Base(p)
		if skipped[name] {
			continue
		}

		file, err := parser.ParseFile(token.NewFileSet(), p, nil, parser.SkipObjectResolution)
		if err != nil {
			return nil, fmt.Errorf("r5import: parse %s: %w", p, err)
		}

		files = append(files, generated{name: name, path: p, declares: typeNames(file)})
	}

	return files, nil
}

func typeNames(file *ast.File) []string {
	var names []string

	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.TYPE {
			continue
		}

		for _, spec := range gen.Specs {
			if ts, ok := spec.(*ast.TypeSpec); ok {
				names = append(names, ts.Name.Name)
			}
		}
	}

	return names
}

func intersect(names []string, set map[string]bool) []string {
	var out []string

	for _, n := range names {
		if set[n] {
			out = append(out, n)
		}
	}

	return out
}

func copyFile(f generated, dst string) error {
	raw, err := os.ReadFile(f.path)
	if err != nil {
		return fmt.Errorf("r5import: read %s: %w", f.path, err)
	}

	out, err := format.Source(adapt(raw))
	if err != nil {
		return fmt.Errorf("r5import: format %s: %w", f.name, err)
	}

	// O_EXCL refuses to replace a file that is already there, even one whose
	// types do not clash, so a hand-written file is never overwritten.
	file, err := os.OpenFile(filepath.Join(dst, f.name), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return fmt.Errorf("r5import: create %s: %w", f.name, err)
	}

	if _, err := file.Write(out); err != nil {
		_ = file.Close()

		return fmt.Errorf("r5import: write %s: %w", f.name, err)
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("r5import: close %s: %w", f.name, err)
	}

	return nil
}

// adapt applies the package's conventions to one generated file.
func adapt(src []byte) []byte {
	s := packageLine.ReplaceAllString(string(src), "package fhir500")
	s = bsonTag.ReplaceAllString(s, "")
	s = idField.ReplaceAllString(s, "${1}ID${2}")
	s = idSelector.ReplaceAllString(s, ".ID")

	return []byte(s)
}
