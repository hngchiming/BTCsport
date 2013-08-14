{{template "admin" .}}
{{if .IsAdmin}}
<div class="container well">
	<form class="form-horizontal" action="/admin/Game_New" method="POST">
		<fieldset>
			<label class="control-label" for="TeamA">主队</label>
			<div class="controls">
				<input type="text" id="TeamA" name="TeamA"></div>
			<br>

			<label class="control-label" for="TeamB">客队</label>
			<div class="controls">
				<input type="text" id="TeamB" name="TeamB"></div>
			<br>

			<label class="control-label" for="Oddsa">主队赔率</label>
			<div class="controls">
				<input type="text" id="Oddsa" name="Oddsa"></div>
			<br>

			<label class="control-label" for="Oddsb">客队赔率</label>
			<div class="controls">
				<input type="text" id="Oddsb" name="Oddsb"></div>
			<br>

			<label class="control-label" for="Concede">让分</label>
			<div class="controls">
				<input type="text" id="Concede" name="Concede"></div>
			<br>

			<label class="control-label" for="ScoreSum">总分</label>
			<div class="controls">
				<input type="text" id="ScoreSum" name="ScoreSum"></div>
			<br>

			<label class="control-label" for="TimeStart">开始时间</label>
			<div class="controls">
				<input type="text" id="TimeStart" name="TimeStart"></div>
			<br>

			<label class="control-label">体育类别</label>
			<div class="controls">
				<select name="Type">
					<option value="Football">足球</option>
					<option value="Basketball">篮球</option>
				</select>
			</div>
			<br>
			{{ .xsrf }}
			<div class="control-group">
				<div class="controls">
					<button type="submit" class="btn btn-primary">发布</button>
				</div>
			</div>
		</fieldset>
	</form>
</div>
{{else}}
<div class="container">
	<div class="hero-unit"> <b>YOU ARE NOT ALLOWED HERE</b>
	</div>
</div>
{{end}}
{{template "footer" .}}