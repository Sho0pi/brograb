package brograb

import "errors"

var (
	BadDestFile = errors.New("bad destination type")
)

// Grabber is a basic data grabbing interface. Use Next and Scan to fetch the next grabbing data.
type Grabber interface {

	// Next prepares the next fetch data for reading with the Scan method. It
	// returns true on success, or false if there is no next result row or an error
	// happened while preparing it. Err should be consulted to distinguish between
	// the two cases.
	//
	// Every call to Scan, even the first one, must be preceded by a call to Next.
	Next() bool

	// Scan copies the current fetched data into the interface pointed at
	// by dest.
	Scan(dest interface{}) error

	// Err returns the error, if any, that was encountered during iteration.
	Err() error

	// Close stops the iteration, and closes all open databases, connections,
	// and files used to fetch the browser data.
	Close() error
}
