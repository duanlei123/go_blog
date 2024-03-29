{{template "header"}}
<title>添加文章-我的 beego 博客</title>
</head>

<body style="padding-top: 50px;">
    <vav class="navbar navbar-default navbar-fixed-top">
        <div class="container">
            {{template "navbar".}}
        </div>
    </vav>
    <div class="container">
        <h1>添加文章</h1>
        <form method="post" action="/topic" enctype="multipart/form-data">
            <div class="form-group">
                <label>文章标题:</label>
                <input type="text" name="title" class="form-control">
            </div>
            <div class="form-group">
                <label>文章分类:</label>
                <input type="text" name="category" class="form-control">
            </div>
            <div class="form-group">
                <label>文章标签:</label>
                <input type="text" name="labels" class="form-control">
            </div>
            <div class="form-group">
                <label>文章内容:</label>
                <textarea name="content" cols="30" rows="10" class="form-control"></textarea>
            </div>
            <div class="form-group">
                <label>文章附件:</label>
                <input type="file" name="attachment" class="form-control">
            </div>
            <button type="submit" class="btn btn-default">添加文章</button>
        </form>
    </div>
</body>
</html>