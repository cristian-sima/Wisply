<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li class="active">Admin</li>
    </ul></div>
    <div class="panel-body">
      <div class="row">
        <div class="col-lg-4">
          <div class="list-group">
            <a class="list-group-item" href="admin/users">
              <span class="badge">{{ .numberOfUsers}}</span>
              Users
            </a>
            <a class="list-group-item" href="admin/sources">
              <span class="badge"> {{ .numberOfSources }}</span>
              Sources
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