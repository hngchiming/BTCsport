{{template "user" .}}
{{if .User}}
<div class="container">
	<div class="alert alert-success">
		<div>您的比特币地址是： {{.Addr}} 直接往该账户充值即可，2个确认后自动到账</div>
		<div>
			当前账户余额: {{.Balance}} <i class="icon-bold"></i>
			TC
		</div>
	</div>
	<div>
		<div class="header">
			<p> <strong>所有充值记录</strong>
			</p>
		</div>
		<table class="table table-bordered table-hover table-striped table-condensed ">
			<thead>
				<tr>
					<th>用户ID</th>
					<th>数额</th>
					<th>时间</th>
				</tr>
			</thead>
			<tbody>
				{{range .UserDeposit}}
				<tr>
					<td>{{.Uid}}</td>
					<td>{{.Amount}}</td>
					<td>{{.Time}}</td>
				</tr>
				{{end}}
				<tr>
					<td>总充值量</td>
					<td>{{.AllTimeDeposit}}</td>
					<td></td>
				</tr>
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