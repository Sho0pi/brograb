package bookmarks

import (
	"errors"
	"github.com/sho0pi/brograb"
	"github.com/sho0pi/brograb/browserutils"
	"github.com/tidwall/gjson"
	"time"
)

type BookmarkType int

const (
	URL BookmarkType = iota
	FOLDER
)

const (
	bookmarkTypeKey      = "type"
	bookmarkUrlKey       = "url"
	bookmarkNameKey      = "name"
	bookmarkDateAddedKey = "date_added"
	bookmarkChildrenKey  = "children"
)

var (
	invalidBookmarkType  = errors.New("bookmark is not of type folder")
	badDestFile          = errors.New("bad dest type")
	bookmarkNameNotFound = errors.New("unable to find bookmark display name")
	bookmarkTypeNotFound = errors.New("unable to fetch bookmark type")
	dateAddedNotFound    = errors.New("unable to get bookmark creation date")
	urlAddedNotFound     = errors.New("unable to fetch bookmark url")
	childrenNotFound     = errors.New("unable to fetch bookmark children of folder")
)

// BookmarkGrabber extends the Grabber interface, for bookmarks specific grabbing.
type BookmarkGrabber interface {
	Iterate() <-chan Bookmark
	brograb.Grabber
}

// Bookmark represents a chromium browser bookmark.
// Bookmark can be of type URL or of type FOLDER
type Bookmark struct {
	BookmarkGrabber
	DateAdded time.Time
	Name      string
	URL       string
	Type      BookmarkType
	bookmarks []gjson.Result
	index     int
}

func (b Bookmark) Err() error {
	return nil
}

// IsFolder returns true if the bookmark type is folder.
// This will help you to iterate on nested  bookmarks and folders.
func (b Bookmark) IsFolder() bool {
	return b.Type == FOLDER
}

// Close stops any future calls to Next and Scan
func (b *Bookmark) Close() error {
	b.bookmarks = nil
	b.index = -1
	return nil
}

// Next prepare the next bookmark to be returned.
// You can call Next only of bookmarks of type FOLDER.
func (b *Bookmark) Next() bool {
	b.index++
	// After Close this will return false always.
	if b.index >= len(b.bookmarks) {
		return false
	}
	return true
}

// Scan scans the next bookmark entry and returns it into the dest.
// You can call Scan only of bookmarks of type FOLDER.
func (b *Bookmark) Scan(dest interface{}) error {
	bookmarkDest, ok := dest.(*Bookmark)
	if !ok {
		return badDestFile
	}
	if !b.IsFolder() {
		return invalidBookmarkType
	}

	// Scan be always called after next
	entry := b.bookmarks[b.index]

	if err := bookmarkDest.setDateAdded(entry); err != nil {
		return err
	}

	if err := bookmarkDest.setName(entry); err != nil {
		return err
	}

	if err := bookmarkDest.setType(entry); err != nil {
		return err
	}

	switch bookmarkDest.IsFolder() {
	case true:
		if err := bookmarkDest.setChildren(entry); err != nil {
			return err
		}
	case false:
		if err := bookmarkDest.setURL(entry); err != nil {
			return err
		}
	}

	return nil
}

// Iterate lets you iterate on all the available bookmarks with a basic for loop.
func (b *Bookmark) Iterate() <-chan Bookmark {
	ch := make(chan Bookmark, 16)
	go func() {
		defer close(ch)
		defer b.Close()
		for b.Next() {
			var bookmark Bookmark
			if err := b.Scan(&bookmark); err != nil {
				break
			}
			ch <- bookmark
			if bookmark.IsFolder() {
				for nestedBookmark := range bookmark.Iterate() {
					ch <- nestedBookmark
				}
			}
		}
	}()
	return ch
}

func (b *Bookmark) setDateAdded(entry gjson.Result) error {
	if d := entry.Get(bookmarkDateAddedKey); d.Exists() {
		b.DateAdded = browserutils.FormatChromiumEpoch(d.Int())
		return nil
	}
	return dateAddedNotFound
}

func (b *Bookmark) setName(entry gjson.Result) error {
	if name := entry.Get(bookmarkNameKey); name.Exists() {
		b.Name = name.String()
		return nil
	}
	return bookmarkNameNotFound
}

func (b *Bookmark) setChildren(entry gjson.Result) error {
	if children := entry.Get(bookmarkChildrenKey); children.Exists() {
		b.bookmarks = children.Array()
		b.index = -1
		return nil
	}
	return childrenNotFound
}

func (b *Bookmark) setURL(entry gjson.Result) error {
	if url := entry.Get(bookmarkUrlKey); url.Exists() {
		b.URL = url.String()
		return nil
	}
	return urlAddedNotFound
}

func (b *Bookmark) setType(entry gjson.Result) error {
	if t := entry.Get(bookmarkTypeKey); t.Exists() {
		if t.String() == "folder" {
			b.Type = FOLDER
		} else {
			b.Type = URL
		}
		return nil
	}
	return bookmarkTypeNotFound
}
