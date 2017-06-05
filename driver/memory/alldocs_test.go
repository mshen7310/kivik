package memory

import (
	"context"
	"io"
	"sort"
	"testing"

	"github.com/flimzy/diff"
	"github.com/flimzy/kivik/driver"
)

func TestAllDocs(t *testing.T) {
	type adTest struct {
		Name        string
		ExpectedIDs []string
		Error       string
		DB          driver.DB
		RowsError   string
	}
	tests := []adTest{
		{
			Name: "NoDocs",
		},
	}
	for _, test := range tests {
		func(test adTest) {
			t.Run(test.Name, func(t *testing.T) {
				db := test.DB
				if db == nil {
					db = setupDB(t, nil)
				}
				rows, err := db.AllDocs(context.Background(), nil)
				var msg string
				if err != nil {
					msg = err.Error()
				}
				if test.Error != msg {
					t.Errorf("Unexpected error: %s", msg)
				}
				if err != nil {
					return
				}
				var row *driver.Row
				var ids []string
				msg = ""
				for {
					e := rows.Next(row)
					if e != nil {
						if e != io.EOF {
							msg = e.Error()
						}
						break
					}
					ids = append(ids, row.ID)
				}
				if test.RowsError != msg {
					t.Errorf("Unexpected rows error: %s", msg)
				}
				sort.Strings(ids)
				if d := diff.TextSlices(test.ExpectedIDs, ids); d != "" {
					t.Error(d)
				}
			})
		}(test)
	}
}
