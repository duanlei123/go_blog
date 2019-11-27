{{template "header"}}
<title>文章-我的 beego 博客</title>
</head>

<body style="padding-top: 50px;">
    <vav class="navbar navbar-default navbar-fixed-top">
        <div class="container">
            {{template "navbar".}}
        </div>
    </vav>
    <div class="container">
        <h1>文章列表</h1>
        <a href="/topic/add" class="btn btn-primary">添加文章</a>
        <table class="table table-striped">
            <thead>
            <tr>
                <th>#</th>
                <th>文章名称</th>
                <th>最后更新</th>
                <th>浏览</th>
                <th>回复数</th>
                <th>最后回复</th>
                <th>操作</th>
            </tr>
            </thead>
            <tbody>
            {{range .Topics}}
                <tr>
                    <th>{{.Id}}</th>
                    <th><a href="/topic/view/{{.Id}}">{{.Title}}</a></th>
                    <th>{{.Updated}}</th>
                    <th>{{.Views}}</th>
                    <th>{{.ReplyCount}}</th>
                    <th>{{.ReplyTime}}</th>
                    <th>
                        <a href="/topic/modify?tid={{.Id}}">修改</a>
                        <a href="/topic/delete/{{.Id}}">删除</a>
                    </th>
                </tr>
            {{end}}
            </tbody>
        </table>
    </div>
</body>
</html>