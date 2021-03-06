package toolbox_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/toolbox"
	"testing"
)

func TestNewFileSetInfoInfo(t *testing.T) {

	fileSetInfo, err := toolbox.NewFileSetInfo("./fileset_info_test/")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 2, len(fileSetInfo.FilesInfo()))

	fileInfo := fileSetInfo.FileInfo("user_test.go")
	assert.NotNil(t, fileInfo)

	addresses := fileSetInfo.Type("Addresses")
	assert.NotNil(t, addresses)

	assert.False(t, fileInfo.HasType("F"))
	assert.True(t, fileInfo.HasType("User"))

	assert.Equal(t, 7, len(fileInfo.Types()))

	address := fileSetInfo.Type("Address")
	assert.NotNil(t, address)

	assert.Equal(t, 2, len(address.Fields()))
	country := address.Field("Country")
	assert.NotNil(t, country)
	assert.True(t, country.IsAnonymous)

	z := fileSetInfo.Type("Z")
	assert.NotNil(t, z)

	address2 := fileSetInfo.Type("Address2")
	assert.Nil(t, address2)

	userInfo := fileInfo.Type("User")
	assert.NotNil(t, userInfo)

	assert.True(t, userInfo.HasField("ID"))
	assert.True(t, userInfo.HasField("Name"))
	assert.False(t, userInfo.HasField("FF"))

	assert.Equal(t, 10, len(userInfo.Fields()))

	idInfo := userInfo.Field("ID")
	assert.True(t, idInfo.IsPointer)
	assert.Equal(t, "int", idInfo.TypeName)
	assert.Equal(t, true, idInfo.IsPointer)

	dobInfo := userInfo.Field("DateOfBirth")

	assert.Equal(t, "time.Time", dobInfo.TypeName)
	assert.Equal(t, "time", dobInfo.TypePackage)

	assert.Equal(t, "`foo=\"bar\"`", dobInfo.Tag)

	addressPointer := userInfo.Field("AddressPointer")
	assert.NotNil(t, addressPointer)
	assert.Equal(t, "Address", addressPointer.TypeName)

	cInfo := userInfo.Field("C")
	assert.True(t, cInfo.IsChannel)

	mInfo := userInfo.Field("M")
	assert.True(t, mInfo.IsMap)
	assert.Equal(t, "string", mInfo.KeyTypeName)
	assert.Equal(t, "[]string", mInfo.ValueTypeName)

	intsInfo := userInfo.Field("Ints")
	assert.True(t, intsInfo.IsSlice)
	assert.Equal(t, "my comments", userInfo.Comment)

	assert.False(t, userInfo.HasReceiver("Abc"))

	assert.Equal(t, 3, len(userInfo.Receivers()))
	assert.True(t, userInfo.HasReceiver("Test"))
	assert.True(t, userInfo.HasReceiver("Test2"))

	receiver := userInfo.Receiver("Test")
	assert.NotNil(t, receiver)

}
