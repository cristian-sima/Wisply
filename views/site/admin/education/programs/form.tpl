<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/education">Education</a></li>
      {{ if eq .action "Modify" }}
      <li><a href="/admin/education/programs/{{ .program.GetID }}">{{ .program.GetName }}</a></li>
      <li><a href="/admin/education/programs/{{ .program.GetID }}/advance-options">Advance options</a></li>
      {{ end }}
      <li class="active">{{ .action }}</li>
    </ul>
  </div>
  <div class="panel-body">
    <form method="POST" class="form-horizontal" >
      {{ .xsrf_input }}
      <fieldset>
        <div class="form-group">
          <label for="program-name" class="col-lg-2 control-label">Name</label>
          <div class="col-lg-10">
            <input type="text" value="{{ .program.GetName }}" class="form-control" name="program-name" id="program-name" placeholder="The name of the program" required pattern=".{3,255}" title="The name has 3 up to 300 characters!">
          </div>
        </div>
        <div class="form-group">
          <div class="col-lg-10 col-lg-offset-2">
            {{ if eq .action "Add"}}
            <input type="submit" class="btn btn-primary" value="Add" /> <a href="/admin/education" class="btn btn-default">Back to list</a>
            {{ else }}
            <input type="submit" class="btn btn-primary" value="Modify" /> <a href="/admin/education/programs/{{ .program.GetID }}/advance-options" class="btn btn-default">Cancel</a>
            {{ end }}
          </div>
        </div>
      </fieldset>
    </form>
  </div>
</div>
<script>
$(document).ready(function() {
  $("#program-name").focus();
});
</script>
