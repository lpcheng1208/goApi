package utils

import "math"

type Page struct {
	PageNo     int         `json:"page_num"`
	PageSize   int         `json:"page_size"`
	TotalPage  int         `json:"total_page"`
	TotalCount int         `json:"total_count"`
	FirstPage  bool        `json:"first_page"`
	LastPage   bool        `json:"last_page"`
	List       interface{} `json:"list"`
}

func PageUtil(count int, pageNo int, pageSize int, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
}

//分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func Paginator(page, pageSize int, nums int64, list interface{}) map[string]interface{} {

	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(pageSize))) //page总数
	nextpage := true
	if page > totalpages {
		nextpage = false
	}
	if page <= 0 {
		page = 1
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["next_page"] = nextpage
	paginatorMap["currpage"] = page
	paginatorMap["list"] = list
	return paginatorMap
}
