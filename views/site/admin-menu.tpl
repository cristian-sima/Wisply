{{ define "admin-menu" }}
  <header class="">
    <button type="button" class="text-center btn btn-default btn-sm visible-xs" id="close-sidebar-admin">
      <span class="glyphicon glyphicon-remove" aria-hidden="true"></span>
    </button>
    </header>
      <ul class="nav">
        <li ><a href="/admin">Dashboard</a></li>
        <li >
          <a href="/admin/institutions" >Institutions </a>
        </li>
        <li><a href="/admin/repositories/">Repositories</a></li>
        <!-- <li class="dropdown">
          <a class="dropdown-toggle" data-toggle="dropdown" href="/admin/repositories" aria-expanded="true">
            Repository <span class="caret"></span>
          </a>
          <ul class="dropdown-menu">
            <li><a href="/admin/repositories/add">Add</a></li>
            <li><a href="/admin/repositories/">Manage</a></li>
          </ul>
        </li> -->
        <li >
          <a href="/admin/accounts" > Accounts </a>
        </li>
        <li >
          <a href="/admin/log" > Log </a>
        </li>
      </ul>
{{ end }}
