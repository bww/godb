package persist

import (
	"os"
	"testing"

	"github.com/bww/godb/test"
)

const dbname = "godb_db_persist_test"
const table = "godb_persist_test"

func TestMain(m *testing.M) {
	test.Init(dbname, false)
	os.Exit(m.Run())
}
