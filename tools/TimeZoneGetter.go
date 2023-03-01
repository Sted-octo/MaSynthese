package tools

import "time"

func TimeZoneGetter(zone string) *time.Location {
	t, _ := time.LoadLocation(zone)
	return t
}
