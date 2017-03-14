{{define "nav"}}
<header>
	<nav class="navbar navbar-default navbar-fixed-top">
	  <div class="container">
	    <ul class="nav navbar-nav">
        <li class="active"><a href="/">每日一图 <span class="sr-only">(current)</span></a></li>
      </ul>
       <ul class="nav navbar-nav" style="float: right;"> 
        <li class="/about"><a href="/about">About<span class="sr-only">(current)</span></a></li>   
         {{if .IsLogin}}
         <li class="#"><a href="/my/photo">上传图片<span class="sr-only">(current)</span></a></li>
         <li class="#"><a href="/my">{{.User.UserName}}<span class="sr-only">(current)</span></a></li>
         {{else}}
          <li class="#"><a href="/login">登录<span class="sr-only">(current)</span></a></li>
         <li class="#"><a href="/register">注册<span class="sr-only">(current)</span></a></li>
         {{end}}
      </ul>
	  </div>
	</nav>
</header>
{{end}}