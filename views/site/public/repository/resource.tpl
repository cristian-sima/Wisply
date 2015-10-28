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
            <li class="active">{{ .record.Keys.GetTitle }}</li>
          </ul>
        </div>
        <div class="panel-body">
          <h1>{{ .record.Keys }}</h1>
            <div class="top-info">
            </div>

            <div class="content-info">
              <div class="embed-responsive embed-responsive-16by9">

              <iframe id="the-iframe" class="embed-responsive-item the-iframe" >

              </iframe>
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
.the-iframe {

}
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
