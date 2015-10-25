<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/repositories">Repositories</a></li>
      {{ if eq .action "Modify" }}
      <li><a href="/admin/repositories/repository/{{ .repository.ID }}">{{ .repository.Name }}</a></li>
      <li><a href="/admin/repositories/repository/{{ .repository.ID }}/advance-options">Advance options</a></li>
      {{ else }}
      <li><a href="/admin/repositories/add?institution={{.selectedInstitution}}">Choose Type</a></li>
      {{ end }}
      <li class="active">{{ .action }}</li>
    </ul>
  </div>
  <div class="panel-body">
  </div>
  <form method="{{.actionType}}" class="form-horizontal" >
    <input type="hidden" name="repository-category" value="{{ .category }}"/>
    {{ .xsrf_input }}
    {{ $safeDescription := .repository.Description|html}}
    <fieldset>
      <div class="form-group">
        <label class="col-lg-2 control-label">Category</label>
        <div class="col-lg-10">
          <strong>{{ .category }}</strong>
        </div>
      </div>
      <div class="form-group">
        <label for="repository-name" class="col-lg-2 control-label">Name</label>
        <div class="col-lg-10">
          <input type="text" value="{{.repository.Name}}" class="form-control" name="repository-name" id="repository-name" placeholder="Name" required pattern=".{3,255}" title="The name has 3 up to 255 characters!">
        </div>
      </div>
      <div class="form-group">
        <label for="repository-URL" class="col-lg-2 control-label">Base URL</label>
        <div class="col-lg-10">
          <input type="url" value="{{.repository.URL}}" class="form-control" name="repository-URL" id="repository-URL" placeholder="http://address.domain" required pattern=".{3,2083}" title="The URL has 3 up to 2083 characters!">
        </div>
      </div>
      {{ if eq .action "Add" }}
      <div class="form-group">
        <label for="institution-URL" class="col-lg-2 control-label">Institution <a href="/admin/institutions/add" target="_blank"><span data-toggle="tooltip" data-placement="top" title="Create institution" class="glyphicon glyphicon-plus-sign text-success"> </span></a></label>
        <div class="col-lg-10">
          <select class="form-control" name="repository-institution" id="repository-institution">
            {{ $selected := .selectedInstitution }}
            {{ $repInstitution := .repository.Institution }}
            {{range $index, $institution := .institutions}}
            {{$safe := $institution.Name|html}}
            <option
            {{ if $selected }}
            {{ if eq $selected $institution.ID }}
            selected
            {{ end }}
            {{ end }}
            value="{{ $institution.ID }}">{{ $safe }}</option>
            {{ end }}
          </select>
        </div>
      </div>
      <div class="form-group">
        <label for="repository-public-url" class="col-lg-2 control-label">Public URL</label>
        <div class="col-lg-10">
          <input type="url" value="{{.repository.PublicURL}}" class="form-control" name="repository-public-url" id="repository-public-url" placeholder="http://address.domain" required pattern=".{3,2083}" title="The URL has 3 up to 2083 characters!">
        </div>
      </div>
      {{ end }}
      <div class="form-group">
        <label for="repository-description" class="col-lg-2 control-label">Description</label>
        <div class="col-lg-10">
          <textarea class="form-control" rows="3" name="repository-description" id="repository-description" maxlength="500" >{{ .repository.Description}}</textarea>
          <span class="help-block">This field may contain notes about the intitution.</span>
        </div>
      </div>
      <div class="form-group">
        <div class="col-lg-10 col-lg-offset-2">
          {{ if eq .action "Add"}}
          <input type="submit" class="btn btn-primary" value="Submit" /> <a href="/admin/repositories" class="btn btn-default">Back to list</a>
          {{ else }}
          <input type="submit" class="btn btn-primary" value="Submit" /> <a class="btn btn-primary" href="/admin/repositories/repository/{{ .repository.ID }}/advance-options" class="btn btn-default">Cancel</a>
          {{ end }}
        </div>
      </div>
    </fieldset>
  </form>
</div>
</div>
<script>
$(document).ready(function() {
  $("#repository-name").focus();
  wisply.activateTooltip();
});
</script>
