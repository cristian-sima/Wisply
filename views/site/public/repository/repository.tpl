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
              <a href="/admin/repositories/repository/{{ .repository.ID }}"><span class="label label-default">Admin this</span></a>
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
                <span class="text-muted">Address:</span> <a target="_blank" href="{{ .repository.PublicURL }}">Web page</a>
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
              <div><i>{{ .repository.Category }}</i> repository</div>
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
              <table class="table">
                <tbody>
                  <tr>
                    <td>
                      <span class="badge badge-info">{{ .process.Collections }}</span>
                    </td>
                    <td>
                      collections
                    </td>
                  </tr>
                  <tr>
                    <td>
                      <span class="badge badge-info">{{ .process.Records }}</span>
                    </td>
                    <td>
                      records
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
            <!-- Resources -->
            <div class="col-lg-6 col-md-6 col-sm-6" >
              <div id="repository-top">

              </div>
              <div id="repository-before-resources">
              </span>
            </div>
              <div id="repository-resources"></div>
              <ul class="pager">
                <li class="previous" style="display:none"><a href="#">← Older</a></li>
                <li class="next" style="display:none"><a href="#">Newer →</a></li>
              </ul>
              <div id="repository-bottom"></div>
            </div>
            <div class="col-lg-4 col-md-4 col-sm-4" >
              <!-- Collections -->
              <div class="list-group" id="repositories">
                {{range $index, $collection := .collections}}
                <a data-id="{{ $collection.ID }}" class="hover list-group-item set-collection">
                  <h5 class="list-group-item-heading"><span class="glyphicon glyphicon-equalizer"></span> {{ $collection.Name }} <span class="badge">{{ $collection.NumberOfResources }}</span></h5>
                  <p class="list-group-item-text">{{ $collection.Description }}</p>
                </a>
                {{ end }}
              </div>
            </div>
          </div>
          {{ end }}
        </div>
      </div>
    </div>
  </div>
</div>
<style>
.text-almost-invisible {
  color:#D8D8D8;
}
</style>
<script src="/static/js/admin/repository/list.js"></script>
<script>
var server = {};
server.repository = {
  id : {{ .repository.ID }},
  totalResources: {{ .process.Records }},
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
</script>
<link href="/static/css/public/institution.css" type="text/css" rel="stylesheet" property='stylesheet' />
<script src="/static/js/public/repository.js"></script>
