package main

import (
	"os"
	"testing"

	dns "github.com/cert-manager/cert-manager/test/acme"
)

var (
	zone               = os.Getenv("TEST_ZONE_NAME")
	kubeBuilderBinPath = "./_out/kubebuilder/bin"
)

func TestRunsSuite(t *testing.T) {
	// The manifest path should contain a file named config.json that is a
	// snippet of valid configuration that should be included on the
	// ChallengeRequest passed as part of the test cases.

	fixture := dns.NewFixture(&customDNSProviderSolver{},
		dns.SetManifestPath(kubeBuilderBinPath),
		dns.SetResolvedZone(zone),
		dns.SetAllowAmbientCredentials(false),
		dns.SetManifestPath("testdata/cpanel"),
		dns.SetStrict(true),
	)

	fixture.RunConformance(t)
}
