package handler

import (
	_ "github.com/parvez3019/goas/model"
)

// @Title Get restaurants list
// @Description Returns a list of restaurants based on filter request
// @Header model.Headers
// @Param count query int32 false "count of restaurants"
// @Param offset query int32 false "offset limit count"
// @Param order_by query model.OrderByEnum false "order restaurants list"
// @Param filter query model.Filter false "In json format"
// @Success 200 {object} model.GetRestaurantsResponse
// @Router /restaurants [get]
func GetRestaurants() {
}
