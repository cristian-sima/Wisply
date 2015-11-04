<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/education">Education</a></li>
      <li class="active">{{ .subject.GetName }}</li>
    </ul>
  </div>
  <div class="panel-body">
    <h2>{{ .subject.GetName }}</h2>
    <ul class="nav nav-tabs">
      <li class="active"><a href="#KAs" data-toggle="tab">Knowledge Areas</a></li>
      <li><a href="#definitions" data-toggle="tab">Definitions</a></li>
    </ul>
    <div id="myTabContent" class="tab-content">
      <div class="tab-pane fade active in" id="KAs">
        <a href="/admin/education/subjects/{{ .subject.GetID }}/ka/add" class="btn btn-sm btn-success">
        <span class="glyphicon glyphicon-plus" ></span> Add Knowledge area</a>
        <br />
        <br />
        {{ if eq (.KAs | len) 0 }}
        There are no knowledge areas.
        {{ else }}
        <div class="table-responsive">
          <table id="list-accounts" class="table table-striped table-hover ">
            <thead>
              <tr>
                <th>Title</th>
                <th>Edit</th>
                <th>Delete</th>
              </tr>
            </thead>
            <tbody>
              {{ $subject := .subject }}
              {{range $index, $ka := .KAs }}
              <tr>
                <td>{{ $ka.GetTitle }}</td>
                <td><a href="/admin/education/subjects/{{ $subject.GetID }}/ka/{{ $ka.GetID }}/modify">Edit</a></td>
                <td><a href="#" data-id="{{ $ka.GetID }}" class="deleteKAButton btn btn-danger btn-xs" >Delete</a></td>
              </tr>
              {{ end }}
            </tbody>
          </table>
        </div>
        {{ end }}
      </div>
      <div class="tab-pane fade" id="definitions">
        <a href="/admin/education/subjects/{{ .subject.GetID }}/definition/add" class="btn btn-sm btn-success">
          <span class="glyphicon glyphicon-plus"></span>
          Add formal definition
        </a>
        <br />
        <br />
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
              {{ $subject := .subject }}
              {{range $index, $definition := .definitions }}
              <tr>
                <td>{{ $definition.GetContent }}</td>
                <td>{{ $definition.GetSource }}</td>
                <td><a href="/admin/education/subjects/{{ $subject.GetID }}/definition/{{ $definition.GetID }}/modify">Edit</a></td>
                <td><a href="#" data-id="{{ $definition.GetID }}" class="deleteDefinitionButton btn btn-danger btn-xs" >Delete</a></td>
              </tr>
              {{ end }}
            </tbody>
          </table>
        </div>
      </div>
    {{ end }}
  </div>
    <hr />
    <div>
      <a class="btn btn-primary" href="/admin/education/subjects/{{ .subject.GetID }}/advance-options">Advance options</a>
    </div>
  </div>
</div>
<script src="/static/js/admin/education/subject/home.js"></script>
<script>

var subject = {
  id: {{ .subject.GetID }},
};

$(document).ready(function(){
  var module = wisply.getModule("admin-education-subject-home"),
  manager = new module.Manager(subject);
  manager.init();
});
</script>
