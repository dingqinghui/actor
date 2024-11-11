/**
 * @Author: dingQingHui
 * @Description:
 * @File: reflect_test
 * @Version: 1.0.0
 * @Date: 2024/11/11 11:32
 */

package examples

import (
	"reflect"
	"testing"
)

type baseStruct struct {
}

func (receiver baseStruct) BaseName() {

}

type testStruct struct {
	baseStruct
}

func (receiver testStruct) name() {

}

func (receiver testStruct) Name() {

}

func A(a ...int) {
	//println(a)
	//println("====================")
}

func TestReflect(t *testing.T) {
	atyp := reflect.ValueOf(A)
	paramList := []reflect.Value{reflect.ValueOf(struct {
	}{})}
	atyp.Call(paramList)
}

func BenchmarkReflect(b *testing.B) {
	b.Run("reflect", func(b *testing.B) {
		atyp := reflect.ValueOf(A)
		paramList := []reflect.Value{reflect.ValueOf(1)}
		for i := 0; i < 5; i++ {
			paramList = append(paramList, reflect.ValueOf(2))
		}
		atyp.Call(paramList)
	})

	b.Run("call", func(b *testing.B) {
		A(1, 2, 3, 4, 5)
	})
}
