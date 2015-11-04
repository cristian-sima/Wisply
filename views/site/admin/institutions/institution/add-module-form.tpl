<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/education">Education</a></li>
      <li><a href="/admin/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></li>
      <li><a href="/admin/institutions/{{ .institution.ID }}/program/{{ .program.GetID }}">{{ .program.GetTitle }}</a></li>
      <li class="active">Add module</li>
    </ul>
  </div>
  <div class="panel-body">
    {{ if eq (.modulesToAdd | len) 0 }}
    There are no modules to add. <a href="/admin/institutions/{{ .institution.ID }}/program/{{ .program.GetID }}">Go back</a>
    {{ else }}
    <form method="POST" class="form-horizontal" >
      {{ .xsrf_input }}
      <fieldset>
        <div class="form-group">
          <label for="program-year" class="col-lg-2 control-label">Module</label>
          <div class="col-lg-10">
            <select class="form-control" id="module-id" name="module-id">
              {{ range $index, $module := .modulesToAdd }}
              <option value="{{ $module.GetID }}">
                {{ $module.GetCode }} - {{$module.GetTitle }}
              </option>
              {{ end }}
            </select>
          </div>
        </div>
        <div class="form-group">
          <div class="col-lg-10 col-lg-offset-2">
            <input type="submit" class="btn btn-primary" value="Add module" /> <a href="/admin/institutions/{{ .institution.ID }}/program/{{ .program.GetID }}" class="btn btn-default">Cancel</a>
          </div>
        </div>
      </fieldset>
    </form>
    {{ end }}
  </div>
</div>
<script>
$(document).ready(function() {
  $("#module-title").focus();
});
</script>
