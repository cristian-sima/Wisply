<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/institutions">Institutions</a></li>
            <li class="active">{{ .institution.Name }}
              {{ if .currentAccount.IsAdministrator  }}
              <a href="/admin/institutions/institution/{{ .institution.ID }}"><span class="label label-default">Admin this</span></a>
              {{ end }}
            </li>
          </ul>
        </div>
        <div class="panel-body">
          <div class="row">
            <div class="col-lg-2 col-md-2 col-sm-2 text-center" >
              <div class="institution-profile">
                <div class="insider">
                  {{ if eq .institution.LogoURL "" }}
                  <span class="glyphicon glyphicon-education institution-logo-default"></span>
                  {{ else }}
                  <img alt="{{ .institution.Name }}" src="{{ .institution.LogoURL }}" class="inlogo" />
                  {{ end }}
                </div>
              </div>
              <hr />
              <div class="text-left">
                <a href="{{ .institution.URL }}">Web page</a>
              </div>
            </div>
            <div class="col-lg-6 col-md-6 col-sm-6" >
              <div>
                <h1>{{ .institution.Name }}</h1>
                <span class="text-muted">Institution</span>
              </div>
              <div>
                {{ .institution.Description}} <a target="_blank" href="{{ .institution.WikiURL }}">Wikipedia</a>
              </div>
            </div>
            <div class="col-lg-4 col-md-4 col-sm-4" >
              <!-- Repositories -->
              <div class="list-group" id="repositories">
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
</div>
<div>
<style scoped>
.institution-logo {
  font-size: 13em;
}
</style>
</div>
<link href="/static/css/public/institution.css" type="text/css" rel="stylesheet" property='stylesheet' />
<script src="/static/js/admin/institution/list.js"></script>
