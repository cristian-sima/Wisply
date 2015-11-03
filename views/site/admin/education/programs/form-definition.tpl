<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/education">Education</a></li>
      <li><a href="/admin/education/programs/{{ .program.GetID }}">{{ .program.GetName }}</a></li>
      <li class="active">{{ .action }}</li>
    </ul>
  </div>
  <div class="panel-body">
    <form method="POST" class="form-horizontal" >
      {{ .xsrf_input }}
      <fieldset>
        <div class="form-group">
          <label for="definition-source" class="col-lg-2 control-label">Source</label>
          <div class="col-lg-10">
            <input type="text" value="{{ .definition.GetSource }}" class="form-control" name="definition-source" id="definition-source" placeholder="The source of the definition" required pattern=".{0,200}" title="The source has up to 200 characters!">
          </div>
        </div>
        <div class="form-group">
          <label for="definition-content" class="col-lg-2 control-label">Content</label>
          <div class="col-lg-10">
            <textarea class="form-control" name="definition-content" id="definition-content" placeholder="Content">{{ .definition.GetContent }}</textarea>
          </div>
        </div>
        <div class="form-group">
          <div class="col-lg-10 col-lg-offset-2">
            <input type="submit" class="btn btn-primary" value="{{ .action }}" /> <a href="/admin/education/programs/{{ .program.GetID }}" class="btn btn-default">Cancel</a>
          </div>
        </div>
      </fieldset>
    </form>
  </div>
</div>
<script>
$(document).ready(function() {
  $("#definition-source").focus();
});
</script>
