{{define "admin"}}
{{template "header"}}
<title>管理中心</title>
</head>
{{template "nav-bar" .}}
			    <div class="pull-left span2 well">
			  			<ul class="nav nav-list">
							<li class="nav-header">会员管理</li>
			  				<li {{if .User_All}}class="active"{{end}}><a href="/admin/User_All">所有用户</a></li>
				    		<li {{if .User_Refer}}class="active"{{end}}><a href="/admin/User_Refer">推广系统</a></li>

				    		<li class="nav-header">游戏管理</li>
				    		<li {{if .Game_New}}class="active"{{end}}><a href="/admin/Game_New">发布新体育竞猜</a></li>
				    		<li {{if .Game_Pending}}class="active"{{end}}><a href="/admin/Game_Pending">正在投注中</a></li>
				    		<li {{if .Game_End}}class="active"{{end}}><a href="/admin/Game_End">已过期赛事</a></li>
			    			
			    			<li class="nav-header">钱包管理</li>
			    			<li {{if .Wallet_Deposit}}class="active"{{end}}><a href="/admin/Wallet_Deposit">用户充值记录</a></li>
			    			<li {{if .Wallet_Withdraw}}class="active"{{end}}><a href="/admin/Wallet_Withdraw">用户提现记录</a></li>
			    			<li {{if .Wallet_Pay}}class="active"{{end}}><a href="/admin/Wallet_Pay">自定义转账</a></li>

			    		</ul>
			    </div>
{{end}}