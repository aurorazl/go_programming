package utils

import "reflect"

func ConverseInterfaceToSliceInterface(in interface{}) []interface{} {
	var out []interface{}
	switch reflect.TypeOf(in).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(in)
		for i := 0; i < s.Len(); i++ {
			out = append(out, s.Index(i).Interface())
		}
	}
	return out
}

func GetTwoSliceDiff(srcIn, dstIn interface{}) ([]interface{}, []interface{}) {
	onlyInSrc, onlyInDst := make([]interface{}, 0), make([]interface{}, 0)
	src := ConverseInterfaceToSliceInterface(srcIn)
	dst := ConverseInterfaceToSliceInterface(dstIn)
	if src == nil && dst == nil {
		return nil, nil
	}
	srcMap := make(map[interface{}]bool, len(src))
	for _, oneKey := range src {
		srcMap[oneKey] = true
	}
	for _, dstKey := range dst {
		_, ok := srcMap[dstKey]
		if ok {
			srcMap[dstKey] = false
		} else {
			onlyInDst = append(onlyInDst, dstKey)
		}
	}
	for imei, exist := range srcMap {
		if exist == true {
			onlyInSrc = append(onlyInSrc, imei)
		}
	}
	return onlyInSrc, onlyInDst
}
