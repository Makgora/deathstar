<!DOCTYPE html>

<html>
<head>
  <title>Parked</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="shortcut icon" href="../static/img/logo_parked.png" type="image/x-icon" />

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-image: url('../static/img/background_flou.png');
    }

    header,
    footer {
      width: 960px;
      margin-left: auto;
      margin-right: auto;
    }

    .logo {
      background-image: url('../static/img/logo_parked.png');
      background-repeat: no-repeat;
      -webkit-background-size: 100px 100px;
      background-size: 100px 100px;
      background-position: center center;
      text-align: center;
      font-size: 42px;
      padding: 250px 0 70px;
      font-weight: normal;
      text-shadow: 0px 1px 2px #ddd;
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 150px 0;
      color: #999;
    }

    .description {
      text-align: center;
      font-size: 16px;
    }

    a {
      color: #444;
      text-decoration: none;
    }

  </style>
</head>

<body>
  <header>
    <h1 class="logo">Welcome to Parked's optimization tool</h1>
    <div class="description">
      Building...
    </div>
  </header>
  <footer>
    <div class="author">
      Official website:
      <a href="http://{{.Website}}">{{.Website}}</a>
      <br>
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
  </footer>

  <script src="/static/js/reload.min.js"></script>
</body>
</html>
