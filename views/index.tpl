{{template "header"}}
<title>BTC-Sport</title>
</head>
{{template "nav-bar" .}}
<div class="container">
	<div class="well">

		<!-- <div class="">
		-->
		<div id="top" class="pull-left">
			<div class="text-center">
				<h3> <i class="icon-bold"></i>
					TC 体育竞猜
				</h3>
				<div class="copy-headline">
					<br>
					<br>
					<img id="main-image" src="/static/img/logo.png" alt="BTCSport Logo"></div>
				<br>
				<br>
				<hr></div>
		</div>
		<!-- </div>
		-->
		<br>
		<div class="row">
			<div class="span4">
				<h4> <i class="icon icon-user"></i>
					用户注册
				</h4>
				<p>
					用户需要先
					<a href="/login">注册</a>
					一个账户，获得专用的比特币地址，登录然后充值比特币，方能参与竞猜
				</p>
			</div>

			<div class="span4">
				<h4>
					<i class="icon-star-empty"></i>
					推广系统
				</h4>
				<p>
					推荐更多用户参与竞猜，推荐者将获取更多回报，回报详情请查看
					<a href="/about">关于</a>
				</p>
			</div>
			<div class="span4">
				<h4>
					<i class="icon-bold"></i>
					Blockchain在线钱包
				</h4>
				<p>竞猜平台使用Blockchain在线钱包服务，所有充值、提现、竞猜都是自动转账的，安全、高效，所有交易都受到监控</p>
			</div>
			<div class="span4">
				<h4>
					<i class="icon-ok"></i>
					零手续费
				</h4>
				<p>竞猜平台不收取额外手续费，所扣除的0.0005BTC手续费为转账时Blockchain在线钱包收取的矿工费用</p>
			</div>
			<div class="span4">
				<p></p>
			</div>
			<div class="span4">
				<h4>
					<i class="icon-book"></i>
					钱包地址
				</h4>
				<p>
					<ul>
						<li>
							投注池：
							<a href="http://blockchain.info/address/{{.BetAddr}}">点击这里查看</a>
						</li>
						<li>
							平台收益：
							<a href="http://blockchain.info/address/{{.ProfitAddr}}">点击这里查看</a>
						</li>
						<li>每个用户都有专用充值地址</li>
					</ul>
				</p>
			</div>
		</div>
	</div>
</div>
{{template "footer" .}}