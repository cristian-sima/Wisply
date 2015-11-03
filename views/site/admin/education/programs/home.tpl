<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/education">Education</a></li>
      <li class="active">{{ .program.GetName }}</li>
    </ul>
  </div>
  <div class="panel-body">
    <h2>{{ .program.GetName }}</h2>
    <hr />
    <div>
    <a href="/admin/education/programs/{{ .program.GetID }}/definition/add" class="btn btn-primary sm">Add formal definition</a>
    <!--<a href="/admin/education/programs/{{ .program.GetID }}/knowledge-area/add" class="btn btn-primary sm">Add Knowledge Area</a>-->
    </div>
    <h3>Formal definitions</h3>
    {{ if eq (.definitions | len) 0 }}
    There are no formal definitions
    {{ else }}
    <div class="table-responsive">
      <table id="list-accounts" class="table table-striped table-hover ">
        <thead>
          <tr>
            <th>Content</th>
            <th>Source</th>
            <th>Edit</th>
            <th>Delete</th>
          </tr>
        </thead>
        <tbody>
          {{ $program := .program }}
          {{range $index, $definition := .definitions }}
          <tr>
            <td>{{ $definition.GetContent }}</td>
            <td>{{ $definition.GetSource }}</td>
            <td><a href="/admin/education/programs/{{ $program.GetID }}/definition/{{ $definition.GetID }}/modify">Edit</a></td>
            <td><a href="#" data-id="{{ $definition.GetID }}" class="deleteDefinitionButton btn btn-danger btn-xs" >Delete</a></td>
          </tr>
          {{ end }}
        </tbody>
      </table>
    </div>
    {{ end }}
    <hr />
    <div>
      <a class="btn btn-primary" href="/admin/education/programs/{{ .program.GetID }}/advance-options">Advance options</a>
    </div>
  </div>
</div>
<script src="/static/js/admin/education/program/home.js"></script>
<script src="/static/js/wisply/server.js"></script>
<script>

var program = {
  id: {{ .program.GetID }},
};

$(document).ready(function(){
    var module = wisply.getModule("admin-education-program-home"),
      manager = new module.Manager(program);
      manager.init();
});
</script>
