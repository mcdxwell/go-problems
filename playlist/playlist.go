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

type PlayMap map[string]struct{}

func DeleteWGenerics[K string, V any]() {

}

type T interface {
	~int | string | struct{}
}

// ref: https://cs.opensource.google/go/x/exp/+/39d4317d:maps/maps.go;l=77
func Copy[M ~map[K]V, K comparable, V any](dst, src M) {
	for k, v := range src {
		dst[k] = v
	}
}

// TODO
// Merge music libraries (merge maps to create one map with a bunch of linked lists/playlists)
func MergeLib[M ~map[K]V, K comparable, V any](dst, src M) {
	Copy(dst, src)
	fmt.Println(dst)
	fmt.Println(src)
}

// TODO
// Merge music playlists (merge linked lists into 1 linked list, delete copies)

func RemoveDuplicates[M ~map[K]V, K comparable, V any]() {

}

type Remover interface {
	map[string]any | Playlist
}

// Remove linked list dupes

func PlaylistDupes(a, b Playlist) Playlist {
	if a.Start == nil || b.Start == nil {
		return Playlist{}
	}

	if a.Start != nil && b.Start == nil {
		return a
	}

	if a.Start == nil && b.Start != nil {
		return b
	}
	// Make map
	// Using the showAllSongs logic, get every song from both Playlists
	// Put each song in a map if the song does not exist in the map
	// Get each value from the map and create a new Playlist (linked list) with the values

	return Playlist{}
}
