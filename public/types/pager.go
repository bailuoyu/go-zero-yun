package types

import "math"

// GetPager 获取分页
func GetPager(req PageReq) Pager {
	var p Pager
	p.Set(req)
	return p
}

// Set 设置分页
func (p *Pager) Set(req PageReq) {
	p.Page = req.Page
	p.PageSize = req.PageSize
	p.checkPage()
}

// checkPage 检查分页
func (p *Pager) checkPage() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	} else if p.PageSize > 1000 {
		p.PageSize = 1000
	}
}

// Count 设置总数
func (p *Pager) Count(total int64) {
	p.Total = int(total)
	p.PageCount = int(math.Ceil(float64(total) / float64(p.PageSize)))
}

// Offset 偏移量
func (p *Pager) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// Limit 限制条数
func (p *Pager) Limit() int {
	return p.PageSize
}
