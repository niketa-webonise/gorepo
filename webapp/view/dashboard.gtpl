<html>
	<head>
		<title></title>
	</head>

	<h1>
		wow! Login successful...
	</h1>
	<body>
	<form action="/logout" method="post">
		<p>Your EmailID</p>
		<p>{{ .Email }}</p>

		<p>Your firstname</p>
		<p>{{ .FirstName }}</p>

		<p>Your lastname</p>
		<p>{{ .LastName }}</p>
		
		<center>
			<button type="submit" name="logout">LOGOUT</button>
		</center>
</form>
</body>
</html>