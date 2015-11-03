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
          <label for="ka-title" class="col-lg-2 control-label">Title</label>
          <div class="col-lg-10">
            <input type="text" value="{{ .ka.GetTitle }}" class="form-control" name="ka-title" id="ka-title" placeholder="The name of knowledge area" required pattern=".{3,100}" title="The code has 3 up to 100 characters">
          </div>
        </div>
        <div class="form-group">
          <label for="ka-code" class="col-lg-2 control-label">Code</label>
          <div class="col-lg-10">
            <input type="text" value="{{ .ka.GetCode }}" class="form-control" name="ka-code" id="ka-code" placeholder="A short code which represents the KA" required pattern=".{2,10}" title="The code has 2 up to 10 characters">
          </div>
        </div>
        <div class="form-group">
          <label for="ka-source" class="col-lg-2 control-label">Source</label>
          <div class="col-lg-10">
            <input type="text" value="{{ .ka.GetSource }}" class="form-control" name="ka-source" id="ka-source" placeholder="The source may be: ACM Computer Science Curriculum 2013" required pattern=".{0,200}" title="The source has up to 200 characters">
          </div>
        </div>
        <div class="form-group">
          <label for="ka-content" class="col-lg-2 control-label">Description</label>
          <div class="col-lg-10">
            <textarea class="form-control" name="ka-content" id="ka-content" placeholder="Content">{{ .ka.GetContent }}</textarea>
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
  $("#ka-title").focus();
});
</script>
