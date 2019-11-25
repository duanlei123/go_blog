{{template "header"}}
<title>分类-我的 beego 博客</title>
</head>

<body style="padding-top: 50px;">
    <vav class="navbar navbar-default navbar-fixed-top">
        <div class="container">
            {{template "navbar".}}
        </div>
    </vav>

    <div class="container">
        <h1>分类列表</h1>
        <form method="post" action="/category">
            <div class="form-group">
                <label>分类名称</label>
                <input id="name" type="text" class="form-control" placeholder="例如:Goland入门第一篇" name="name">
            </div>
            <input type="hidden" name="op" value="add">
            <button type="submit" class="btn btn-default" onclick="return checkInput()">添加</button>
        </form>
        <script type="text/javascript">
            function checkInput() {
                var name = document.getElementById("name")
                if (name.value.length == 0){
                    alert("请输入分类名称")
                    return false
                }
                return true
            }
        </script>
        <table class="table table-striped">
            <thead>
                <tr>
                    <th>#</th>
                    <th>名称</th>
                    <th>文章数</th>
                    <th>操作</th>
                </tr>
            </thead>
            <tbody>
                {{range .Categories}}
                <tr>
                    <th>{{.Id}}</th>
                    <th>{{.Title}}</th>
                    <th>{{.TopicCount}}</th>
                    <th>
                        <a href="/category?op=del&id={{.Id}}">删除</a>
                    </th>
                </tr>
                {{end}}
            </tbody>
        </table>
    </div>
</body>
</html>