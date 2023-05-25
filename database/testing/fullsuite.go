// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package testing

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"sync/atomic"

	"github.com/canonical/sqlair"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	coredatabase "github.com/juju/juju/core/database"
	"github.com/juju/juju/database/app"
	"github.com/juju/juju/domain/schema"
)

// DBSuite is used to provide a Dqlite-backed sql.DB reference to tests.
type DBSuite struct {
	testing.IsolationSuite

	dbPath   string
	rootPath string
	uniqueID int64

	dqlite    *app.App
	db        *sql.DB
	trackedDB coredatabase.TxnRunner
}

// SetUpSuite creates a new Dqlite application and waits for it to be ready.
func (s *DBSuite) SetUpSuite(c *gc.C) {
	s.IsolationSuite.SetUpSuite(c)

	s.rootPath = c.MkDir()

	path := filepath.Join(s.rootPath, "dqlite")
	err := os.Mkdir(path, 0700)
	c.Assert(err, jc.ErrorIsNil)
	s.dbPath = path

	port := FindTCPPort(c)

	url := fmt.Sprintf("%s:%d", "127.0.0.1", port)
	c.Logf("Opening sqlite3 db with: %v", url)

	s.dqlite, err = app.New(s.dbPath, app.WithAddress(url))
	c.Assert(err, jc.ErrorIsNil)

	err = s.dqlite.Ready(context.TODO())
	c.Assert(err, jc.ErrorIsNil)
}

// TearDownSuite terminates the Dqlite node, releasing all resources.
func (s *DBSuite) TearDownSuite(c *gc.C) {
	if s.dqlite != nil {
		err := s.dqlite.Close()
		c.Assert(err, jc.ErrorIsNil)
	}

	s.IsolationSuite.TearDownSuite(c)
}

// SetUpTest opens a new, randomly named database and
// makes it available for use by test the next test.
func (s *DBSuite) SetUpTest(c *gc.C) {
	s.IsolationSuite.SetUpTest(c)

	// Increment the id and use it as the database name, this prevents
	// tests from interfering with each other.
	uniqueID := atomic.AddInt64(&s.uniqueID, 1)

	var err error
	s.db, err = s.dqlite.Open(context.TODO(), strconv.FormatInt(uniqueID, 10))
	c.Assert(err, jc.ErrorIsNil)

	s.trackedDB = &txnRunner{
		db: sqlair.NewDB(s.db),
	}

	s.ApplyControllerDDL(c)
}

// TearDownTest closes the database opened in SetUpTest.
// TODO (manadart 2022-09-12): There is currently no avenue for dropping a DB.
func (s *DBSuite) TearDownTest(c *gc.C) {
	if s.db != nil {
		c.Logf("Closing DB")
		err := s.db.Close()
		c.Assert(err, jc.ErrorIsNil)
	}

	s.IsolationSuite.TearDownTest(c)
}

// DB returns the database for the current test.
func (s *DBSuite) DB() *sql.DB {
	return s.db
}

// TxnRunner returns the tracked database for the current test.
func (s *DBSuite) TxnRunner() coredatabase.TxnRunner {
	return s.trackedDB
}

// DBApp returns the dqlite application.
func (s *DBSuite) DBApp() *app.App {
	return s.dqlite
}

// DBPath returns the path to the dqlite database.
func (s *DBSuite) DBPath() string {
	return s.dbPath
}

// RootPath returns the path to the root directory, where the DBPath is located.
func (s *DBSuite) RootPath() string {
	return s.rootPath
}

// FindTCPPort finds an unused TCP port and returns it.
// It is prone to racing, so the port should be used as soon as it is acquired
// to minimise the change of another process using it in the interim.
// The chances of this should be negligible during testing.
func FindTCPPort(c *gc.C) int {
	l, err := net.Listen("tcp", ":0")
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(l.Close(), jc.ErrorIsNil)
	return l.Addr().(*net.TCPAddr).Port
}

// ApplyControllerDDL applies the controller schema to the provided sql.DB.
// This is useful for tests that need to apply the schema to a new DB.
func (s *DBSuite) ApplyControllerDDL(c *gc.C) {
	tx, err := s.db.Begin()
	c.Assert(err, jc.ErrorIsNil)

	for idx, delta := range schema.ControllerDDL(0x2dc171858c3155be) {
		c.Logf("Executing schema DDL index: %v", idx)
		_, err := tx.Exec(delta.Stmt(), delta.Args()...)
		c.Assert(err, jc.ErrorIsNil)
	}

	c.Logf("Committing schema DDL")
	err = tx.Commit()
	c.Assert(err, jc.ErrorIsNil)
}
