<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li class="active">Admin</li>
    </ul></div>
    <br />
    <br />
    <div class="panel-body">
      <div class="row">
        <div class="cube text-center col-md-3 col-sm-6 col-xs-12">
          <a href="/admin/accounts">
            <span class="glyphicon glyphicon-user"></span> Accounts <span class="badge">{{ .numberOfAccounts}}</span>
          </a>
        </div>
        <div class="cube text-center col-md-3 col-sm-6  col-xs-12">
          <a href="/admin/repositories">
            <span class="glyphicon glyphicon-equalizer"></span> Repositories <span class="badge">{{ .numberOfRepositories}}</span>
          </a>
        </div>
        <div class="cube text-center col-md-3 col-sm-6 col-xs-12">
          <a href="/admin/institutions">
            <span class="glyphicon glyphicon-user"></span> Institutions
          </a>
        </div>
        <div class="cube text-center col-md-3 col-sm-6  col-xs-12">
          <a href="/admin/developers">
            <span class="glyphicon glyphicon-dashboard"></span> Developers
          </a>
        </div>
        <div class="cube text-center col-md-3 col-sm-6 col-xs-12">
          <a href="/admin/log">
            <span class="glyphicon glyphicon-list-alt"></span> Log
          </a>
        </div>
        <div class="cube text-center col-md-3 col-sm-6 col-xs-12">
          <a href="/admin/education">
            <span class="glyphicon glyphicon-list-alt"></span> Education
          </a>
        </div>
      </div>
  </div>
</div>
<div>
  <style scoped>
  .cube {
    height: 100px;
  }
  </style>
</div>
