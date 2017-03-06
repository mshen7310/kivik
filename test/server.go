package test

import (
	"net/http"

	"github.com/flimzy/kivik/test/kt"
)

func init() {
	RegisterSuite(SuiteKivikServer, kt.SuiteConfig{
		"AllDBs.expected":                            []string{},
		"AllDBs/RW.skip":                             true, // FIXME: Enable this when it's possible to delete DB from the server
		"Config/Admin/GetAll.expected_sections":      []string{"log"},
		"Config/Admin/GetSection.sections":           []string{"log", "chicken"},
		"Config/Admin/GetSection/log.keys":           []string{"capacity"},
		"Config/Admin/GetSection/chicken.keys":       []string{},
		"Config/Admin/GetItem.items":                 []string{"log.capacity", "log.foobar", "logx.level"},
		"Config/Admin/GetItem/log.foobar.status":     http.StatusNotFound,
		"Config/Admin/GetItem/logx.level.status":     http.StatusNotFound,
		"Config/Admin/GetItem/log.capacity.expected": "10",
		"Config/NoAuth.skip":                         true, // FIXME: Update this when the server supports auth
		"Config/RW/NoAuth.skip":                      true, // FIXME: Update this when the server supports auth
		"Config/RW.skip":                             true, // FIXME: Update this when the server can write config
	})
}
