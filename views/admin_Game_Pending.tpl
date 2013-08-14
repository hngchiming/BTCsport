{{template "admin" .}}
{{if .IsAdmin}}
<div class="well span5 pull-right">
	<form class="form-horizontal" action="/admin/Game_Result">
		<div class="control-group">
			<label class="control-label" for="gid">赛事ID</label>
			<div class="controls">
				<input type="text" id="gid" placeholder="赛事ID" name="gid"></div>
		</div>
		<div class="control-group">
			<label class="control-label" for="result">结果</label>
			<div class="controls">
				<input type="text" id="result" placeholder="赛事结果(1:2)" name="result"></div>
		</div>
		<div class="control-group">
			<div class="controls">
				{{.xsrf}}
				<button type="submit" class="btn btn-primary">提交</button>
			</div>
		</div>
	</form>
</div>
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
			{{range .GamePending}}
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