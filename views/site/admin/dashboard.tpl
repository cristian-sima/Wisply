<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li class="active">Admin</li>
    </ul></div>
    <div class="panel-body">
      <div class="row">
        <div class="col-lg-4">
          <div class="list-group">
            <a class="list-group-item" href="/admin/accounts">
              <span class="badge">{{ .numberOfAccounts}}</span>
              Accounts
            </a>
            <a class="list-group-item" href="/admin/repositories">
              <span class="badge"> {{ .numberOfRepositories }}</span>
              Repositories
            </a>
          </div>
        </div>
        <div class="col-lg-4">
        </div>
        <div class="col-lg-4">
        </div>
      </div>
    </div>
  </div>
