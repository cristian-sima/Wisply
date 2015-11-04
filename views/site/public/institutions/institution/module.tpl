<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/institutions">Institutions</a></li>
            <li><a href="/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></li>
            <li class="active">{{ .module.GetTitle }}</li>
          </ul>
        </div>
        <div class="panel-body">
          <div style="margin:0px">
            <div class="row">
              <div class="col-md-9">
                <h1>{{ .module.GetTitle }}</h1>
                <span class="text-muted">Module</span> &bull; <a href="/institutions/{{ .institution.ID }}">{{ .institution.Name}}</a>
                <div class="well">
                {{ .module.GetContent }}
              </div>
              </div>
              <div class="col-md-3">
                <h2>Information</h2>
                <table class="table">
                  <tbody>
                    <tr>
                      <td>Code</td>
                      <td>{{ .module.GetCode }}</td>
                    </tr>
                    <tr>
                      <td>Year</td>
                      <td>{{ .module.GetYear }}</td>
                    </tr>
                    <tr>
                      <td>CATS</td>
                      <td>{{ .module.GetCredits "CATS" }}</td>
                    </tr>
                    <tr>
                      <td>ECTS</td>
                      <td>{{ .module.GetCredits "ECTS" }}</td>
                    </tr>
                    <tr>
                      <td>US credits</td>
                      <td>{{ .module.GetCredits "US" }}</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
          <br />
          <br />
          <h2>Is part of</h2>
          {{ $programs := .module.GetPrograms }}
          {{ if eq ($programs | len) 0 }}
          <div class="text-muted">There are no programs of study which include this module.</div>
          {{ else }}
          <div class="list-group">
            {{ $institution := .institution }}
            {{ range $index, $program := $programs }}
            <a href="/institutions/{{ $institution.ID }}/program/{{ $program.GetID }}" class="list-group-item active">
              {{ $program.GetCode }} - {{ $program.GetTitle }}
            </a>
            {{ end }}
          </div>
          {{ end }}
        </div>
      </div>
    </div>
  </div>
</div>
