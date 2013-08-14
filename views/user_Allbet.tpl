{{template "user" .}}
{{if .User}}
<div class="container">
	<div class="alert alert-info">
		<p> <strong>共有  {{.BetCount}}  个投注记录</strong>
		</p>
	</div>
	<div>
		<div class="header">
			<p> <strong>所有投注记录</strong>
			</p>
		</div>
		<table class="table table-bordered table-hover table-striped table-condensed ">
			<thead>
				<tr>
					<th>数额</th>
					<th>盈利</th>
					<th>投注详情</th>
					<th>时间</th>
				</tr>
			</thead>
			<tbody>
				{{range .UserAllBets}}
				<tr>
					<td>{{.Betamount}}</td>
					<td>{{.Profit}}</td>
					<td>
						<a href="http://blockchain.info/tx/{{.Txhash}}">{{.Txhash}}</a>
					</td>
					<td>{{.Bettime}}</td>
				</tr>
				{{end}}
			</tbody>
		</table>
	</div>
</div>
{{else}}
<div class="container">
	<div class="hero-unit"> <b>YOU ARE NOT ALLOWED HERE</b>
	</div>
</div>
{{end}}

{{template "footer" .}}