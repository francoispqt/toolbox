package scp_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/toolbox/storage/scp"
	"net/url"
	"testing"
)

// -rw-r--r--  1 awitas  wheel  6668 Oct 25 11:41:44 2017 /build/dcm/dcm-server/target/classes/dcm.properties

func Test_ExtractFileInfo(t *testing.T) {

	var parserURL, _ = url.Parse("scp://127.0.0.1/")

	{
		var parser = scp.Parser{IsoTimeStyle: false}
		objects, err := parser.Parse(parserURL, "-rw-r--r--  1 awitas  1742120565   414 Jun  8 14:14:08 2017 f.pub", true)
		if assert.Nil(t, err) {
			var object = objects[0]
			assert.Equal(t, "scp://127.0.0.1/f.pub", object.URL())
			assert.Equal(t, int64(1496931248), object.FileInfo().ModTime().Unix())
			assert.Equal(t, int64(414), object.FileInfo().Size())
			assert.Equal(t, true, object.IsContent())
		}
	}
	{
		parserURL, _ = url.Parse("scp://127.0.0.1:22/")
		var parser = scp.Parser{IsoTimeStyle: true}
		var objects, err = parser.Parse(parserURL, "-rw-r--r-- 1 awitas awitas 2002 2017-11-04 22:29:33.363458941 +0000 aerospikeciads_aerospike.conf", true)
		if assert.Nil(t, err) {
			var object = objects[0]
			assert.Equal(t, "scp://127.0.0.1:22/aerospikeciads_aerospike.conf", object.URL())
			assert.Equal(t, int64(1509834573), object.FileInfo().ModTime().Unix())
			assert.Equal(t, int64(2002), object.FileInfo().Size())
			assert.Equal(t, true, object.IsContent())
		}
	}

}
