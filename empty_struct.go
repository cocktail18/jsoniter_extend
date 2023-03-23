package jsoniter_extend

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
	"reflect"
	"unsafe"
)

type tolerateEmptyStructExtension struct {
	jsoniter.DummyExtension
}

func (extension *tolerateEmptyStructExtension) DecorateDecoder(typ reflect2.Type, decoder jsoniter.ValDecoder) jsoniter.ValDecoder {
	if typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array {
		return &tolerateEmptyStructDecoder{decoder}
	}
	return decoder
}

type tolerateEmptyStructDecoder struct {
	valDecoder jsoniter.ValDecoder
}

func (decoder *tolerateEmptyStructDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	if iter.WhatIsNext() == jsoniter.ObjectValue {
		iter.Skip()
		newIter := iter.Pool().BorrowIterator([]byte("[]"))
		defer iter.Pool().ReturnIterator(newIter)
		decoder.valDecoder.Decode(ptr, newIter)
	} else {
		decoder.valDecoder.Decode(ptr, iter)
	}
}
