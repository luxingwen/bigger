{{define "user"}}
<a href="#">
		 <img src="{{.User.Pic}}" alt="" width="100%"></a>
		<h3>{{.User.UserName}}</h3>
		<hr />
		<span>{{.User.Email}}</span>
{{end}}