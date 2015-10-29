<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/institutions">Institutions</a></li>
            <li><a href="/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></li>
            <li><a href="/repository/{{ .repository.ID }}">{{ .repository.Name }}</a></li>
            <li class="active">{{ .resource.Keys.GetTitle }}</li>
          </ul>
        </div>
        <div class="panel-body">
          <h1>{{ .resource.Keys.GetTitle }}</h1>
          <div class="top-info">
            {{range $index, $description := .resource.Keys.Get "description" }}
            {{ $description }}
            {{ end }}
          </div>
          <div class="row">
            <div class="col-md-3">
              <table class="table h5">
                <tbody>
                  <tr>
                    <td> <span class="glyphicon glyphicon-education"></span> <a href="/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></td>
                  </tr>
                  <tr>
                    <td> <span class="glyphicon glyphicon glyphicon-equalizer "></span> <a href="/repository/{{ .repository.ID }}">{{ .repository.Name }}</a></td>
                  </tr>
                </tbody>
              </table>
            </div>
            <div class="col-md-6">
              <!-- Other information -->
            </div>
          </div>
          <div class="content-info">
              <a href="http://bootswatch.com/paper/#" >
                <span><span class="glyphicon glyphicon-download"></span> Download</a> &bull;</span>
              <a href="{{ .resource.Keys.GetURL }}"><span class="glyphicon glyphicon-blackboard"></span> See it on {{ .repository.Name }}</a>

            {{ if not .resource.IsVisible }}
            <div class="well">
              <div class="row">
                <div class="col-md-1 text-center">
                  <span class="h3"><span class="glyphicon glyphicon-lock"></span></span>
                </div>
                <div class="col-md-11">
                  Wisply does not have access to the content of this resource.
                  <br />
                  This may happen because this resource can not be previewed in the browser or that the author of the resource maked it as private.
                </div>
              </div>
            </div>
            {{ else }}
            <div class="embed-responsive embed-responsive-16by9">
              <div id="the-iframe" class="embed-responsive-item the-iframe" >
              </div>
            </div>
            {{ end }}
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
<div>
<style scoped>
</style>
</div>
<script src="/static/js/public/resource.js"></script>
<script>
$(document).ready(function() {
  "use strict";
  var data = {
    repository : {
      name : "{{ .repository.Name }}",
      id: {{ .repository.ID }},
    },
    resource : {
      id: "{{ .resource.ID }}",
      identifier: "{{ .resource.Identifier }}",
    },
  },
  module = wisply.getModule("public-resource");

  module.init(data);
});
</script>
