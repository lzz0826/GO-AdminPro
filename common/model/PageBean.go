package model

type PageBean struct {
	Total      int64       `json:"total"`      // 总记录数
	Pages      int         `json:"pages"`      // 总页数
	IsLastPage bool        `json:"isLastPage"` // 是否最后一页
	BeanList   interface{} `json:"beanList"`   // 数据列表
}

// Set 用于设置PageBean的属性，并返回PageBean实例
func (pb *PageBean) Set(total int64, page, size int, beanList interface{}) *PageBean {
	pb.Total = total
	// 計算總頁數
	if size > 0 {
		pageCount := (total + int64(size) - 1) / int64(size)
		pb.Pages = int(pageCount)
		if page <= 0 {
			page = 1
		}
		pb.IsLastPage = page >= pb.Pages
	}
	pb.BeanList = beanList
	return pb
}

// Empty 返回一个空的PageBean实例
func Empty() *PageBean {
	return &PageBean{
		Total:      0,
		Pages:      0,
		IsLastPage: true,
		BeanList:   []interface{}{},
	}
}

/**
 * 生成PageBean
 * @param total    总记录数
 * @param page     當前頁
 * @param size     頁大小
 * @param  interface{} 范行
 */
// Of 根据传入的参数创建并返回一个PageBean实例
func Of(total int64, page, size int, beanList interface{}) *PageBean {
	pb := &PageBean{}
	return pb.Set(total, page, size, beanList)
}
