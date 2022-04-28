package playlist

import (
	"fmt"
)

type Song struct {
	Name     string
	Artist   string
	Previous *Song
	Next     *Song
}

type Playlist struct {
	Name       string
	Start      *Song
	End        *Song
	NowPlaying *Song
}

// Credit: https://medium.com/backendarmy/linked-lists-in-go-f7a7b27a03b9
// Mapping Doubly Circular Linked Lists
// create playlist
// add Song
// show all Songs
// start playing
// next/previous Song

func CreatePlaylist(name string) *Playlist {
	return &Playlist{
		Name: name,
	}
}

func (p *Playlist) AddSong(name, artist string) error {
	s := &Song{
		Name:   name,
		Artist: artist,
	}
	if p.Start == nil {
		p.Start = s
	} else {
		currentNode := p.End
		currentNode.Next = s
		s.Previous = p.End
	}

	p.End = s
	return nil
}

func (p *Playlist) ShowAllSongs() error {
	currentNode := p.Start
	if currentNode == nil {
		fmt.Println("Playlist is empty!")
		return nil
	}

	fmt.Printf("%+v\n", *currentNode)
	for currentNode.Next != nil {
		currentNode = currentNode.Next
		fmt.Printf("%+v\n", *currentNode)
	}
	return nil
}

func (p *Playlist) StartPlaying() *Song {
	p.NowPlaying = p.Start
	return p.NowPlaying
}

func (p *Playlist) NextSong() *Song {
	p.NowPlaying = p.NowPlaying.Next
	if p.NowPlaying == nil {
		p.NowPlaying = p.Start
	}
	return p.NowPlaying
}

func (p *Playlist) PrevSong() *Song {
	p.NowPlaying = p.NowPlaying.Previous
	if p.NowPlaying == nil {
		p.NowPlaying = p.End
	}
	return p.NowPlaying
}

func AddToLib(p *Playlist, m map[string]Playlist) {

	m[p.Name] = *p
	fmt.Printf("Added %v to library\n", p.Name)
}

func DisplayPlaylists(m map[string]Playlist) {
	for _, v := range m {
		fmt.Println("Playlist count: ", len(m))
		fmt.Println("Playlist: ", v.Name)
		v.ShowAllSongs()
	}
}

func DeletePlaylists(m map[string]Playlist, playlists ...string) {
	for _, playlist := range playlists {
		_, ok := m[playlist]
		if ok {
			delete(m, playlist)
			fmt.Printf("Deleted %v\n", playlist)
		}
	}
}
