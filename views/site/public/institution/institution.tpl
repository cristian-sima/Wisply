<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/institutions">Institutions</a></li>
            <li class="active">{{ .institution.Name }}</li>
          </ul>
        </div>
        <div class="panel-body">
          <div class="row">
            <div class="col-lg-3 col-md-3 col-sm-3 text-center" >
              <span class="glyphicon glyphicon-education institution-logo img-responsive "></span>
              <div class="text-left">
                <span class="text-muted">Address:</span> <a href="{{ .institution.URL }}">Web page</a>
              </div>
            </div>
            <div class="col-lg-5 col-md-5 col-sm-5" >
              <div>
                <h1>{{ .institution.Name }}</h2>
                  <span class="text-muted">Institution</span>
                </div>
                <div>
                  {{ .institution.Description}}
                </div>
              </div>
              <div class="col-lg-4 col-md-4 col-sm-4" >
                    {{ if .currentAccount.IsAdministrator  }}
                    <div><a href="/admin/institutions/institution/{{ .institution.ID }}"><span class="label label-default">Admin this</span></a></div>
                    {{ end }}
                  <!-- Repositories -->
                  <div class="list-group">
                  {{range $index, $repository := .repositories}}
                    <a href="/repository/{{ $repository.ID }}" class="list-group-item">
                      <h4 class="list-group-item-heading"><span class="glyphicon glyphicon-equalizer"></span> {{ $repository.Name }}</h4>
                      <p class="list-group-item-text">{{ $repository.Description }}</p>
                    </a>
                  {{ end }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <style>
    .institution-logo {
      font-size: 13em;
    }
    </style>
    <script src="/static/js/admin/institution/list.js"></script>
