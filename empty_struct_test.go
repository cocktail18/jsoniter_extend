package jsoniter_extend

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_tolerateEmptyStructDecoder_Decode(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	//  treat {} as empty array
	jsoniter.RegisterExtension(&tolerateEmptyStructExtension{})
	var testList TestList
	var err error
	err = json.Unmarshal([]byte(`{"list":[]}`), &testList)
	assert.Nil(t, err)
	assert.True(t, len(testList.List) == 0)
	err = json.Unmarshal([]byte(`{"list":{}}`), &testList)
	assert.Nil(t, err)
	assert.True(t, len(testList.List) == 0)

	err = json.Unmarshal([]byte(`{"list":[{"id":1}]}`), &testList)
	assert.Nil(t, err)
	assert.True(t, len(testList.List) == 1)
	assert.True(t, testList.List[0].Id == 1)
}

type TestList struct {
	List []TestData `json:"list"`
}

type TestData struct {
	Id int `json:"id"`
}
