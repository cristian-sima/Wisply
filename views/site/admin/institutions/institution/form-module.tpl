<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/education">Education</a></li>
      <li><a href="/admin/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></li>
      <li><a href="/admin/institutions/{{ .institution.ID }}/program/{{ .program.GetID }}">{{ .program.GetTitle }}</a></li>
      <li class="active">{{ .action }}</li>
    </ul>
  </div>
  <div class="panel-body">
    <form method="POST" class="form-horizontal" >
      {{ .xsrf_input }}
      <fieldset>
        <div class="form-group">
          <label for="module-title" class="col-lg-2 control-label">Title</label>
          <div class="col-lg-10">
            <input type="text" value="{{ .module.GetTitle }}" class="form-control" name="module-title" id="module-title" placeholder="Title" required pattern=".{3,200}" title="The code has 3 up to 200 characters">
          </div>
        </div>
        <div class="form-group">
          <label for="module-code" class="col-lg-2 control-label">Code</label>
          <div class="col-lg-10">
            <input type="text" value="{{ .module.GetCode }}" class="form-control" name="module-code" id="module-code" placeholder="A short code which represents the module" required pattern=".{2,10}" title="The code has 2 up to 10 characters">
          </div>
        </div>
        <div class="form-group">
          <label for="module-year" class="col-lg-2 control-label">Year (e.g. 1 or 6)</label>
          <div class="col-lg-10">
            <input type="text" value="{{ .module.GetYear }}" class="form-control" name="module-year" id="module-year" placeholder="Year" required pattern=".{1,2}" title="The year contains between one to two characters">
          </div>
        </div>
        <div class="form-group">
          <label for="module-CATS" class="col-lg-2 control-label">CATS</label>
          <div class="col-lg-10">
            <input type="text" value="{{ .module.GetCATS }}" class="form-control" name="module-CATS" id="module-CATS" placeholder="CATS" required pattern=".{0,5}" title="The CATS field has up to 5 characters">
          </div>
        </div>
        <div class="form-group">
          <label for="module-content" class="col-lg-2 control-label">Description</label>
          <div class="col-lg-10">
            <textarea class="form-control" rows="5" name="module-content" id="module-content"> {{ .module.GetContent }}</textarea>
          </div>
        </div>
        <div class="form-group">
          <div class="col-lg-10 col-lg-offset-2">
            <input type="submit" class="btn btn-primary" value="{{ .action }}" /> <a href="/admin/institutions/{{ .institution.ID }}/program/{{ .program.GetID }}" class="btn btn-default">Cancel</a>
          </div>
        </div>
      </fieldset>
    </form>
  </div>
</div>
<script>
$(document).ready(function() {
  $("#module-title").focus();
});
</script>
