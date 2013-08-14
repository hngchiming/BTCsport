{{define "nav-bar"}}
<body style="padding-top:90px;font-family:Microsoft Yahei">
<div class="navbar navbar-inverse navbar-fixed-top" style>
  <div class="navbar-inner">
    <div class="container">
      <a class="btn btn-navbar" data-toggle="collapse" data-target=".nav-collapse">
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </a>
      <a class="brand" href="/" name="top">{{.App_Name}}</a>
      <div class="nav-collapse collapse" style="font-family:Microsoft Yahei; font-size:18px;">

        <ul class="nav">
          <li class="divider-vertical"></li>
          <li class="{{if .IsFootball}}active{{end}}">
            <a href="/football">足球</a>
          </li>
          <li class="divider-vertical"></li>
          <li class="{{if .IsBasketball}}active{{end}}">
            <a href="/basketball">篮球</a>
          </li>
          <li class="divider-vertical"></li>
          <li class="{{if .IsChat}}active{{end}}">
            <a href="/chat"> <i class="icon-comments icon-white"></i>
              聊天室
            </a>
          </li>
          <li class="divider-vertical"></li>
          <li class="{{if .IsAbout}}active{{end}}">
            <a href="/about"> <i class="icon-question-sign"></i>
              关于
            </a>
          </li>
          <li class="divider-vertical"></li>
        </ul>
        <div class="btn-group pull-right">
          <a class="btn dropdown-toggle" data-toggle="dropdown" href="#">
            <i class="icon-user"></i>
            {{if .User}}{{.User.Username}}{{else}}游客{{end}}
            <span class="caret"></span>
          </a>
          {{if .User}}
          <ul class="dropdown-menu">
            <li>
              <a href="/user">
                <i class="icon-book"></i>
                账户信息
              </a>
            </li>
            <li>
              <a href="">
                <i class="icon-envelope"></i>
                联系客服
              </a>
            </li>
            <li class="divider"></li>
            <li>
              <a href="/home?logout=true" >
                <i class="icon-off"></i>
                登出
              </a>
            </li>
          </ul>
          {{else}}
          <ul class="dropdown-menu">
            <li>
              <a href="/login">
                <i class="icon-signin"></i>
                登录
              </a>
            </li>
            <li>
              <a href="/login#forget">
                <i class="icon-question-sign"></i>
                忘记密码
              </a>
            </li>
            <li>
              <a href="/login#create">
                <i class="icon-plus"></i>
                注册新用户
              </a>
            </li>
            <li>
              <a href="">
                <i class="icon-envelope"></i>
                联系客服
              </a>
            </li>
          </ul>
          {{end}}
        </div>
      </div>
    </div>
  </div>
</div>
{{end}}