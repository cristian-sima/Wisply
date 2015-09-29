<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/institutions">Institutions</a></li>
      <li class="active">{{.action}}</li>
    </ul></div>
    <div class="panel-body">
        <form method="{{.actionType}}" class="form-horizontal" >
          {{ .xsrf_input }}
          {{ $safeDescription := .institutionDescription|html}}
          <fieldset>
            <div class="form-group">
              <label for="institution-name" class="col-lg-2 control-label">Name</label>
              <div class="col-lg-10">
                <input type="text" value="{{.institutionName}}" class="form-control" name="institution-name" id="institution-name" placeholder="Name" required pattern=".{3,255}" title="The name has 3 up to 255 characters!">
              </div>
            </div>
            {{ if eq .action "Add" }}
            <div class="form-group">
              <label for="institution-URL" class="col-lg-2 control-label">Base URL</label>
              <div class="col-lg-10">
                <input type="url" value="{{.institutionUrl}}" class="form-control" name="institution-URL" id="institution-URL" placeholder="URL address" required pattern=".{3,2083}" title="The URL has 3 up to 2083 characters!">
              </div>
            </div>
            {{ end }}
            <div class="form-group">
              <label for="institution-description" class="col-lg-2 control-label">Description</label>
              <div class="col-lg-10">
                <textarea class="form-control" rows="3" name="institution-description" id="institution-description" maxlength="255" >{{ .institutionDescription}}</textarea>
                <span class="help-block">This field may contain notes about the intitution.</span>
              </div>
            </div>
            <div class="form-group">
              <div class="col-lg-10 col-lg-offset-2">
                <input type="submit" class="btn btn-primary" value="Submit" /> <a href="/admin/institutions" class="btn btn-default">Cancel</a>
              </div>
            </div>
          </fieldset>
        </form>
    </div>
  </div>
  <script>
  $(document).ready(function() {
    $("#institution-name").focus();
  });
  </script>
