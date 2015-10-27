{{ define "menu" }}
<div class="navbar navbar-default navbar-fixed-top">
  <div class="container">
    <div class="navbar-header">
      <a href="/" class="navbar-brand" id="full-logo">
        <img style="display:inline" src="/static/img/wisply/logo/jpg.jpg" alt="Wisply Top Logo"/> Wisply
      </a>
      <button class="navbar-toggle" type="button" title="Show Menu" label="Show Menu" data-toggle="collapse" data-target="#navbar-main">
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <button class="navbar-toggle" type="button" id="show-small-search-button" title="Show small search button">
        <span class="glyphicon glyphicon-search"></span>
      </button>
      <form class="navbar-brand navbar-right hideMe " id="search-small" style="width:50%;"  >
        <div class="form-group visible-xs" style="width:100%;" >
          <div style="position:relative">
            <input type="text" style="width:276px;"  id="search-small-input" class="wisply-search-field form-control" placeholder="Search">
            <img alt="Loading Spinner Mobile" style="display:none" class="wisply-search-field-spinner search-spinner" src='/static/img/wisply/load.gif' />
          </div>
        </div>
      </form>
    </div>
    <nav class="navbar-collapse collapse" id="navbar-main">
      <ul class="nav navbar-nav hidden-sm">
        <li class="">
          <a href="/about">About</a>
        </li>
        <li class="dropdown">
          <a href="" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
            Curricula <span class="caret"></span>
          </a>
          <ul class="dropdown-menu" role="menu">
            {{range $index, $program := .programs}}
            <li>
              <a href="/curriculum/{{$program.GetID}}">
                {{ $program.GetName }}
              </a>
            </li>
            {{ end }}
          </ul>
        </li>
        <li class="hidden-sm">
          <a href="/institutions">
            Institutions
          </a>
        </li>
      </ul>
      <ul class="nav navbar-nav visible-sm">
        <li class="dropdown" >
          <a href="#" class="dropdown-toggle" data-toggle="dropdown">
            More
            <b class="caret"></b>
          </a>
          <ul class="dropdown-menu multi-level">
            <li>
              <a href="/curricula" >
                Curricula
              </a>
            </li>
            <li class="divider"></li>
            <li><a href="/about">About</a></li>
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
          <a href="/account" data-toggle='tooltip' data-placement='bottom' title='' data-original-title='Dashboard'>
            <b>{{ .currentAccount.Name }}</b>
            <span class="text-success"><span class="glyphicon glyphicon-user"></span></span>
          </a>
        </li>
        {{ if .currentAccount.IsAdministrator }}
        <li>
          <a href="/admin" data-toggle='tooltip' data-placement='bottom' title='' data-original-title='Admin'>
            <span class="glyphicon glyphicon-cog"></span>
          </a>
        </li>
        {{ end }}
        <li>
          <a id="menu-logout-button" href="#" data-toggle='tooltip' data-placement='bottom' data-original-title='Logout' title="Logout">
            <span class="glyphicon glyphicon-log-in"></span>
          </a>
        </li>
        {{ end }}
      </ul>
      <form class="navbar-form navbar-right hidden-xs">
        <div class="form-group">
          <div style="position:relative">
            <input aria-label="Search" id="search-2" type="text" style="width: 278px;" class="wisply-search-field form-control" placeholder="Search">
            <img alt="Loading Spinner Desktop" style="display:none" class="wisply-search-field-spinner search-spinner" src='/static/img/wisply/load.gif' />
          </div>
        </div>
      </form>
    </nav>
  </div>
</div>
{{ end }}
