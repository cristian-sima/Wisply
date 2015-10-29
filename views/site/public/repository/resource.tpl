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
          <div class="top-info">
          </div>
          <div class="content-info">
            <div class="embed-responsive embed-responsive-16by9">
              <div id="the-iframe" class="embed-responsive-item the-iframe" >
              </div>
            </div>
            {{ .record }}
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
