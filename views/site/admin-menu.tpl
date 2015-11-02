{{ define "admin-menu" }}
<div class="admin-menu">
  <ul class="nav">
    <li>
      <a href="/admin">
        <span class="glyphicon glyphicon-th-large"></span>
        <span class="hidden-sm hidden-xs">Home</span>
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
      <a href="/admin/developers" >
        <span class="glyphicon glyphicon-save"></span>
        <span class="hidden-sm hidden-xs">
          Developers
        </span>
      </a>
    </li>
    <li>
      <a href="/admin/education" >
        <span class="glyphicon glyphicon-tasks"></span>
        <span class="hidden-sm hidden-xs">
          Education
        </span>
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
      <a href="/admin/log" >
        <span class="glyphicon glyphicon-list-alt"></span>
        <span class="hidden-sm hidden-xs">
          Logs
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
  </ul>
</div>
<div>
<style scoped>
.admin-menu {
  font-size: 15px;
}
</style>
</div>
{{ end }}
