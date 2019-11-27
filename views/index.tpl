{{template "header"}}
<title>首页-我的 beego 博客</title>
</head>
<body style="padding-top: 50px;">
    <vav class="navbar navbar-default navbar-fixed-top">
        <div class="container">
           {{template "navbar".}}
        </div>
    </vav>
    <div class="container">
        <div class="col-md-9">
            {{range .Topics}}
                <div class="page-header">
                    <h1><a href="/topic/view/{{.Id}}">{{.Title}}</a></h1>
                    <h6 class="text-muted">文章发表于 {{.Created}},共有 {{.Views}} 次浏览,{{.ReplyCount}}个评论</h6>
                    <p>
                        {{.Content}}
                    </p>
                </div>
            {{end}}
        </div>

        <div class="col-md-3">
            <h3>文章分类</h3>
            <ul class="list-group">
               {{range .Categories}}
                   <li class="list-group-item"><a href="/?cate={{.Title}}">{{.Title}}</a></li>
               {{end}}
            </ul>
        </div>
    </div>
</body>
</html>
