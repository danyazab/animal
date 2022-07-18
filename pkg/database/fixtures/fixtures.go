package fixtures

import (
	"io/ioutil"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/romanyx/polluter"
)

func PrepareTestDatabase(db *sqlx.DB, paths ...string) {
	p := polluter.New(polluter.PostgresEngine(db.DB))
	for _, f := range paths {
		content, err := ioutil.ReadFile(f)
		if err != nil {
			panic(err)
		}

		if err := p.Pollute(strings.NewReader(string(content))); err != nil {
			panic(err)
		}
	}
}
