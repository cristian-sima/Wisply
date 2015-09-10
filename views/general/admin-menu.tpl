{{ define "admin-menu" }}
  <div class="panel panel-default">
    <div class="panel-heading">Menu</div>
    <div class="panel-body">
      <ul class="nav nav-pills nav-stacked">
        <li ><a href="/admin">Home</a></li>
        <li class="dropdown">
          <a class="dropdown-toggle" data-toggle="dropdown" href="/admin/sources" aria-expanded="true">
            Source <span class="caret"></span>
          </a>
          <ul class="dropdown-menu">
            <li><a href="/admin/sources/add">Add</a></li>
            <li><a href="/admin/sources/">Manage</a></li>
          </ul>
        </li>
        <li >
          <a href="/admin/users" > Users </a>
        </li>
      </ul>
    </div>
</div>
{{ end }}
