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
              <span class="glyphicon glyphicon-equalizer institution-logo-default "></span>
              <div class="text-left"></div>
              <hr />
              <div class="text-left">
                <span class="text-muted">Address:</span> <a target="_blank" href="{{ .repository.PublicURL }}">Web page</a>
              </div>
            </div>
            <div class="col-lg-6 col-md-6 col-sm-6" >
              <div>
                <h1>{{ .repository.Name }}</h1>
                <span class="text-muted">Repository</span>
              </div>
              <div>
                {{ .repository.Description }}
              </div>
            </div>
            <div class="col-lg-4 col-md-4 col-sm-4" >
              <div>Part of <a href="/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></div>
              <div><i>{{ .repository.Category }}</i> repository</div>
            </div>
          </div>
          <div>
            <br />
            <br />
          </div>
          <!-- Things -->
          {{ if eq .repository.LastProcess 0 }}
          <div class="text-center text-muted">
            Wisply did not process this repository yet... :(
          </div>
          {{ else }}
          <div class="row">
            <!-- Statistics -->
            <div class="col-lg-2 col-md-2 col-sm-2 text-left" >
              <table class="table">
                <tbody>
                  <tr>
                    <td>
                      <span class="badge badge-info">{{ .process.Formats }}</span>
                    </td>
                    <td>
                      formats
                    </td>
                  </tr>
                  <tr>
                    <td>
                      <span class="badge badge-info">{{ .process.Collections }}</span>
                    </td>
                    <td>
                      collections
                    </td>
                  </tr>
                  <tr>
                    <td>
                      <span class="badge badge-info">{{ .process.Records }}</span>
                    </td>
                    <td>
                      records
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
            <!-- Resources -->
            <div class="col-lg-6 col-md-6 col-sm-6" >
              <div id="repository-before-resources"></div>
              <div id="repository-resources"></div>
            </div>
            <div class="col-lg-4 col-md-4 col-sm-4" >
              <!-- Collections -->
              <div class="list-group" id="repositories">
                {{range $index, $collection := .collections}}
                <a href="#" class="list-group-item">
                  <h5 class="list-group-item-heading"><span class="glyphicon glyphicon-equalizer"></span> {{ $collection.Name }}</h5>
                  <p class="list-group-item-text">{{ $collection.Description }}</p>
                </a>
                {{ end }}
              </div>
            </div>
          </div>
          {{ end }}
        </div>
      </div>
    </div>
  </div>
</div>
<script src="/static/js/admin/repository/list.js"></script>
<script>
var server = {}
server.repository = {
  id : {{ .repository.ID }},
  totalRecords: {{ .process.Records }},
}
</script>
<link href="/static/css/public/institution.css" type="text/css" rel="stylesheet" property='stylesheet' />
<script src="/static/js/public/repository.js"></script>
