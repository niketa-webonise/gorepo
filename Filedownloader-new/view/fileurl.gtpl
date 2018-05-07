<!DOCTYPE html>
<head>
	  <meta name="viewport" content="width=device-width,initial-scale=1.0"/>

  <link href="static/css/bootstrap.min.css" rel="stylesheet">

  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>

  <link href="static/css/file.css" type="text/css" rel="stylesheet"/>
</head>
<body>
    <div id="wrapper">

    <div class="form_div">
    
    <center>
    <p class="form_label">

    <h3>Type File URL To Download File Here !!</h3>

    </p>
<form action="/download" method="post"> 

 <div class="container">

	<input type="text" placeholder="enter file url" name="furl" required>
	<button type="submit" name="download">Download</button>

</div>
</form>
</center>
</body>


</html>