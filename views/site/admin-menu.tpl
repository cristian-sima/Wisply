{{ define "admin-menu" }}
<header class="">
  <button type="button" class="text-center btn btn-default btn-sm visible-xs" id="close-sidebar-admin">
    <span class="glyphicon glyphicon-remove" aria-hidden="true"></span>
  </button>
</header>
<ul class="nav">
  <li>
    <a href="/admin">Dashboard</a>
  </li>
  <li>
    <a href="/admin/institutions" >Institutions </a>
  </li>
  <li>
    <a href="/admin/repositories/">Repositories</a>
  </li>
  <li>
    <a href="/admin/accounts" > Accounts </a>
  </li>
  <li>
    <a href="/admin/api" > API Settings</a>
  </li>
  <li>
    <a href="/admin/log" > Log </a>
  </li>
</ul>
{{ end }}
