<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/curriculum">Programs of study</a></li>
      <li class="active">{{ .action }}</li>
    </ul>
  </div>
  <div class="panel-body">
  </div>
  <form method="POST" class="form-horizontal" >
    {{ .xsrf_input }}
    <input type="hidden" name="program-id" value="{{ .program.GetID  }}"/>
    <fieldset>
      <div class="form-group">
        <label for="repository-name" class="col-lg-2 control-label">Name</label>
        <div class="col-lg-10">
          <input type="text" value="{{ .program.GetName }}" class="form-control" name="program-name" id="program-name" placeholder="The name of the program of study" required pattern=".{3,255}" title="The name has 3 up to 300 characters!">
        </div>
      </div>
      <div class="form-group">
        <div class="col-lg-10 col-lg-offset-2">
          {{ if eq .action "Add"}}
          <input type="submit" class="btn btn-primary" value="Add" /> <a href="/admin/curriculum" class="btn btn-default">Back to list</a>
          {{ else }}
          <input type="submit" class="btn btn-primary" value="Modify" /> <a href="/admin/curriculum/programs/{{ .program.GetID }}/advance-options" class="btn btn-default">Cancel</a>
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
  wisply.activateTooltip();
});
</script>
