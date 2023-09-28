# go-gin-frame
go-gin-frame for clean architecture


```
├── assets
│   └── assets.go //embed 静态文件
├── config
│   │── dev_config.yaml              // 本地开发环境配置文件 
│   │── fat_config.yaml              // 测试环境配置文件 
│   │── pro_config.yaml              // 正式环境配置文件 
│   └── uat_config.yaml              // 预上线环境配置文件 
├── constants
│   └── constants.go //全局变量引用，如数据库、kafka等
├── internal
│       │── code // 自定义业务代码
│       │   └── code.go
│       │── service // 业务逻辑层
│       │   └── article
|       |       └── create.go //业务逻辑处理类
|       |       └── create_test.go
|       |       └── detail.go
|       |       └── delete.go
│       │── model 
│       │   └── article // 数据库实体
|       |       └── article.go //实体类
|       |       └── model.go
|       |       └── create.go //业务逻辑处理
|       |       └── create_test.go
|       |       └── detail.go
|       |       └── detail_test.go
│       │── controller // 接收外部请求的代码
│       │   └── article
|       |       └── create.go //路由对应的接口实现
|       |       └── detail.go
|       |       └── delete.go
│       │── router
│       │   └── router.go //路由
│       └── socket
├── pkg      // 扩展方法工具类       
│   │── env
│   │── errors
│   │── file
│   └── timeutil
├── initial.go // 初始化方法
```


#### 面向接口编程
除了 model 层，层与层之间应该通过接口交互而不是实现。

调用过程: controller层 --> service层 --> model层

例如想要将所有文章查询出来，那么可以在 repo 提供这样的接口：

```
package repo

import (
    "context"
    "my-clean-rchitecture/models"
    "time"
)

// IArticleRepo represent the article's repository contract
type IArticleRepo interface {
    Fetch(ctx context.Context, createdDate time.Time, num int) (res []models.Article, err error)
}

```
接口的实现类就可以根据需求变更，想要 mysql 来作为存储查询，只需要提供一个 mysqlArticleRepository 基类：

```
type mysqlArticleRepository struct {
    DB *gorm.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlArticleRepository(DB *gorm.DB) IArticleRepo {
    return &mysqlArticleRepository{DB}
}

func (m *mysqlArticleRepository) Fetch(ctx context.Context, createdDate time.Time,
    num int) (res []models.Article, err error) {

    err = m.DB.WithContext(ctx).Model(&models.Article{}).
        Select("id,title,content, updated_at, created_at").
        Where("created_at > ?", createdDate).Limit(num).Find(&res).Error
    return
}
```

如果想要换成 MongoDB 来实现存储，那么只需要定义一个新的结构体 mongoArticleRepository 实现 IArticleRepo 接口即可。

如果要用 service 调用 repo 层，那么应该调用 repo 的接口。在 service 层实现的时候就可以按照需求来将对应的 repo 实现注入即可，从而不需要改动 service 层的实现：

```
type articleService struct {
    articleRepo repo.IArticleRepo
}

// NewArticleService will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewArticleService(a repo.IArticleRepo) IArticleService {
    return &articleService{
        articleRepo: a,
    }
}

// Fetch
func (a *articleService) Fetch(ctx context.Context, createdDate time.Time, num int) (res []models.Article, err error) {
    if num == 0 {
        num = 10
    }
    res, err = a.articleRepo.Fetch(ctx, createdDate, num)
    if err != nil {
        return nil, err
    }
    return
}
```

#### 依赖注入 DI
wire

#### 测试

- model 层：由于没有依赖任何其他代码，所以可以直接用 go 的单测框架直接测试；
- repo 层：由于使用了 mysql 数据库，那么需要 mock mysql，这样即使不用连接 mysql 也可以正常测试，这里推荐使用 `github.com/DATA-DOG/go-sqlmock`；
- service 层：因为 service 层依赖了 model 层，由于它们之间是通过接口来关联，所以这里可以使用 `github.com/golang/mock/gomock` 来 mock model 层；
- controller 层：这一层依赖 service 层，并且它们之间是通过接口来关联，这里也可以使用 gomock 来 mock service 层。因为接入层用的是 gin 框架，所以还需要在单测的时候模拟发送请求

