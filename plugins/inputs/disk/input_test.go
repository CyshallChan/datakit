package disk

import (
	"math"
	"testing"
	"time"

	tu "gitlab.jiagouyun.com/cloudcare-tools/cliutils/testutil"
)

func TestCollect(t *testing.T) {
	//i := &Input{
	//	MountPoints: []string{},
	//	IgnoreFS: []string{
	//		"tmpfs", "devtmpfs",
	//		"devfs", "iso9660",
	//		"overlay", "aufs", "squashfs",
	//	},
	//	Tags: map[string]string{
	//		"tag1": "a",
	//	},
	//	diskStats: &PSDisk{},
	//}
	i := newDefaultInput()
	for x := 0; x < 5; x++ {
		if err := i.Collect(); err != nil {
			t.Error(err)
		}
		time.Sleep(time.Second * 1)
	}
	if len(i.collectCache) < 1 {
		t.Error("Failed to collect, no data returned")
	}
	tmap := map[string]bool{}
	for _, v := range i.collectCache {
		m, ok := v.(*diskMeasurement)
		if !ok {
			t.Error("v expect to be *diskMeasurement")
			continue
		}

		tmap[m.ts.String()] = true
	}
	if len(tmap) != 1 {
		t.Error("Need to clear collectCache.")
	}
}

func TestWrapUint64(t *testing.T) {
	tu.Assert(t, wrapUint64(math.MaxInt64+1) == -1, "")
	tu.Assert(t, wrapUint64(math.MaxInt64-1) == math.MaxInt64-1, "")
	tu.Assert(t, wrapUint64(1023) == 1023, "")
}
