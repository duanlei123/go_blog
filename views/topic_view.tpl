{{template "header"}}
<title>{{.Topic.Title}}-我的 beego 博客</title>
</head>

<body style="padding-top: 50px;">
    <vav class="navbar navbar-default navbar-fixed-top">
        <div class="container">
            {{template "navbar".}}
        </div>
    </vav>
    <div class="container">
        {{$label := .Labels}}
        {{with .Topic}}
            <h1>{{.Title}}<small>   {{.Category}}</small></h1>
            <h5>
                {{range $label}}
                    <a href="/?label={{.}}">{{.}}</a>
                {{end}}
            </h5>
            <p>{{.Content}}</p>
            <h5>文章附件:<a href="/attachment/{{.Attachment}}">{{.Attachment}}</a></h5>
        {{end}}
    </div>

    <div class="container">
        {{$tid := .Topic.Id}}
        {{$isLogin := .IsLogin}}
        {{range .Replies}}
            <h4>{{.Name}} <small>{{.Created}}</small> {{if $isLogin}}<a href="/reply/delete?tid={{$tid}}&rid={{.Id}}">删除</a>{{end}}</h4>
            {{.Content}}
        {{end}}
        <h3>本文回复</h3>
        <form method="post" action="/reply/add">
            <input type="hidden" name="tid" value="{{.Topic.Id}}">
            <div class="form-group">
                <label>昵称:</label>
                <input type="text" class="form-control" name="nickname">
            </div>

            <div class="form-group">
                <label>内容:</label>
                <textarea name="content" id="" cols="30" rows="10" class="form-control"></textarea>
            </div>
            <button class="btn btn-default" type="submit">提交</button>
        </form>
    </div>
</body>
</html>