# go-gin-frame
go-gin-frame for clean architecture


```
├── cmd
│   └── main.go //启动函数
├── etc
│   └── dev_conf.yaml              // 配置文件 
├── global
│   └── global.go //全局变量引用，如数据库、kafka等
├── internal
│       └── service // 业务逻辑层
│           └── xxx_service.go //业务逻辑处理类
│           └── xxx_service_test.go 
│       └── model // 数据库实体类
│           └── xxx_info.go //结构体
│       └── api // 接收外部请求的代码
│           └── xxx_api.go //路由对应的接口实现
│       └── repo // 数据库操作类，数据库CRUD
│           └── xxx_repo.go //业务逻辑处理类
│           └── xxx_repo_test.go 
│       └── router //
│           └── router.go //路由
│       └── pkg
│           └── tool //工具类
│               └── date.go //时间工具类
│               └── json.go //json 工具类
```


#### 面向接口编程
除了 models 层，层与层之间应该通过接口交互而不是实现。如果要用 service 调用 repo 层，那么应该调用 repo 的接口。那么修改底层实现的时候我们上层的基类不需要变更，只需要更换一下底层实现即可。

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
这个接口的实现类就可以根据需求变更，比如说当我们想要 mysql 来作为存储查询，那么只需要提供一个这样的基类：

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


在 service 层实现的时候就可以按照需求来将对应的 repo 实现注入即可，从而不需要改动 service 层的实现：

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
- service 层：因为 service 层依赖了 repo 层，由于它们之间是通过接口来关联，所以这里可以使用 `github.com/golang/mock/gomock` 来 mock repo 层；
- api 层：这一层依赖 service 层，并且它们之间是通过接口来关联，这里也可以使用 gomock 来 mock service 层。因为接入层用的是 gin 框架，所以还需要在单测的时候模拟发送请求；

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

3. api 层测试
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

