package driver

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKubernetesInterfaces(t *testing.T) {
	var _ Driver = &Kubernetes{}
	var _ Configurable = &Kubernetes{}
}

var quickHashRegexp = regexp.MustCompile("^[-._a-zA-Z0-9]+$")

func TestQuickHash(t *testing.T) {
	// This is a canary to make sure we calculate sum consistently
	in := "this is a test"
	out1 := quickHash(in)
	out2 := quickHash(in)
	in3 := "other test"
	out3 := quickHash(in3)

	is := assert.New(t)
	is.Equal(out1, out2)
	is.NotEqual(out1, out3)

	// Make sure the sums pass the Kubernetes regexp
	is.True(quickHashRegexp.MatchString(out1))
	is.True(quickHashRegexp.MatchString(out3))
	is.False(quickHashRegexp.MatchString("illegal/char"))
}
