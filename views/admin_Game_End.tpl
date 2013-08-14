{{template "admin" .}}
{{if .IsAdmin}}
<div class="container">
	<table class="table table-bordered table-hover table-striped table-condensed ">
		<thead>
			<tr>
				<th>Id</th>
				<th>主队</th>
				<th>客队</th>
				<th>开始时间</th>
				<th>让分</th>
				<th>主队赔率</th>
				<th>客队赔率</th>
				<th>总分</th>
			</tr>
		</thead>
		<tbody>
			{{range .GameEnded}}
			<tr>
				<td>{{.Id}}</td>
				<td>{{.Teama}}</td>
				<td>{{.Teamb}}</td>
				<td>{{.Timestarted}}</td>
				<td>{{.Concede}}</td>
				<td>{{.Oddsa}}</td>
				<td>{{.Oddsb}}</td>
				<td>{{.Scoresum}}</td>
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