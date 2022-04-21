package timescale

import (
	"fmt"
	"os"

	"strings"
	"sync"
	"time"

	"github.com/Akkurate/utils/logging"
	"github.com/Akkurate/utils/system"
	"github.com/jmoiron/sqlx"

	// should be imported
	_ "github.com/lib/pq"
)

var debug = system.GetEnvOrDefault("QUERY_DEBUG", "") == "1"

// Timescale parameters.
type Timescale struct {
	uri            string
	DB             *sqlx.DB
	MaxConnections int    // Max. number of connections, 50 as default
	Error          string // Error string
	sync.Mutex
}

// New Timescale connection.
func New(uri string) (timescale *Timescale) {
	timescale = &Timescale{
		uri:            uri,
		MaxConnections: 50, // hardcoded for now
	}
	err := timescale.Connect()
	if err != nil {
		timescale.Error = fmt.Sprintf("%v", err)
		logging.Error("%v", err)
	}
	return timescale
}

// Connect to Timescale. SSL is set as disabled and hostname is used as the application name. In case of not being able to connect, Error string is updated.
func (ts *Timescale) Connect() (err error) {

	hostname, _ := os.Hostname()
	a := strings.Index(ts.uri, "?sslmode=disable")
	if a == -1 {
		ts.uri = fmt.Sprintf("%v?sslmode=disable", ts.uri)
	}
	ts.uri = fmt.Sprintf("%v&fallback_application_name=%v", ts.uri, hostname)
	logging.Info("<cyan>** Connecting to Timescale</> <gray>%v</>", ts.uri)

	var db *sqlx.DB

	maxRetries := 3
	for retry := 0; retry < maxRetries; retry++ {
		db, err = sqlx.Connect("postgres", ts.uri)

		if err != nil {
			logging.Warn("<cyan>** Couldn't connect to Timescale. Retrying in 10 seconds...</>")
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}

	if err != nil {
		return err
	}

	db.SetMaxOpenConns(ts.MaxConnections)
	ts.DB = db

	return err
}

// Close connection.
func (ts *Timescale) Close() {
	ts.DB.Close()
}
