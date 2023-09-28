package article

import (
	"go-gin-frame/internal/code"
	"go-gin-frame/internal/service/article"
	"go-gin-frame/pkg/hash"
	"go-gin-frame/pkg/timeutil"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type listRequest struct {
	Page     int    `json:"page"`      // 第几页
	PageSize int    `json:"page_size"` // 每页显示条数
	Title    string `json:"title"`     // 标题
}

type listData struct {
	Id        int    `json:"id"`         // ID
	HashID    string `json:"hashid"`     // hashid
	Title     string `json:"title"`      // 标题
	Content   string `json:"content"`    // 内容
	CreatedAt string `json:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at"` // 更新时间
}

type listResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int `json:"total"`
		CurrentPage  int `json:"current_page"`
		PerPageCount int `json:"per_page_count"`
	} `json:"pagination"`
}

// 文章列表
func List(c *gin.Context) {
	req := new(listRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(200, gin.H{
			"code":    code.ParamBindError,
			"message": code.Text(code.ParamBindError),
		})
		return
	}

	page := req.Page
	if page == 0 {
		page = 1
	}

	pageSize := req.PageSize
	if pageSize == 0 {
		pageSize = 10
	}

	searchData := new(article.ListSearchData)
	searchData.Page = req.Page
	searchData.PageSize = req.PageSize
	searchData.Title = req.Title

	s := article.NewService()
	list, err := s.PageList(searchData)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    code.ArticleListError,
			"message": code.Text(code.ArticleListError),
		})
	}
	count, err := s.PageCount(searchData)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    code.ArticleListError,
			"message": code.Text(code.ArticleListError),
		})
	}

	resp := new(listResponse)
	resp.Pagination.Total = int(count)
	resp.Pagination.CurrentPage = page
	resp.Pagination.PerPageCount = pageSize
	resp.List = make([]listData, len(list))

	for k, v := range list {
		hashId, err := hash.HashidsEncode([]int{cast.ToInt(v.ID)})
		if err != nil {
			c.JSON(200, gin.H{
				"code":    code.ArticleListError,
				"message": code.Text(code.ArticleListError),
			})
			return
		}

		data := listData{
			Id:        int(v.ID),
			HashID:    hashId,
			Title:     v.Title,
			Content:   v.Content,
			CreatedAt: v.CreatedAt.Format(timeutil.CSTLayout),
			UpdatedAt: v.UpdatedAt.Format(timeutil.CSTLayout),
		}

		resp.List[k] = data
	}

	c.JSON(200, gin.H{
		"code":    200,
		"data":    resp,
		"message": "List",
	})
}
