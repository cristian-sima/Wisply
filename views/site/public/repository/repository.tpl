<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/institutions">Institutions</a></li>
            <li><a href="/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></li>
            <li class="active">{{ .repository.Name }}
              {{ if .currentAccount.IsAdministrator  }}
              <a href="/admin/repositories/repository/{{ .repository.ID }}"><span class="label label-default">Admin this</span></a>
              {{ end }}
            </li>
          </ul>
        </div>
        <div class="panel-body">
          <div class="row">
            <div class="col-lg-2 col-md-2 col-sm-2 text-center" >
              <span class="glyphicon glyphicon-equalizer institution-logo img-responsive "></span>
              <div class="text-left">
              </div>
            </div>
            <div class="col-lg-6 col-md-6 col-sm-6" >
              <div>
                <h1>{{ .repository.Name }}</h2>
                  <span class="text-muted">Repository</span>
                </div>
                <div>
                  {{ .repository.Description }}
                </div>
              </div>
              <div class="col-lg-4 col-md-4 col-sm-4" >
                  <div>Is part of <a href="/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></div>
                  <div>{{ .repository.Category }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <style>
    .institution-logo {
      font-size: 8em;
    }
    </style>
    <script src="/static/js/admin/institution/list.js"></script>
