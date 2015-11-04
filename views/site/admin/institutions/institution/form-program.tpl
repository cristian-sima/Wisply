<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/education">Education</a></li>
      <li><a href="/admin/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></li>
      <li class="active">{{ .action }}</li>
    </ul>
  </div>
  <div class="panel-body">
    <form method="POST" class="form-horizontal" >
      {{ .xsrf_input }}
      <fieldset>
        <div class="form-group">
          <label for="program-title" class="col-lg-2 control-label">Title</label>
          <div class="col-lg-10">
            <input type="text" value="{{ .program.GetTitle }}" class="form-control" name="program-title" id="program-title" placeholder="Title" required pattern=".{3,200}" title="The code has 3 up to 200 characters">
          </div>
        </div>
        <div class="form-group">
          <label for="program-level" class="col-lg-2 control-label">Year</label>
          <div class="col-lg-10">
            <select class="form-control" id="program-level" name="program-level">
              <option value="undergraduate" {{ if eq .program.GetLevel "undergraduate"}}selected {{ end }}>Undergraduate</option>
              <option value="postgraduate" {{ if eq .program.GetLevel "postgraduate"}}selected {{ end }}>Postgraduate</option>
            </select>
          </div>
        </div>
        <div class="form-group">
          <label for="program-code" class="col-lg-2 control-label">Code</label>
          <div class="col-lg-10">
            <input type="text" value="{{ .program.GetCode }}" class="form-control" name="program-code" id="program-code" placeholder="A short code which represents the Program" required pattern=".{2,10}" title="The code has 2 up to 10 characters">
          </div>
        </div>
        <div class="form-group">
          <label for="program-program" class="col-lg-2 control-label">Category</label>
          <div class="col-lg-10">
            <select class="form-control" id="program-program" name="program-program">
              {{ range $index, $program := .subjects }}
              <option value="{{ $program.GetID }}">{{ $program.GetName }}</option>
              {{ end }}
            </select>
          </div>
        </div>
        <div class="form-group">
          <label for="program-year" class="col-lg-2 control-label">Year</label>
          <div class="col-lg-10">
            <select class="form-control" id="program-year" name="program-year">
              <option>2015</option>
              <option>2014</option>
              <option>2013</option>
              <option>2012</option>
              <option>2011</option>
              <option>2010</option>
            </select>
          </div>
        </div>
        <div class="form-group">
          <label for="program-content" class="col-lg-2 control-label">Description</label>
          <div class="col-lg-10">
            <textarea class="form-control" rows="5" name="program-content" id="program-content"> {{ .program.GetContent }}</textarea>
          </div>
        </div>
        <div class="form-group">
          <label for="program-ucas-code" class="col-lg-2 control-label">UCAS code</label>
          <div class="col-lg-10">
            <input type="text" value="{{ .program.GetCode }}" class="form-control" name="program-ucas-code" id="program-ucas-code" placeholder="A short code which represents the Program" pattern=".{0,20}" title="The ucas-code has 0 up to 20 characters">
          </div>
        </div>
        <div class="form-group">
          <div class="col-lg-10 col-lg-offset-2">
            <input type="submit" class="btn btn-primary" value="{{ .action }}" /> <a href="/admin/institutions/{{ .institution.ID }}#programs" class="btn btn-default">Cancel</a>
          </div>
        </div>
      </fieldset>
    </form>
  </div>
</div>
<script>
$(document).ready(function() {
  $("#program-title").focus();
});
</script>
