{{template "admin" .}}
{{if .IsAdmin}}
<div class="container">
	<table class="table table-bordered table-hover table-striped table-condensed ">
		<thead>
			<tr>
				<th>Id</th>
				<th>用户名</th>
				<th>余额</th>
				<th>盈利</th>
				<th>总押注</th>
				<th>推荐人</th>
			</tr>
		</thead>
		<tbody>
			{{range .UserAll}}
			<tr>
				<td>{{.Id}}</td>
				<td>{{.Username}}</td>
				<td>{{.Balance}}</td>
				<td>{{.Profit}}</td>
				<td>{{.Alltimebet}}</td>
				<td>{{.Referral}}</td>
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