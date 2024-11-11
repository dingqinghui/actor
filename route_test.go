/**
 * @Author: dingQingHui
 * @Description:
 * @File: route_test
 * @Version: 1.0.0
 * @Date: 2024/11/5 15:27
 */

package actor

import "testing"

func RouteFunc(s *System) {

}

func TestRoute(t *testing.T) {
	b := RouteFunc
	funInfo, _ := GetFuncInfo(RouteFunc)
	println(funInfo, b)
	_ = funInfo
}
