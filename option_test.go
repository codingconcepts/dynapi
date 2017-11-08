package dynapi

import (
	"testing"

	"github.com/codingconcepts/dynapi/test"
	"github.com/facebookgo/clock"
)

func TestOptionCertsDir(t *testing.T) {
	exp := `/etc/certs`
	s := NewServer("", 0, CertsDir(exp))
	test.Equals(t, exp, s.certsDir)
}

func TestOptionSSL(t *testing.T) {
	exp := true
	s := NewServer("", 0, SSL(exp))
	test.Equals(t, exp, s.ssl)
}

func TestOptionClock(t *testing.T) {
	exp := clock.NewMock()
	s := NewServer("", 0, Clock(exp))

	_, ok := s.clock.(*clock.Mock)
	test.Assert(t, ok)
}

func TestOptionBuildInfo(t *testing.T) {
	expVersion := "1.2.3"
	expTimestamp := "2017-11-08 07:38:46"
	s := NewServer("", 0, BuildInfo(expVersion, expTimestamp))
	test.Equals(t, expVersion, s.buildVersion)
	test.Equals(t, expTimestamp, s.buildTimestamp)
}

func TestOptionRoutes(t *testing.T) {
	exp := RouteConfig{
		URI: "expected route",
	}
	s := NewServer("", 0, Routes(exp))
	test.Equals(t, exp.URI, s.routes[0].URI)
}
