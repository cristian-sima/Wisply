{{ define "account-menu" }}
<div class="navbar">
  <div class="container-fluid">
    <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-2">
      <ul class="nav navbar-nav">
        <li class="active"><a href="/account">Dashboard<span class="sr-only">(current)</span></a></li>
        <li class="dropdown">
          <a href="http://bootswatch.com/paper/#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">Activity <span class="caret"></span></a>
          <ul class="dropdown-menu" role="menu">
            <li><a href="/account/search">Search</a></li>
          </ul>
        </li>
        <li><a href="/account/settings">Settings</a></li>
      </ul>
    </div>
  </div>
</div>

{{ end }}
