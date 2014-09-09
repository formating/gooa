<html><!doctype html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no" />
	<!-- Apple devices fullscreen -->
	<meta name="apple-mobile-web-app-capable" content="yes" />
	<!-- Apple devices fullscreen -->
	<meta names="apple-mobile-web-app-status-bar-style" content="black-translucent" />
	
	<title>{{.context.Sitename}} - 注册</title>

	<!-- Bootstrap -->
	<link rel="stylesheet" href="{{"static/css/bootstrap.min.css"|u}}">
	<!-- Bootstrap responsive -->
	<link rel="stylesheet" href="{{"static/css/bootstrap-responsive.min.css"|u}}">
	<!-- icheck -->
	<link rel="stylesheet" href="{{"static/css/plugins/icheck/all.css"|u}}">
	<!-- Theme CSS -->
	<link rel="stylesheet" href="{{"static/css/style.css"|u}}">
	<!-- Color CSS -->
	<link rel="stylesheet" href="{{"static/css/themes.css"|u}}">


	<!-- jQuery -->
	<script src="{{"static/js/jquery.min.js"|u}}"></script>
	
	<!-- Nice Scroll -->
	<script src="{{"static/js/plugins/nicescroll/jquery.nicescroll.min.js"|u}}"></script>
	<!-- Validation -->
	<script src="{{"static/js/plugins/validation/jquery.validate.min.js"|u}}"></script>
	<script src="{{"static/js/plugins/validation/additional-methods.min.js"|u}}"></script>
	<!-- icheck -->
	<script src="{{"static/js/plugins/icheck/jquery.icheck.min.js"|u}}"></script>
	<!-- Bootstrap -->
	<script src="{{"static/js/bootstrap.min.js"|u}}"></script>
	<script src="{{"static/js/eakroko.js"|u}}"></script>

	<!--[if lte IE 9]>
		<script src="{{"static/js/plugins/placeholder/jquery.placeholder.min.js"|u}}"></script>
		<script>
			$(document).ready(function() {
				$('input, textarea').placeholder();
			});
		</script>
	<![endif]-->
	

	<!-- Favicon -->
	<link rel="shortcut icon" href="{{"static/img/favicon.ico"|u}}" />
	<!-- Apple devices Homescreen icon -->
	<link rel="apple-touch-icon-precomposed" href="{{"static/img/apple-touch-icon-precomposed.png"|u}}" />

</head>

<body class='login'>
	<div class="wrapper wrapperreg">
		<h1><a href="{{""|u}}"><img src="{{"static/img/logo.png"|u}}" alt="" class='retina-ready' width="59" height="49">FLAT</a></h1>
		{{if .flash.error}}
			{{str2html .flash.error}}
		{{end}}
		<div class="login-body">
			<h2>注册</h2>
			<form action="{{"reg"|u}}" method='post' class='form-validate' id="test">
				<div class="control-group">
					<div class="email controls">
						<input type="text" name='name' placeholder="企业名称" value="{{.flash.name}}" class='input-block-level' data-rule-required="false">
					</div>
				</div>
				<div class="control-group">
					<div class="email controls">
						<input type="text" name="phone" placeholder="电话" value="{{.flash.phone}}" class='input-block-level' data-rule-required="false">
					</div>
				</div>
				<div class="control-group">
					<div class="email controls">
						<input type="text" name="address" placeholder="地址" value="{{.flash.address}}" class='input-block-level' data-rule-required="false">
					</div>
				</div>
				<div class="control-group">
					<div class="email controls">
						<input type="text" name="email" placeholder="邮箱" value="{{.flash.email}}" class='input-block-level' data-rule-required="false" data-rule-email="true">
					</div>
				</div>
				<div class="control-group">
					<div class="email controls">
						<input type="password" name="pw" placeholder="密码" class='input-block-level' data-rule-required="false">
					</div>
				</div>
				<div class="control-group">
					<div class="email controls">
						<input type="password" name="pw_confirm" placeholder="密码确认" class='input-block-level' data-rule-required="false">
					</div>
				</div>
				<div class="submit">
					<input type="submit" value="注册" class='btn btn-primary'>
				</div>
			</form>
			<div class="forget">
				<a href="{{""|u}}"><span>已有账号，登录?</span></a>
			</div>
		</div>
	</div>
</body>

</html>