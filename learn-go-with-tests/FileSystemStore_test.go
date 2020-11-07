package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func createTmpFile(t *testing.T, initData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	tmpfile.Write([]byte(initData))

	removeFile := func() {
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didnt expect an error but got one, %v", err)
	}
}
func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTmpFile(t, `[{"Name":"likuo","Wins":10},{"Name":"xiong","Wins":20}]`)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		want := []Player{
			{"xiong", 20},
			{"likuo", 10},
		}
		got := store.GetLeague()
		assertLeague(t, got, want)
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTmpFile(t, `[{"Name":"likuo","Wins":10},{"Name":"xiong","Wins":20}]`)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		got := store.GetPlayerScore("likuo")

		want := 10

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTmpFile(t, `[{"Name":"likuo","Wins":10},{"Name":"xiong","Wins":20}]`)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		store.RecordWin("xiong")

		got := store.GetPlayerScore("xiong")
		want := 21
		assertScoreEquals(t, got, want)

		store.RecordWin("xiong2")
		got = store.GetPlayerScore("xiong2")
		want = 1
		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTmpFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemStore(database)

		assertNoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTmpFile(t, `[{"Name":"likuo","Wins":10},{"Name":"xiong","Wins":20}]`)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		want := []Player{
			{"xiong", 20},
			{"likuo", 10},
		}
		got := store.GetLeague()
		assertLeague(t, got, want)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
