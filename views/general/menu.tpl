{{ define "menu" }}
<div class="navbar navbar-default navbar-fixed-top">
  <div class="container">
    <div class="navbar-header">
      <a href="/" class="navbar-brand"> <img id="logo" src="/static/img/logo/logo.jpg" alt="Logo"/> Wisply</a>
      <button class="navbar-toggle" type="button" data-toggle="collapse" data-target="#navbar-main">
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
    </div>
    <div class="navbar-collapse collapse" id="navbar-main">
      <ul class="nav navbar-nav">
        <li class="dropdown">
          <a href="/about" id="themes">About</a>
        </li>
        <li>
          <a href="/webscience">Web science</a>
        </li>
        <li>
          <a href="/contact">Contact</a>
        </li>
      </ul>
      <ul class="nav navbar-nav navbar-right" id="menu-top-left">
        {{ if .userDisconnected }}
        <li><a href="/auth/login">Login</a></li>
        <li><a href="/auth/register">Register</a></li>
        {{ end }}
        {{ if .userConnected }}
        <li  class="text-muted"><a>Hi, <b>{{ .currentUser.Username }}</b></a></li>
        {{ if .currentUser.Administrator }}
        <li><a href="/admin">Admin</a></li>
        {{ end }}
        <li><a id="menu-logout-button" href="#">Logout</a></li>
        {{ end }}
      </ul>
    </div>
  </div>
</div>
{{ end }}
