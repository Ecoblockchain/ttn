// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package device

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/redis.v3"

	"github.com/TheThingsNetwork/ttn/core/types"
	. "github.com/smartystreets/assertions"
)

func getRedisClient() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost"
	}
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", host),
		Password: "", // no password set
		DB:       1,  // use default DB
	})
}

func TestDeviceStore(t *testing.T) {
	a := New(t)

	stores := map[string]Store{
		"local": NewDeviceStore(),
		"redis": NewRedisDeviceStore(getRedisClient()),
	}

	for name, s := range stores {

		t.Logf("Testing %s store", name)

		// Get non-existing
		dev, err := s.Get("AppID-1", "DevID-1")
		a.So(err, ShouldNotBeNil)
		a.So(dev, ShouldBeNil)

		// Create
		err = s.Set(&Device{
			DevAddr: types.DevAddr([4]byte{0, 0, 0, 1}),
			DevEUI:  types.DevEUI([8]byte{0, 0, 0, 0, 0, 0, 0, 1}),
			AppEUI:  types.AppEUI([8]byte{0, 0, 0, 0, 0, 0, 0, 1}),
			AppID:   "AppID-1",
			DevID:   "DevID-1",
		})
		a.So(err, ShouldBeNil)

		// Get existing
		dev, err = s.Get("AppID-1", "DevID-1")
		a.So(err, ShouldBeNil)
		a.So(dev, ShouldNotBeNil)

		// Create extra
		err = s.Set(&Device{
			DevAddr: types.DevAddr([4]byte{0, 0, 0, 2}),
			DevEUI:  types.DevEUI([8]byte{0, 0, 0, 0, 0, 0, 0, 2}),
			AppEUI:  types.AppEUI([8]byte{0, 0, 0, 0, 0, 0, 0, 1}),
			AppID:   "AppID-1",
			DevID:   "DevID-2",
		})
		a.So(err, ShouldBeNil)

		// List
		devices, err := s.List()
		a.So(err, ShouldBeNil)
		a.So(devices, ShouldHaveLength, 2)

		// Delete
		err = s.Delete("AppID-1", "DevID-1")
		a.So(err, ShouldBeNil)

		// Get deleted
		dev, err = s.Get("AppID-1", "DevID-1")
		a.So(err, ShouldNotBeNil)
		a.So(dev, ShouldBeNil)

		// Cleanup
		s.Delete("AppID-1", "DevID-2")
	}

}
