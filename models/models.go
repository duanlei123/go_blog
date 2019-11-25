package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	_DB_NAME = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

// 分类
type Category struct {
	Id int64
	Title string
	Created time.Time `orm:"index"`
	Views int64 `orm:"index"`
	TopicTime time.Time `orm:"index"`
	TopicCount int64
	TopicLastUserId int64
}

// 文章
type Topic struct {
	Id int64
	Uid int64
	Title string
	Content string `orm:"size(5000)"`
	Attachment string
	Created time.Time `orm:"index"`
	Updated time.Time `orm:"index"`
	Views int64 `orm:"index"`
	Author string
	ReplyTime time.Time `orm:"index"`
	ReplyCount int64
	ReplyLastUserId int64

}

func RegisterDB()  {
	if !com.IsExist(_DB_NAME) {
		_ = os.MkdirAll(path.Dir(_DB_NAME), os.ModeAppend)
		_, _ = os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic))
	_ = orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	_ = orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

// 新增分类
func AddCategory(name string) error{
	o := orm.NewOrm()
	cate := &Category{Title: name}
	//查询操作
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil{
		return err
	}
	_, err = o.Insert(cate)
	if err != nil{
		// 插入失败
		return err
	}
	return nil
}

//获取所有分类
func GetAllCategory() ([]*Category, error){
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

// 删除分类
func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}