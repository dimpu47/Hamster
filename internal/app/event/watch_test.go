// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package event

import (
	"github.com/nbio/st"
	"io/ioutil"
	"testing"
)

// TestWatch test cases
func TestWatch(t *testing.T) {

	var watch Watch

	dat, err := ioutil.ReadFile("../../../fixtures/watch.json")

	if err != nil {
		t.Errorf("File fixtures/watch.json is invalid!")
	}

	ok, _ := watch.LoadFromJSON(dat)

	if !ok {
		t.Errorf("Testing with file fixtures/watch.json is invalid")
	}

	st.Expect(t, watch.Repository.Owner.Login, "Clivern")
}
