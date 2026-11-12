package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// newAlias creates contracts/<alias> in a temporary directory, with one file in it, and returns its path
func newAlias(t *testing.T, content string) string {
	t.Helper()
	aliasPath := filepath.Join(t.TempDir(), "contracts", "alias")
	if err := os.MkdirAll(aliasPath, 0o755); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(aliasPath, "previous.go"), content)
	return aliasPath
}

func writeFile(t *testing.T, name, content string) {
	t.Helper()
	if err := os.WriteFile(name, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func readFile(t *testing.T, name string) string {
	t.Helper()
	data, err := os.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

func TestAliasSwap_BackupIsIgnoredByGo(t *testing.T) {
	swap := newAliasSwap(filepath.Join("some", "contracts", "alias"))
	if got, want := filepath.Base(swap.backup), ".alias.replaced"; got != want {
		t.Fatalf("backup name = %q, want %q: it must start with a dot so Go tooling ignores it", got, want)
	}
	if filepath.Dir(swap.backup) != filepath.Dir(swap.path) {
		t.Fatalf("backup %q must be next to the alias %q (same file system, atomic rename)", swap.backup, swap.path)
	}
}

func TestAliasSwap_NoPreviousAlias(t *testing.T) {
	aliasPath := filepath.Join(t.TempDir(), "contracts", "alias")
	swap := newAliasSwap(aliasPath)

	if err := swap.moveAside(); err != nil {
		t.Fatal(err)
	}
	if swap.moved {
		t.Fatal("nothing was moved: there was no alias")
	}
	// Whatever was generated is left as it is, as it is without the swap
	if err := os.MkdirAll(aliasPath, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := swap.restore(); err != nil {
		t.Fatal(err)
	}
	if !exists(aliasPath) {
		t.Fatal("restore must not remove what was generated when there was no previous alias")
	}
	if err := swap.discard(); err != nil {
		t.Fatal(err)
	}
}

func TestAliasSwap_Restore(t *testing.T) {
	aliasPath := newAlias(t, "previous")
	swap := newAliasSwap(aliasPath)

	if err := swap.moveAside(); err != nil {
		t.Fatal(err)
	}
	if exists(aliasPath) {
		t.Fatal("the alias must be moved aside")
	}
	// The new one is partially generated when the import fails
	if err := os.MkdirAll(aliasPath, 0o755); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(aliasPath, "partial.go"), "partial")

	if err := swap.restore(); err != nil {
		t.Fatal(err)
	}
	if got := readFile(t, filepath.Join(aliasPath, "previous.go")); got != "previous" {
		t.Fatalf("the previous alias was not restored: %q", got)
	}
	if exists(filepath.Join(aliasPath, "partial.go")) {
		t.Fatal("what was generated must be removed")
	}
	if exists(swap.backup) {
		t.Fatal("the backup must not be left behind")
	}
}

func TestAliasSwap_Discard(t *testing.T) {
	aliasPath := newAlias(t, "previous")
	swap := newAliasSwap(aliasPath)

	if err := swap.moveAside(); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(aliasPath, 0o755); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(aliasPath, "new.go"), "new")

	if err := swap.discard(); err != nil {
		t.Fatal(err)
	}
	if exists(swap.backup) {
		t.Fatal("the previous alias must be removed once the new one is complete")
	}
	if got := readFile(t, filepath.Join(aliasPath, "new.go")); got != "new" {
		t.Fatalf("the new alias must be kept: %q", got)
	}
	if exists(filepath.Join(aliasPath, "previous.go")) {
		t.Fatal("nothing of the previous alias must be in the new one")
	}
}

// The import was interrupted after moving the alias aside and before creating the new one: the backup is still
// there and the alias is not. It must not be ignored, or a later failure would leave the previous alias lost
func TestAliasSwap_InterruptedImportWithoutAlias(t *testing.T) {
	aliasPath := filepath.Join(t.TempDir(), "contracts", "alias")
	swap := newAliasSwap(aliasPath)
	if err := os.MkdirAll(swap.backup, 0o755); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(swap.backup, "previous.go"), "previous")

	err := swap.moveAside()
	if err == nil || !strings.Contains(err.Error(), "interrupted") {
		t.Fatalf("expected an error about the interrupted import, got %v", err)
	}
	if got := readFile(t, filepath.Join(swap.backup, "previous.go")); got != "previous" {
		t.Fatalf("the backup was modified: %q", got)
	}
}

// A backup that is already there is the previous alias of an import that was interrupted: it must not be overwritten
func TestAliasSwap_InterruptedImportIsNotOverwritten(t *testing.T) {
	aliasPath := newAlias(t, "previous")
	swap := newAliasSwap(aliasPath)
	if err := os.MkdirAll(swap.backup, 0o755); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(swap.backup, "older.go"), "older")

	err := swap.moveAside()
	if err == nil || !strings.Contains(err.Error(), "interrupted") {
		t.Fatalf("expected an error about the interrupted import, got %v", err)
	}
	if swap.moved {
		t.Fatal("nothing must be moved")
	}
	if got := readFile(t, filepath.Join(swap.backup, "older.go")); got != "older" {
		t.Fatalf("the backup of the interrupted import was modified: %q", got)
	}
	if got := readFile(t, filepath.Join(aliasPath, "previous.go")); got != "previous" {
		t.Fatalf("the alias was modified: %q", got)
	}
}
