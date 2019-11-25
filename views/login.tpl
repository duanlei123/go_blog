{{template "header"}}
<title>登录-我的 beego 博客</title>
</head>
<body style="padding-top: 200px;">
    <div class="container" style="width: 500px">
        <form method="post" action="/login">
            <div class="form-group">
                <label>Account</label>
                <input id="uname" type="text" class="form-control" placeholder="Enter Account" name="uname">
            </div>
            <div class="form-group">
                <label >Password</label>
                <input id="pwd" type="password" class="form-control" placeholder="Password" name="pwd">
            </div>
            <div class="checkbox">
                <label>
                    <input type="checkbox" name="autoLogin"> 自动登录
                </label>
            </div>
            <button type="submit" class="btn btn-default" onclick="return checkInput()">登录</button>
            <button class="btn btn-default" onclick="return backToHome()">返回</button>
        </form>
        <script type="text/javascript">
            function checkInput() {
                var unam = document.getElementById("uname")
                if (unam.value.length == 0){
                    alert("请输入账号")
                    return false
                }
                var pwd = document.getElementById("pwd")
                if (pwd.value.length == 0){
                    alert("请输入密码")
                    return false
                }
                return true
            }

            function backToHome() {
                window.location.href = "/"
                return false
            }

        </script>
    </div>
</body>
</html>