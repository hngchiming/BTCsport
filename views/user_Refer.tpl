{{template "user" .}}
{{if .User}}
<div class="container">
	<div class="alert alert-block">
		<p> <strong>您推荐的用户总数: <i class="icon-user"></i>
				{{.Count}}</strong> 
		</p>
		<p> <strong>复制下面的地址，推荐给其他用户进行注册~</strong>
			<div class="alert alert-info">
				<a href="/login?username={{.Username}}">http://cryptoption.com/login?username={{.Username}}</a>
			</div>
		</p>
		<p>
			<strong>您本月投注量为： {{.betAmount}} BTC ---- 盈利为: {{.betProfit}} BTC</strong>
		</p>
		<p>
			<strong>若您本月投注量多于 10 BTC, 可以获得 {{.ReferFee}} BTC推广费，推广费详细规则请查看
				<a href="/about">关于</a></strong> 
		</p>
	</div>
	<table class="table table-bordered table-hover table-striped table-condensed ">
		<thead>
			<tr>
				<th>所推荐的用户ID</th>
				<th>本月总投注额</th>
				<th>本月盈利</th>
			</tr>
		</thead>
		<tbody>
			{{range .UserData}}
			<tr>
				<td>{{.Uid}}</td>
				<td>{{.Betamount}}</td>
				<td>{{.Profit}}</td>
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