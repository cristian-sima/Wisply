{{ define "admin-menu" }}
<ul class="nav">
  <li>
    <a href="/admin">
      <span class="glyphicon glyphicon-th-large"></span>
      <span class="hidden-sm hidden-xs">Dashboard</span>
    </a>
  </li>
  <li>
    <a href="/admin/institutions" >
      <span class="glyphicon glyphicon-education"></span>
      <span class="hidden-sm hidden-xs">
        Institutions
      </span>
    </a>
  </li>
  <li>
    <a href="/admin/repositories/">
      <span class="glyphicon glyphicon-equalizer"></span>
      <span class="hidden-sm hidden-xs">
        Repositories
      </span>
    </a>
  </li>
  <li>
    <a href="/admin/accounts" >
      <span class="glyphicon glyphicon-user"></span>
      <span class="hidden-sm hidden-xs">
        Accounts
      </span>
    </a>
  </li>
  <li>
    <a href="/admin/api" >
      <span class="glyphicon glyphicon-save"></span>
      <span class="hidden-sm hidden-xs">
        API Settings
      </span>
    </a>
  </li>
  <li>
    <a href="/admin/log" >
      <span class="glyphicon glyphicon-list-alt"></span>
      <span class="hidden-sm hidden-xs">
        Logs
      </span>
    </a>
  </li>
  <li>
    <a href="/admin/curriculum" >
      <span class="glyphicon glyphicon-tasks"></span>
      <span class="hidden-sm hidden-xs">
        Curricula
      </span>
    </a>
  </li>
</ul>
{{ end }}
