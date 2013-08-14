{{template "admin" .}}
{{if .IsAdmin}}
<div class="container">
	<table class="table table-bordered table-hover table-striped table-condensed ">
		<thead>
			<tr>
				<th>Id</th>
				<th>用户名</th>
				<th>邮箱</th>
				<th>生日</th>
				<th>地址</th>
				<th>IP</th>
			</tr>
		</thead>
		<tbody>
			{{range .UserAll}}
			<tr>
				<td>{{.Id}}</td>
				<td>{{.Username}}</td>
				<td>{{.Email}}</td>
				<td>{{.Birth}}</td>
				<td>{{.Btcaddress}}</td>
				<td>{{.Lastip}}</td>
			</tr>
			{{end}}
		</tbody>
	</table>
</div>
{{else}}
<div class="container">
	<div class="hero-unit"> <b>YOU ARE NOT ALLOWED HERE</b>
	</div>
</div>
{{end}}

{{template "footer" .}}