package udf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Length(t *testing.T) {
	{
		value, err := Length(4.3, nil)
		assert.Nil(t, err)
		assert.Equal(t, 0, value)
	}
	{
		value, err := Length("abcd", nil)
		assert.Nil(t, err)
		assert.Equal(t, 4, value)
	}
	{
		value, err := Length(map[int]int{
			2: 3,
			1: 1,
			6: 3,
		}, nil)
		assert.Nil(t, err)
		assert.Equal(t, 3, value)
	}
	{
		value, err := Length([]int{1, 2, 3}, nil)
		assert.Nil(t, err)
		assert.Equal(t, 3, value)
	}
}

func Test_Keys(t *testing.T) {

	{
		var keys, err = Keys(map[string]interface{}{}, nil)
		assert.Nil(t, err)
		assert.NotNil(t, keys)
	}
	{
		var keys, err = Keys("{\"abc\":1}", nil)
		assert.Nil(t, err)
		assert.EqualValues(t, []interface{}{"abc"}, keys)
	}
}

func Test_Values(t *testing.T) {

	{
		var keys, err = Values(map[string]interface{}{}, nil)
		assert.Nil(t, err)
		assert.NotNil(t, keys)
	}
	{
		var keys, err = Values("{\"abc\":1}", nil)
		assert.Nil(t, err)
		assert.EqualValues(t, []interface{}{1}, keys)
	}
}

func Test_Join(t *testing.T) {
	{
		var joined, err = Join([]interface{}{
			[]interface{}{1, 2, 3},
			",",
		}, nil)
		assert.Nil(t, err)
		assert.NotNil(t, joined)
		assert.EqualValues(t, "1,2,3", joined)
	}
}

func Test_IndexOf(t *testing.T) {
	{
		index, _ := IndexOf([]interface{}{"this is test", "is"}, nil)
		assert.EqualValues(t, 2, index)
	}
	{
		index, _ := IndexOf([]interface{}{[]string{"this", "is", "test"}, "is"}, nil)
		assert.EqualValues(t, 1, index)
	}
}

func Test_Base64Decode(t *testing.T) {
	decoded, _ := Base64DecodeText("IkhlbGxvIFdvcmxkIg==", nil)
	assert.EqualValues(t, `"Hello World"`, decoded)

}

func TestTrimSpace(t *testing.T) {
	trimmed, _ := TrimSpace(" erer ", nil)
	assert.EqualValues(t, `erer`, trimmed)

}
