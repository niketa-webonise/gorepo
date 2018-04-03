<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width,initial-scale=1.0"/>

<link href="static/css/bootstrap.min.css" rel="stylesheet">

<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>

<link href="static/css/login.css" type="text/css" rel="stylesheet"/>
</head>
<body>
<div id="wrapper">

<div class="form_div">
<p class="form_label">LOGIN FORM</p>
<form method="post" action="/login">



  <div class="container">
    <label for="uname"><b>Email</b></label>
    <input type="text" placeholder="Enter Email" name="email" required>

    <label for="psw"><b>Password</b></label>
    <input type="password" placeholder="Enter Password" name="pwd" required>

    <button type="submit">Login</button>
    <label>
      <input type="checkbox" checked="checked" name="remember"> Remember me
    </label>
  </div>

  </form>
  </div>
  </div>

  </body>
  <center><p>{{ .Msg }}</p></center>
  </html>