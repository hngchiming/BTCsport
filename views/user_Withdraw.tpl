{{template "user" .}}
{{if .User}}
<div class="container">
	{{if .Error}}
	<div class="alert alert-error">
		<button type="button" class="close" data-dismiss="alert">&times;</button>
		<p> <strong>{{.Error}}</strong>
		</p>
	</div>
	{{end}}
	<div class="well">
		<form class="inline" action="/user" method="POST">
			<fieldset>
				<div>
					<input type="text" class="input-xlarge" placeholder="比特币提现地址" name="address">
					<input type="text" class="input-small" placeholder="提现数额" name="amount"></div>
				<div>
					<input type="password" class="input-medium" placeholder="资金密码" name="fundpassword">
					<input type="text" class="input-xlarge" placeholder="Email" name="email"></div>
				{{.xsrf}}
				<button type="submit" class="btn">提现</button>
			</fieldset>
		</form>
	</div>

	<div class="container">
		<div class="header">
			<p> <strong>所有提现记录</strong>
			</p>
		</div>
		<table class="table table-bordered table-hover table-striped table-condensed ">
			<thead>
				<tr>
					<th>用户ID</th>
					<th>数额</th>
					<th>地址</th>
					<th>时间</th>
				</tr>
			</thead>
			<tbody>
				{{range .UserWithdraw}}
				<tr>
					<td>{{.Uid}}</td>
					<td>{{.Amount}}</td>
					<td>
						<a href="http://blockchain.info/address/{{.Address}}"></a>
						{{.Address}}
					</td>
					<td>{{.Time}}</td>
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