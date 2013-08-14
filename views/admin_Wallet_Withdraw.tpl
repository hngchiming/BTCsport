{{template "admin" .}}
{{if .IsAdmin}}
<div class="container">
	<table class="table table-bordered table-hover table-striped table-condensed ">
		<thead>
			<tr>
				<th>用户Id</th>
				<th>数额</th>
				<th>提现地址</th>
				<th>日期</th>
			</tr>
		</thead>
		<tbody>
			{{range .Deposit}}
			<tr>
				<td>{{.Uid}}</td>
				<td>{{.Amount}}</td>
				<td>{{.Address}}</td>
				<td>{{.Time}}</td>
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