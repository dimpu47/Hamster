// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package event

import (
	"github.com/nbio/st"
	"io/ioutil"
	"testing"
)

// TestIssueComment test cases
func TestIssueComment(t *testing.T) {

	var issueComment IssueComment

	dat, err := ioutil.ReadFile("../../../fixtures/issue_comment.json")

	if err != nil {
		t.Errorf("File fixtures/issue_comment.json is invalid!")
	}

	ok, _ := issueComment.LoadFromJSON(dat)

	if !ok {
		t.Errorf("Testing with file fixtures/issue_comment.json is invalid")
	}

	st.Expect(t, issueComment.Issue.User.Login, "Clivern")
}
