// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package scriptstore

import (
	"fmt"
	"os"
	"testing"
)

func TestCall(t *testing.T) {
	LoadDefaultDotPScript2Store()
	ReloadAllGitReposDotPScript2Store(nil)
	ReloadAllRemoteDotPScript2Store(nil)
	_ = os.WriteFile("/tmp/nginx-time123.p", []byte(`
	json(_, time)
	set_tag(bb, "aa0")
	default_time(time)
	`), os.FileMode(0o755))
	LoadDotPScript2StoreWithNS("xxxx", "", []string{"/tmp/nginx-time.p123"})
	_ = os.Remove("/tmp/nginx-time123.p")
	LoadDotPScript2StoreWithNS("xxx", "", nil)
}

func TestPlScriptStore(t *testing.T) {
	store := NewScriptStore()

	err := store.UpdateScriptsWithNS(DefaultScriptNS, map[string]string{"abc.p": "default_time(time)"})
	if err != nil {
		t.Error(err)
	}

	err = store.UpdateScriptsWithNS(GitRepoScriptNS, map[string]string{"abc.p": "default_time(time)"})
	if err != nil {
		t.Error(err)
	}

	err = store.UpdateScriptsWithNS(RemoteScriptNS, map[string]string{"abc.p": "default_time(time)"})
	if err != nil {
		t.Error(err)
	}

	for i, ns := range plScriptNSSearchOrder {
		store.UpdateScriptsWithNS(ns, nil)
		if i < len(plScriptNSSearchOrder)-1 {
			sInfo, ok := store.Get("abc.p")
			if !ok {
				t.Error(fmt.Errorf("!ok"))
				return
			}
			if sInfo.ns != plScriptNSSearchOrder[i+1] {
				t.Error(sInfo.ns, plScriptNSSearchOrder[i+1])
			}
		} else {
			_, ok := store.Get("abc.p")
			if ok {
				t.Error(fmt.Errorf("shoud not be ok"))
				return
			}
		}
	}
}
