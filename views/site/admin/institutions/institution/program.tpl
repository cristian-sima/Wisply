<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/institutions">Institutions</a></li>
      <li><a href="/admin/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></li>
      <li class="active">{{ .program.GetTitle }}</li>
    </ul></div>
    <div class="panel-body">
      <a href="/admin/institutions/{{ .institution.ID }}/program/{{ .program.GetID }}/module/add" class=" btn-sm btn btn-primary">
        <span class="glyphicon glyphicon-plus" ></span>
        Add module
      </a>
      <br />
      <br />
      {{ if eq (not .modules) true }}
      <div class="text-muted">
        :( there are no moduels for this program of study
      </div>
      {{ else }}
      <h4>Modules ({{ .modules | len }})</h4>
      <div class="table-responsive">
        <table class="table table-striped table-hover " id="modules-list">
          <thead>
            <tr>
              <th>Code</th>
              <th>Title</th>
              <th>Content</th>
              <th>CATS</th>
            </tr>
          </thead>
          <tbody>
            {{ $institution   := .institution }}
            {{ $program       := .program }}
            {{range $index, $module := .modules}}
            <tr>
              <td>{{ $module.GetCode }}</td>
              <td>{{ $module.GetTitle }}</td>
              <td>{{ $module.GetContent }}</td>
              <td>{{ $module.GetCATS }}</td>
              <td><a href="/admin/institutions/{{ $institution.ID }}/program/{{ $program.GetID }}/module/{{ $module.GetID }}/modify">Edit</td>
                <td><a href="#" data-id="{{ $module.GetID }}" class="deleteModuleButton btn btn-danger btn-xs" >Delete</a></td>
              </tr>
              {{ end }}
            </tbody>
          </table>
        </div>
        {{ end }}
      </div>
    </div>
    <div>
    <style scoped>
    .big-number {
      font-size: 30px;
    }
    </style>
  </div>
  <script src="/static/js/admin/institutions/institution/program.js"></script>
  <script>

  var institution = {
    id: {{ .institution.ID }},
    program : {
      id: {{ .program.GetID }},
    },
  };

  console.log(institution)

  $(document).ready(function(){
    var module = wisply.getModule("admin-institutions-program"),
    manager = new module.Manager(institution);
    manager.init();
  });
  </script>
