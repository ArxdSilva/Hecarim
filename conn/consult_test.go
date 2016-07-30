package conn

import "gopkg.in/check.v1"

func testConsult(c *check.C) {
	c.Check(GetPlayer, check.DeepEquals, []interface{}{""})
}
