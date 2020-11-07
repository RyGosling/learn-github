package poker

import (
	"testing"
)

// 学生玩家
type StuPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StuPlayerStore) GetPlayerScore(name string) int {
	score, ok := s.scores[name]
	if !ok {
		return -1
	}
	return score
}

func (s *StuPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StuPlayerStore) GetLeague() League {
	return s.league
}

func AssertPlayerWin(t *testing.T, store *StuPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got '%s' want '%s'", store.winCalls[0], winner)
	}
}
