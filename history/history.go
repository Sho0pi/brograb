package history

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sho0pi/brograb"
	"github.com/sho0pi/brograb/browserutils"
	"time"
)

const (
	queryChromiumHistory = `SELECT url, title, visit_count, last_visit_time FROM urls`
)

// HistoryGrabber extends the Grabber interface, for bookmarks specific grabbing.
type HistoryGrabber interface {
	Iterate() <-chan History
	brograb.Grabber
}

// History represents the history from the database of the browserutils.
type History struct {
	URL         string
	Title       string
	VisitCount  int
	LastVisited time.Time
}

// ChromiumGrabber is a grabber for chromium history entries.
type ChromiumGrabber struct {
	HistoryGrabber
	db   *sql.DB
	rows *sql.Rows
}

// Close closes the connected history database.
func (c *ChromiumGrabber) Close() error {
	defer c.db.Close()
	if err := c.rows.Close(); err != nil {
		return err
	}
	if err := c.db.Close(); err != nil {
		return err
	}
	return nil
}

// Err returns the error, if any that was encountered during iteration.
func (c *ChromiumGrabber) Err() error {
	return c.rows.Err()
}

// Next fetches the next History.
func (c *ChromiumGrabber) Next() bool {
	return c.rows.Next()
}

// Scan copies the columns in the current row into the History pointed
// at by dest.
// The dest var must be of type History or else the method will fail.
func (c *ChromiumGrabber) Scan(dest interface{}) error {
	historyEntry, ok := dest.(*History)
	if !ok {
		return brograb.BadDestFile
	}
	var (
		url             string
		title           string
		visitedCount    int
		lastVisitedTime int64
	)

	if err := c.rows.Scan(&url, &title, &visitedCount, &lastVisitedTime); err != nil {
		return err
	}
	historyEntry.URL = url
	historyEntry.Title = title
	historyEntry.VisitCount = visitedCount
	historyEntry.LastVisited = browserutils.FormatChromiumEpoch(lastVisitedTime)
	return nil
}

// NewChromiumGrabber returns a new History grabber that grabs history from the given date.
// If you specify date to be nil, it will grab all the History from the db.
func NewChromiumGrabber(profile browserutils.ProfileDir, date time.Time) (*ChromiumGrabber, error) {
	historyDb, err := sql.Open("sqlite3", profile.HistoryDB())
	if err != nil {
		return nil, err
	}
	// Todo: query from specific time.
	rows, err := historyDb.Query(queryChromiumHistory)
	if err != nil {
		return nil, err
	}
	return &ChromiumGrabber{
		db:   historyDb,
		rows: rows,
	}, nil
}
