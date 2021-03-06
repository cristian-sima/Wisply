<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/institutions">Institutions</a></li>
            <li><a href="/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></li>
            <li class="active">{{ .repository.Name }}
              {{ if .currentAccount.IsAdministrator  }}
              <a href="/admin/repositories/{{ .repository.ID }}"><span class="label label-default">Admin this</span></a>
              {{ end }}
            </li>
          </ul>
        </div>
        <div class="panel-body">
          <div class="row">
            <div class="col-lg-2 col-md-2 col-sm-2 text-center" >
              <span class="glyphicon glyphicon-equalizer institution-logo-default "></span>
              <div class="text-left"></div>
              <hr />
              <div class="text-left">
                <a target="_blank" href="{{ .repository.PublicURL }}">Web page</a>
                <div>{{ .repository.Category }} repository</div>
              </div>
            </div>
            <div class="col-lg-6 col-md-6 col-sm-6" >
              <div>
                <h1>{{ .repository.Name }}</h1>
                <span class="text-muted">Repository</span>
              </div>
              <div>
                {{ .repository.Description }}
              </div>
            </div>
            <div class="col-lg-4 col-md-4 col-sm-4" >
              <div>Part of <a href="/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></div>
            </div>
          </div>
          <div>
            <br />
            <br />
          </div>
          <!-- Things -->
          {{ if eq .repository.LastProcess 0 }}
          <div class="text-center text-muted">
            Wisply did not process this repository yet... :(
          </div>
          {{ else }}
          <div class="row">
            <!-- Statistics -->
            <div class="col-lg-2 col-md-2 col-sm-2 text-left" >

            </div>
            <!-- Resources -->
            <div class="col-lg-6 col-md-6 col-sm-6" >
              <div id="repository-top"></div>
              <div id="repository-before-resources"></div>
              <div id="repository-resources"></div>
              <ul class="pager">
                <li class="previous" style="display:none"><a href="#">← Previous</a></li>
                <li class="next" style="display:none"><a href="#">Next →</a></li>
              </ul>
              <div id="repository-bottom"></div>
            </div>
            <div class="col-lg-4 col-md-4 col-sm-4" >
              <!-- Collections -->
              <div class="list-group" id="repository-side"></div>
            </div>
          </div>
          {{ end }}
        </div>
      </div>
    </div>
  </div>
</div>
<div>
<style scoped>
.text-almost-invisible {
  color:#D8D8D8;
}
.change-resources-per-page {
  width: 50px;
  background-color: #f9f9f9;
}
a.resource:focus{
  outline: 0px auto transparent;
}
</style>
</div>
<script>
var server = {};
server.repository = {
  id : {{ .repository.ID }},
  totalResources: {{ .process.Records }},
  name: "{{ .repository.Name }}",
  collections: JSON.parse({{ .collectionsJSON }}),
};
server.repository.getCollection = function (requestedID) {
  var i, collection;
  for(i=0; i < this.collections.length; i++) {
    collection = this.collections[i];
    if(parseInt(collection.ID, 10) === parseInt(requestedID, 10) ) {
      return collection;
    }
  }
  return false;
};
server.repository.getBySpec = function (requestedSet) {
  var i, collection;
  for(i=0; i < this.collections.length; i++) {
    collection = this.collections[i];
    if(collection.Spec === requestedSet ) {
      return collection;
    }
  }
  return false;
};
</script>
<link href="/static/css/public/institution.css" type="text/css" rel="stylesheet" property='stylesheet' />
<script src="/static/js/public/repositories/repository/home.js"></script>
<script>
$(document).ready(function(){
    var module = wisply.getModule("public-repositories-repository"),
      manager = new module.Manager(server.repository);
      manager.init();
});
</script>
