<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/institutions">Institutions</a></li>
            <li><a href="/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></li>
            <li class="active">{{ .program.GetTitle }}</li>
          </ul>
        </div>
        <div class="panel-body">
          <h1>{{ .program.GetTitle }}</h1>
          <div class="well">{{ .program.GetContent }}</div>
          <br />
          <br />
          {{ if eq (not .modules) true }}
          <div class="text-muted">
            :( there are no moduels for this program of study
          </div>
          {{ else }}
          <h4>Modules:</h4>
          <div class="table-responsive">
            <table class="table table-striped table-hover " id="modules-list">
              <thead>
                <tr>
                  <th>Title</th>
                  <th>Code</th>
                  <th>Year</th>
                </tr>
              </thead>
              <tbody>
                {{ $institution   := .institution }}
                {{ $program       := .program }}
                {{range $index, $module := .modules}}
                <tr>
                  <td><a href="{{ $program.GetID }}/module/{{ $module.GetID }}">{{ $module.GetTitle }}</a></td>
                  <td>{{ $module.GetCode }}</td>
                  <td>{{ $module.GetYear }}</td>
                  {{ end }}
                </tbody>
              </table>
            </div>
            {{ end }}
        </div>
      </div>
    </div>
  </div>
</div>