通过 `github.com/golang/mock/gomock` 来进行 mock ，需要执行代码生成，生成的mock 代码保存到 mock 目录
```
mockgen -destination .\mock\repo_mock.go -source .\repo\repo.go -package mock

mockgen -destination .\mock\service_mock.go -source .\service\service.go -package mock
```

1. repo 层测试

项目用了 gorm 作为 orm库，所以需要使用 github.com/DATA-DOG/go-sqlmock 结合 gorm 来进行 mock：
```
func getSqlMock() (mock sqlmock.Sqlmock, gormDB *gorm.DB) {
    //创建sqlmock
    var err error
    var db *sql.DB
    db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
    if err != nil {
        panic(err)
    }
    //结合gorm、sqlmock
    gormDB, err = gorm.Open(mysql.New(mysql.Config{
        SkipInitializeWithVersion: true,
        Conn:                      db,
    }), &gorm.Config{})
    if nil != err {
        log.Fatalf("Init DB with sqlmock failed, err %v", err)
    }
    return
}

func Test_mysqlArticleRepository_Fetch(t *testing.T) {
    createAt := time.Now()
    updateAt := time.Now()
    //id,title,content, updated_at, created_at
    var articles = []models.Article{
        {1, "test1", "content", updateAt, createAt},
        {2, "test2", "content2", updateAt, createAt},
    }

    limit := 2
    mock, db := getSqlMock()

    mock.ExpectQuery("SELECT id,title,content, updated_at, created_at FROM `articles` WHERE created_at > ? LIMIT 2").
        WithArgs(createAt).
        WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "updated_at", "created_at"}).
            AddRow(articles[0].ID, articles[0].Title, articles[0].Content, articles[0].UpdatedAt, articles[0].CreatedAt).
            AddRow(articles[1].ID, articles[1].Title, articles[1].Content, articles[1].UpdatedAt, articles[1].CreatedAt))

    repository := NewMysqlArticleRepository(db)
    result, err := repository.Fetch(context.TODO(), createAt, limit)

    assert.Nil(t, err)
    assert.Equal(t, articles, result)
}
```

2. service 层测试
```
func Test_articleService_Fetch(t *testing.T) {
    ctl := gomock.NewController(t)
    defer ctl.Finish()
    now := time.Now()
    mockRepo := mock.NewMockIArticleRepo(ctl)

    gomock.InOrder(
        mockRepo.EXPECT().Fetch(context.TODO(), now, 10).Return(nil, nil),
    )

    service := NewArticleService(mockRepo)

    fetch, _ := service.Fetch(context.TODO(), now, 10)
    fmt.Println(fetch)
}
```

3. controller 层测试
不仅要 mock service 层，还需要发送 httptest 来模拟请求发送：
```
func TestArticleHandler_FetchArticle(t *testing.T) {

    ctl := gomock.NewController(t)
    defer ctl.Finish()
    createAt, _ := time.Parse("2006-01-02", "2021-12-26")
    mockService := mock.NewMockIArticleService(ctl)

    gomock.InOrder(
        mockService.EXPECT().Fetch(gomock.Any(), createAt, 10).Return(nil, nil),
    )

    article := NewArticleHandler(mockService)

    gin.SetMode(gin.TestMode)

    // Setup your router, just like you did in your main function, and
    // register your routes
    r := gin.Default()
    r.GET("/articles", article.FetchArticle)

    req, err := http.NewRequest(http.MethodGet, "/articles?num=10&create_date=2021-12-26", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    w := httptest.NewRecorder()
    // Perform the request
    r.ServeHTTP(w, req)

    // Check to see if the response was what you expected
    if w.Code != http.StatusOK {
        t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
    }
}

```

