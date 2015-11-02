<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/education">Education</a></li>
      <li class="active">{{ .program.GetName }}</li>
    </ul>
  </div>
  <div class="panel-body">
    <span class="h4">{{ .program.GetName }}</span>
    <hr />
    <div>
    <a href="/admin/education/{{ .program.GetID }}/add-knowledge-area" class="btn btn-primary sm">Add Knowledge Area</a>
    </div>
    <hr />
    <div>
      <a class="btn btn-primary" href="/admin/education/programs/{{ .program.GetID }}/advance-options">Advance options</a>
    </div>
  </div>
</div>
