package btree

import (
	"os"
	"testing"
)

func Test_creation(t *testing.T) {
	tempfile, err := os.CreateTemp("/tmp", "")
	if err != nil {
		t.Fatalf("failed to create temp file: %s", err.Error())
	}
	btree, err := NewBTree(tempfile)
	if err != nil {
		t.Fatalf("failed to create btree: %s", err.Error())
	}
	if btree.Root == nil {
		t.Errorf("Root node was not created")
	}
	info, err := tempfile.Stat()
	if err != nil {
		t.Fatal("tempfile corrupted")
	}
	if info.Size() != 4096 {
		t.Errorf("File size after initiation should be 4096, instead for %d", info.Size())
	}
}

func Test_creation_from_file_with_existing_data(t *testing.T) {
	tempfile, err := os.CreateTemp("/tmp", "")
	if err != nil {
		t.Fatalf("failed to create temp file: %s", err.Error())
	}
	btree, err := NewBTree(tempfile)
	if err != nil {
		t.Fatalf("failed to create btree: %s", err.Error())
	}
	if btree.Root == nil {
		t.Errorf("Root node was not created")
	}
	info, err := tempfile.Stat()
	if err != nil {
		t.Fatal("tempfile corrupted")
	}
	if info.Size() != 4096 {
		t.Errorf("File size after initiation should be 4096, instead for %d", info.Size())
	}

	err = btree.Insert("a", "a")
	if err != nil {
		t.Fatalf("failed to insert btree: %s", err.Error())
	}

	newBtree, err := NewBTree(tempfile)
	if err != nil {
		t.Fatalf("failed to create btree: %s", err.Error())
	}
	val, exist := newBtree.Search("a")
	if !exist || val != "a" {
		t.Fatalf("wrong retrieved value, expected b got: %s", val)
	}

}

func Test_insert_and_search(t *testing.T) {
	tempfile, err := os.CreateTemp("/tmp", "")
	if err != nil {
		t.Fatalf("failed to create temp file: %s", err.Error())
	}
	btree, err := NewBTree(tempfile)
	if err != nil {
		t.Fatalf("failed to create btree: %s", err.Error())
	}

	// Basic insert and search functionality
	err = btree.Insert("a", "a")
	if err != nil {
		t.Fatalf("failed to insert btree: %s", err.Error())
	}
	val, exist := btree.Search("a")
	if !exist || val != "a" {
		t.Fatalf("wrong retrieved value, expected a got: %s", val)
	}

	// Search for non existent
	_, exist = btree.Search("b")
	if exist {
		t.Fatal("returned a value for non existent key: b")
	}

	// Key override
	err = btree.Insert("a", "b")
	if err != nil {
		t.Fatalf("failed to insert btree: %s", err.Error())
	}
	val, exist = btree.Search("a")
	if !exist || val != "b" {
		t.Fatalf("wrong retrieved value, expected b got: %s", val)
	}

}
