package poker

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type FileSystemStore struct {
	database *json.Encoder
	league   League
}

func initPlayerDBFile(file *os.File) error {
	file.Seek(0, 0)
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}
	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}
	return nil
}

func FileSystemStoreFromFile(path string) (*FileSystemStore, error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("problem opening %s, %v", path, err)
	}

	store, err := NewFileSystemStore(db)
	if err != nil {
		return nil, fmt.Errorf("problem creating file system player store, %v ", err)
	}

	return store, nil
}

func NewFileSystemStore(file *os.File) (*FileSystemStore, error) {

	err := initPlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemStore{
		json.NewEncoder(&tape{file}),
		league,
	}, nil
}

func (f *FileSystemStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return -1
}

func (f *FileSystemStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	//f.database.Seek(0, 0)
	f.database.Encode(f.league)
}
