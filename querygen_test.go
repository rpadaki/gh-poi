package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetQueryHashes(t *testing.T) {
	assert.Equal(t,
		[]string{
			"hash:356a192b7913b04c54574d18c28d46e6395428ab " +
				"hash:da4b9237bacccdf19c0760cab7aec4a8359010b0 " +
				"hash:77de68daecd823babbb58edb1c8e14d7106e83bb " +
				"hash:1b6453892473a467d07372d45eb05abc2031647a " +
				"hash:ac3478d69a3c81fa62e60f5c3696165a4e5e6ac4 ",
			"hash:c1dfd96eea8cc2b62785275bca38ac261256e278",
		},
		GetQueryHashes([]Branch{
			Branch{false, "issue1", "356a192b7913b04c54574d18c28d46e6395428ab", []PullRequest{}, Unknown},
			Branch{false, "issue2", "da4b9237bacccdf19c0760cab7aec4a8359010b0", []PullRequest{}, Unknown},
			Branch{false, "issue3", "77de68daecd823babbb58edb1c8e14d7106e83bb", []PullRequest{}, Unknown},
			Branch{false, "issue4", "1b6453892473a467d07372d45eb05abc2031647a", []PullRequest{}, Unknown},
			Branch{false, "issue5", "ac3478d69a3c81fa62e60f5c3696165a4e5e6ac4", []PullRequest{}, Unknown},
			Branch{false, "issue6", "c1dfd96eea8cc2b62785275bca38ac261256e278", []PullRequest{}, Unknown},
		}))
}