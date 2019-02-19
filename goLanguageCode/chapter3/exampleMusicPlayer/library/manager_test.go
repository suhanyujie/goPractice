package library

import "testing"

func TestOps(t *testing.T)  {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("New MusicManager failed...")
	}
	if mm.Len() != 0 {
		t.Error("MusicManager faild,not empty...")
	}
	m0 := &MusicEntry {
		"1","My heart will go on","Celion Dion",
		"http://qbox.me/24501234","mp3",
	}
	mm.Add(m0)
	if mm.Len()!=1 {
		t.Error("Music Add() faild")
	}
	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManager Find() failed")
	}
	if m.Id != m0.Id || m.Artist != m0.Artist {
		t.Error("MusicManager.Find() failed.Found item is mismatch.")
	}
	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() failed.")
	}
}