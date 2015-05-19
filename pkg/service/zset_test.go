// Copyright 2015 Reborndb Org. All Rights Reserved.
// Licensed under the MIT (MIT-LICENSE.txt) license.

package service

import (
	"strconv"

	. "gopkg.in/check.v1"
)

func (s *testServiceSuite) checkZSet(c *C, k string, expect map[string]int64) {
	ay := s.checkBytesArray(c, "zgetall", k)
	if expect == nil {
		c.Assert(ay, IsNil)
	} else {
		c.Assert(ay, HasLen, len(expect)*2)
		for i := 0; i < len(expect); i++ {
			k := string(ay[i*2])
			v := string(ay[i*2+1])
			f, err := strconv.ParseInt(v, 10, 64)
			c.Assert(err, IsNil)
			c.Assert(expect[k], Equals, f)
		}
	}
}

func (s *testServiceSuite) TestZAdd(c *C) {
	k := randomKey(c)
	s.checkInt(c, 1, "zadd", k, 1, "one")
	s.checkInt(c, 2, "zadd", k, 2, "two", 3, "three")
	s.checkInt(c, 1, "zadd", k, 1, "one", 4, "four", 5, "four")
	s.checkZSet(c, k, map[string]int64{"one": 1, "two": 2, "three": 3, "four": 5})
	s.checkInt(c, 0, "zadd", k, 1, "one", 4, "four")
	s.checkZSet(c, k, map[string]int64{"one": 1, "two": 2, "three": 3, "four": 4})
}

func (s *testServiceSuite) TestZCard(c *C) {
	k := randomKey(c)
	s.checkInt(c, 0, "zcard", k)
	s.checkInt(c, 1, "zadd", k, 1, "one")
	s.checkInt(c, 1, "zcard", k)
	s.checkInt(c, 2, "zadd", k, 2, "two", 3, "three")
	s.checkInt(c, 3, "zcard", k)
	s.checkInt(c, 0, "zadd", k, 4, "two")
	s.checkInt(c, 3, "zcard", k)
}

func (s *testServiceSuite) TestZScore(c *C) {
	k := randomKey(c)
	s.checkNil(c, "zscore", k, "one")
	s.checkInt(c, 1, "zadd", k, 1, "one")
	s.checkFloat(c, 1, "zscore", k, "one")
	s.checkNil(c, "zscore", k, "two")
}

func (s *testServiceSuite) TestZRem(c *C) {
	k := randomKey(c)
	s.checkInt(c, 3, "zadd", k, 1, "key1", 2, "key2", 3, "key3")
	s.checkInt(c, 0, "zrem", k, "key")
	s.checkInt(c, 1, "zrem", k, "key1")
	s.checkZSet(c, k, map[string]int64{"key2": 2, "key3": 3})
	s.checkInt(c, 2, "zrem", k, "key1", "key2", "key3")
	s.checkZSet(c, k, nil)
	s.checkInt(c, -2, "ttl", k)
}

func (s *testServiceSuite) TestZIncrBy(c *C) {
	k := randomKey(c)
	s.checkFloat(c, 1, "zincrby", k, 1, "one")
	s.checkFloat(c, 1, "zincrby", k, 1, "two")
	s.checkFloat(c, 2, "zincrby", k, 1, "two")
	s.checkZSet(c, k, map[string]int64{"one": 1, "two": 2})
}
