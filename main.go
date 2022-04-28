package main

import (
	"fmt"

	"github.com/mcdxwell/go-problems/playlist"
)

func main() {

	/* x := adder.Adder(1, 2)
	z := adder.Adder("hello", "world")
	fmt.Println(x, z)

	l := linkedlist.LinkedList{}

	fmt.Println("\n************* Insert *************")
	l.Insert(12)
	l.Insert(13)
	l.Insert(14)
	l.Insert(15)
	fmt.Println("************* Print *************")
	l.Print()

	fmt.Println("\n************* Search *************")
	fmt.Println("Position of 12 value: ", l.Search(12))
	fmt.Println("Position of 14 value: ", l.Search(14))
	fmt.Println("Position of 15 value: ", l.Search(15))
	fmt.Println("Position of 100 value: ", l.Search(100))

	fmt.Println("\n************* GetAt *************")
	fmt.Println("Get At 1st position: ", l.GetAt(0))
	fmt.Println("Get At 3rd position: ", l.GetAt(2))
	fmt.Println("Get At 4th position: ", l.GetAt(3))
	fmt.Println("Get At -5 position (head is returned): ", l.GetAt(-5))

	// fmt.Println("\n************* InsertAt *************")
	// l.InsertAt(0, 12)
	// l.InsertAt(-1, 13)
	// l.InsertAt(1, 14)
	// l.InsertAt(2, 15)
	// fmt.Println("************* Print *************")
	// l.Print()

	fmt.Println("\n************* DeleteAt *************")
	l.DeleteAt(3)
	fmt.Println("************* Print *************")
	l.Print()

	// fmt.Println("\n************* DeleteVal *************")
	// l.DeleteVal(14)
	// fmt.Println("************* Print *************")
	// l.Print() */
	playlistName := "Random Playlist"
	myPlaylist := playlist.CreatePlaylist(playlistName)
	fmt.Printf("Created playlist: %v\n", playlistName)
	fmt.Printf("Adding songs to %v\n", playlistName)
	myPlaylist.AddSong("Ophelia", "The Lumineers")
	myPlaylist.AddSong("Shape of you", "Ed Sheeran")
	myPlaylist.AddSong("Stubborn Love", "The Lumineers")
	myPlaylist.AddSong("Feels", "Calvin Harris")
	//fmt.Println("Showing all songs in playlist...")
	//myPlaylist.ShowAllSongs()

	//myPlaylist.StartPlaying()
	//fmt.Printf("Now playing: %s by %s\n", myPlaylist.NowPlaying.Next, myPlaylist.NowPlaying.Artist)

	//myPlaylist.NextSong()
	//fmt.Println("Changing next song...")
	//fmt.Printf("Now playing: %s by %s\n", myPlaylist.NowPlaying.Name, myPlaylist.NowPlaying.Artist)
	//fmt.Println()
	//myPlaylist.NextSong()
	//fmt.Println("Changing next song...")
	//fmt.Printf("Now playing: %s by %s\n", myPlaylist.NowPlaying.Name, myPlaylist.NowPlaying.Artist)
	m := make(map[string]playlist.Playlist)

	m[playlistName] = *myPlaylist
	playlist.DisplayPlaylists(m)
	secondPlayList := "David's Playlist"
	mysecPlaylist := playlist.CreatePlaylist(secondPlayList)
	fmt.Printf("Created playlist: %v\n", secondPlayList)
	fmt.Printf("Adding songs to %v\n", secondPlayList)
	mysecPlaylist.AddSong("Bangarang", "Skrillex")
	mysecPlaylist.AddSong("Visions of Johana", "Bob Dylan")
	mysecPlaylist.AddSong("Under the Bridge", "Red Hot Chili Peppers")
	mysecPlaylist.AddSong("Come As You Are", "Nirvana")
	mysecPlaylist.AddSong("Something in the way", "Nirvana")
	//fmt.Printf("Showing all songs in %v\n", secondPlayList)
	//mysecPlaylist.ShowAllSongs()
	//fmt.Println("Created playlist")

	playlist.AddToLib(mysecPlaylist, m)
	playlist.DisplayPlaylists(m)
	mysecPlaylist.StartPlaying()
	fmt.Println(mysecPlaylist.NowPlaying.Name)
	mysecPlaylist.NextSong()
	fmt.Println(mysecPlaylist.NowPlaying.Name)
	mysecPlaylist.NextSong()
	fmt.Println(mysecPlaylist.NowPlaying.Name)
	mysecPlaylist.NextSong()
	fmt.Println(mysecPlaylist.NowPlaying.Name)
	mysecPlaylist.NextSong()
	fmt.Println(mysecPlaylist.NowPlaying.Name)

	mysecPlaylist.PrevSong()
	fmt.Println(mysecPlaylist.NowPlaying.Name)
	mysecPlaylist.PrevSong()
	fmt.Println(mysecPlaylist.NowPlaying.Name)
	mysecPlaylist.PrevSong()
	fmt.Println(mysecPlaylist.NowPlaying.Name)
	mysecPlaylist.PrevSong()
	fmt.Println(mysecPlaylist.NowPlaying.Name)

	playlist.DeletePlaylists(m, "Random Playlist", "David's Playlist")

	playlist.DisplayPlaylists(m)

}
