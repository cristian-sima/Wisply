{{ define "menu" }}
<div class="navbar navbar-default navbar-fixed-top">
  <div class="container">
    <div class="navbar-header">
      <a href="/" class="navbar-brand" id="full-logo">
        <img id="logo" src="/static/img/wisply/logo/jpg.jpg" alt="Logo"/> Wisply
      </a>
      {{ if .isAdminPage }}
      <button type="button" class="navbar-toggle btn-lg" data-toggle="offcanvas" data-target=".sidebar-nav">
        <span class="glyphicon glyphicon-cog"></span>
      </button>
      {{ end }}
      <button class="navbar-toggle" type="button" data-toggle="collapse" data-target="#navbar-main">
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <button class="navbar-toggle" type="button" id="show-small-search-button">
        <span class="glyphicon glyphicon-search"></span>
      </button>
      <form class="navbar-brand navbar-right hideMe " id="search-small" role="search" style="width:50%;"  >
        <div class="form-group visible-xs" style="width:100%;" >
          <input type="text" style="width:276px;"  id="search-small-input" class="wisply-search-field form-control" placeholder="Search">
        </div>
      </form>
    </div>
    <nav class="navbar-collapse collapse" id="navbar-main">
      <ul class="nav navbar-nav hidden-sm">
        <li class="">
          <a href="/about" id="themes">About</a>
        </li>
        <li class="hidden-sm">
          <a href="/webscience">Web science</a>
        </li>
        <li class="hidden-sm">
          <a href="/institutions">Institutions</a>
        </li>
      </ul>
      <ul class="nav navbar-nav visible-sm">
        <li class="dropdown" >
          <a href="#" class="dropdown-toggle" data-toggle="dropdown">More <b class="caret"></b></a>
          <ul class="dropdown-menu">
            <li class="divider"></li>
            <li><a href="/webscience">Web science</a></li>
            <li><a href="/about" id="themes">About</a></li>
            <li><a href="/institutions">Institutions</a></li>
          </ul>
        </li>
      </ul>
      <ul class="nav navbar-nav navbar-right" id="menu-top-left">
        {{ if .accountDisconnected }}
        <li>
          <a href="/auth/login">Login</a>
        </li>
        <li>
          <a href="/auth/register">Register</a>
        </li>
        {{ end }}
        {{ if .accountConnected }}
        <li  class="text-muted">
          <a><b>{{ .currentAccount.Name }}</b> <span class="text-success"><span class="glyphicon glyphicon-user"></span></span></a>
        </li>
        {{ if .currentAccount.IsAdministrator }}
        <li>
          <a href="/admin"><span class="glyphicon glyphicon-cog"></span></a>
        </li>
        {{ end }}
        <li>
          <a id="menu-logout-button" href="#" title="Logout"><span class="glyphicon glyphicon-log-in"></span></a>
        </li>
        {{ end }}
      </ul>
      <form class="navbar-form navbar-right hidden-xs" role="search">
        <div class="form-group">
          <input type="text" style="width: 278px;" class="wisply-search-field form-control" placeholder="Search">
        </div>
      </form>
    </nav>
  </div>
</div>
{{ end }}
