package test

import (
	"testing"

	"gotest.tools/assert"
	"k8s.io/klog"
)

func TestEqual(t *testing.T) {
	var a = 100
	//var b = 200

	klog.Infof("hello, infof")
	klog.Infof("hello, infof2")
	klog.Errorf("hello, errorf")

	assert.Check(t, a == 100, "a should be 100")
	//assert.Equal(t, a, b, "they should be equal")
	//assert.Assert(t, Equal(a, b), "sss")

}
