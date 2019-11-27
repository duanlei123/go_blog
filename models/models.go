package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"strings"
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
	Category string
	Labels string
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
// 评论
type Comment struct {
	Id int64
	Tid int64
	Name string
	Content string `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

func RegisterDB()  {
	if !com.IsExist(_DB_NAME) {
		_ = os.MkdirAll(path.Dir(_DB_NAME), os.ModeAppend)
		_, _ = os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
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

// 添加文章
func AddTopic(title,category, label, content string) error{
	// 处理标签
	label = "$" + strings.Join(strings.Split(label," "), "#$") + "#"

	o := orm.NewOrm()
	topic := &Topic{
		Title:           title,
		Category:		 category,
		Labels:			 label,
		Content:         content,
		Attachment:      "",
		Created:         time.Now(),
		Updated:         time.Now(),
		Views:           0,
		Author:          "",
		ReplyTime:       time.Now(),
		ReplyCount:      0,
		ReplyLastUserId: 0,
	}
	_, err := o.Insert(topic)
	if err != nil{
		return err
	}
	// 更新分类统计
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil{
		//如果不存在，简单忽略更新操作
		cate.TopicCount++
		_, err = o.Update(cate)
	}
	return err
}
// 获取所有文章
func GetAllTopics(cate string, label string, isDesc bool)([]*Topic, error){
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	if isDesc{
		if len(cate) > 0{
			qs = qs.Filter("category", cate)
		}
		if len(label) > 0{
			qs = qs.Filter("labels__contains","$"+label+"#")
		}
		_, err = qs.OrderBy("-created").All(&topics)
	}else {
		_, err = qs.All(&topics)
	}
	return topics, err
}
// 获取文章详情
func GetTopic(tid string) (*Topic, error){
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil{
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}
	// 更新浏览记录
	topic.Views++
	_, err = o.Update(topic)
	topic.Labels = strings.Replace(strings.Replace(topic.Labels,"#"," ",-1),"$","",-1)
	return topic,err
}

// 修改文章
func ModelsTopic(tid, title, category, label, content string) error{
	label = "$" + strings.Join(strings.Split(label," "), "#$") + "#"
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil{
		return err
	}
	var oldCate string // 文章之前的分类
	o := orm.NewOrm()
	topic := &Topic{Id:tidNum}
	if o.Read(topic) == nil{
		oldCate = topic.Category
		topic.Title = title
		topic.Content = content
		topic.Labels = label
		topic.Category = category
		topic.Updated = time.Now()
		_, err = o.Update(topic)
		if err != nil{
			return err
		}
	}
	//  更新分类统计
	if len(oldCate) > 0{ // 旧分类存在
		category := new(Category)
		qs := o.QueryTable("category")
		err := qs.Filter("title", oldCate).One(category)
		if err == nil{
			category.TopicCount--
			_, err = o.Update(category)
		}
	}
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil{
		cate.TopicCount++
		_, err = o.Update(cate)
	}
	return err
}
//删除文章
func DeleteTopic(tid string) error{
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil{
		return err
	}
	var oldCate string // 文章之前的分类
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil{
		oldCate = topic.Category
		_, err = o.Delete(topic)
		if err != nil{
			return err
		}
	}
	//更新分类统计
	if len(oldCate) > 0{ // 旧分类存在
		category := new(Category)
		qs := o.QueryTable("category")
		err := qs.Filter("title", oldCate).One(category)
		if err == nil{
			category.TopicCount--
			_, err = o.Update(category)
		}
	}
	return err
}

//添加评论
func AddReply(tid, nickname, content string) error{
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil{
		return err
	}
	//创建评论对象
	reply := &Comment{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}
	newOrm := orm.NewOrm()
	_, err = newOrm.Insert(reply)
	if err != nil{
		return err
	}
	// 更新评论统计和最后回复时间
	//1.获取文章
	topic := &Topic{Id: tidNum}
	if newOrm.Read(topic) == nil{
		topic.ReplyTime = time.Now()
		topic.ReplyCount++
		_, err = newOrm.Update(topic)
	}
	return err
}
// 获取某个文章的评论
func GetAllReplies(tid string) (replies []*Comment, err error){
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil{
		return nil, err
	}
	//创建评论对象
	replies = make([]*Comment, 0)
	newOrm := orm.NewOrm()
	qs := newOrm.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).All(&replies)
	return replies, err
}
//删除评论
func DeleteReply(rid string) error{
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil{
		return err
	}
	var tidNum int64
	comment := &Comment{Id: ridNum}
	newOrm := orm.NewOrm()
	if newOrm.Read(comment) == nil{
		tidNum = comment.Tid
		_, err = newOrm.Delete(comment)
		if err != nil{
			return err
		}
	}
	// 更新评论统计和最后更新时间
	replies := make([]*Comment, 0)
	qs := newOrm.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).OrderBy("-created").All(&replies)
	if err != nil{
		return nil
	}
	topic := &Topic{Id: tidNum}
	if newOrm.Read(topic) == nil{
		topic.ReplyTime = replies[0].Created
		topic.ReplyCount = int64(len(replies))
		_, err = newOrm.Update(topic)
	}
	return err
}