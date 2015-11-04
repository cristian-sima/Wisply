<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/institutions">Institutions</a></li>
      <li class="active">{{ .institution.Name }}</li>
    </ul></div>
    <div class="panel-body">
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
              <span class="text-muted">Address:</span> <a href="{{ .institution.URL }}">Web page</a>
            </div>
          </div>
          <div class="col-lg-6 col-md-6 col-sm-6" >
            <div>
              <h1>{{ .institution.Name }}</h1>
              <span class="text-muted">Institution</span>
            </div>
            <div>
              {{ .institution.Description}} <a href="/admin/institutions/{{ .institution.ID}}/modify">Modify</a>
            </div>
          </div>
          <div class="col-lg-4 col-md-4 col-sm-4" >
            <!-- Repositories -->
            <a href="/institutions/{{ .institution.ID}}">Public page</a>
            <br />
            <br />
            <a href="/admin/repositories/add/?institution={{ .institution.ID}}">
              <span class="glyphicon glyphicon-plus-sign text-success"> </span>
              Add repository
            </a>
            <br />
            <a href="/admin/institutions/{{ .institution.ID}}/program/add">
              <span class="glyphicon glyphicon-plus-sign text-success"> </span>
              Add program for this university
            </a>
            <br />
            <br />
            {{ if eq (not .repositories) true }}
            <div class="text-muted">
              :( it does not have repositories
            </div>
            {{ else }}
            <h4>Repositories ({{ .repositories | len }})</h4>
            <div class="list-group" id="repositories">
              {{range $index, $repository := .repositories}}
              <a href="/admin/repositories/{{ $repository.ID }}" class="list-group-item">
                <h4 class="list-group-item-heading"><span class="glyphicon glyphicon-equalizer"></span> {{ $repository.Name }}</h4>
                <p class="list-group-item-text">{{ $repository.Description }}</p>
              </a>
              {{ end }}
            </div>
            {{ end }}
          </div>
        </div>
      </div>
      <h4 id="programs">Programs of study ({{ .institutionPrograms | len }})</h4>
      <a href="/admin/institutions/{{ .institution.ID }}/program/add" class="btn btn-sm btn-success sm">
        <span class="glyphicon glyphicon-plus"></span> Add program
      </a>
      {{ if eq (not .institutionPrograms) true }}
      <div class="text-muted">
        :( there are no programs of study
      </div>
      {{ else }}
      <div class="table-responsive">
        <table class="table table-striped table-hover " id="programs-list">
          <thead>
            <tr>
              <th>Code</th>
              <th>Title</th>
              <th>Level</th>
              <th>Category</th>
              <th>Year</th>
              <th>Edit</th>
              <th>Delete</th>
            </tr>
          </thead>
          <tbody>
            {{ $institution := .institution }}
            {{range $index, $program := .institutionPrograms}}
            {{ $mainProgram := $program.GetProgram }}
            <tr>
              <td>{{ $program.GetCode }}</td>
              <td><a href="/admin/institutions/{{ $institution.ID }}/program/{{ $program.GetID }}">{{ $program.GetTitle }}</a></td>
              <td>{{ $program.GetLevel }}</td>
              <td><a href="/education/programs/{{ $mainProgram.GetID }}">{{ $mainProgram.GetName }}</a></td>
              <td>{{ $program.GetYear }}</td>
              <td><a href="/admin/institutions/{{ $institution.ID }}/program/{{ $program.GetID }}/modify">Edit</td>
              <td><a href="#" data-id="{{ $program.GetID }}" class="deleteProgramButton btn btn-danger btn-xs" >Delete</a></td>
            </tr>
            {{ end }}
          </tbody>
        </table>
      </div>
      {{ end }}
      <h4 id="modules">Modules ({{ .institutionModules | len }})</h4>
      {{ if eq (not .institutionModules) true }}
      <div class="text-muted">
        :( there are no programs modules
      </div>
      {{ else }}
      <a href="/admin/institutions/{{ .institution.ID }}/module/add" class="btn btn-sm btn-success sm">
        <span class="glyphicon glyphicon-plus"></span> Insert new module
      </a>
      <div class="table-responsive">
        <table class="table table-striped table-hover " id="modules-list">
          <thead>
            <tr>
              <th>Code</th>
              <th>Title</th>
              <th>Content</th>
              <th>Credits (CATS)</th>
            </tr>
          </thead>
          <tbody>
            {{ $institution   := .institution }}
            {{range $index, $module := .institutionModules}}
            <tr>
              <td>{{ $module.GetCode }}</td>
              <td>{{ $module.GetTitle }}</td>
              <td>{{ $module.GetContent }}</td>
              <td>{{ $module.GetCredits "CATS" }}</td>
              <td><a href="/admin/institutions/{{ $institution.ID }}/module/{{ $module.GetID }}/modify">Edit</td>
                <td><a href="#" data-id="{{ $module.GetID }}" class="deleteModuleButton btn btn-danger btn-xs" >Delete</a></td>
              </tr>
              {{ end }}
            </tbody>
          </table>
        </div>
        {{ end }}
      <hr />
      <div>
        <a class="btn btn-primary" href="/admin/institutions/{{ .institution.ID }}/advance-options">Advance options</a>
      </div>
    </div>
  </div>
  <div>
  <style scoped>
  .big-number {
    font-size: 30px;
  }
  </style>
</div>
<script src="/static/js/admin/institutions/institution/home.js"></script>
<script>

var institution = {
  id: {{ .institution.ID }},
};

$(document).ready(function(){
  var module = wisply.getModule("admin-institutions-institution-home"),
  manager = new module.Manager(institution);
  manager.init();
});
</script>
