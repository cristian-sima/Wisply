<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/institutions">Institutions</a></li>
            <li><a href="/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></li>
            <li><a href="/repositories/{{ .repository.ID }}">{{ .repository.Name }}</a></li>
            <li class="active">{{ .resource.Keys.GetTitle }}</li>
          </ul>
        </div>
        <div class="panel-body">
          <h1>{{ .resource.Keys.GetTitle }}</h1>
          <div class="h6 text-muted">
            {{range $index, $creator := .resource.Keys.Get "creator" }}
            {{ $creator }}
            {{ end }}
          </div>
          <div class="top-info">
            {{range $index, $description := .resource.Keys.Get "description" }}
            {{ $description }}
            {{ end }}
          </div>
          <div class="row">
            <div class="col-md-12">
              <table class="table h5">
                <tbody>
                  <tr>
                    <td> <span class="glyphicon glyphicon-education"></span> <a href="/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></td>
                  </tr>
                  <tr>
                    <td> <span class="glyphicon glyphicon glyphicon-equalizer "></span> <a href="/repositories/{{ .repository.ID }}">{{ .repository.Name }}</a></td>
                  </tr>
                  {{ if .module }}
                  <tr>
                    <td> <span class="glyphicon glyphicon-random"></span> <span class="text-muted"> Identified as part of</span> <a href="/institutions/{{ .module.GetInstitution }}/module/{{ .module.GetID }}">{{ .module.GetTitle }}</a>
                    </td>
                  </tr>
                  {{ end }}
                </tbody>
              </table>
            </div>
            <div class="col-md-6">
              <!-- Other information -->
            </div>
          </div>
          <div class="content-info">
            {{ if not .resource.IsVisible }}
            <div class="well">
              <div class="row">
                <div class="col-md-1 text-center">
                  <span class="h3"><span class="glyphicon glyphicon-lock"></span></span>
                </div>
                <div class="col-md-11">
                  Wisply does not have access to the content of this resource.
                  <br />
                  This may happen because this resource can not be previewed in the browser or that the author of the resource has maked it as private. <br />
                  <a href="{{ .resource.Keys.GetURL }}"><span class="glyphicon glyphicon-blackboard"></span> See it on {{ .repository.Name }}</a>                </div>
                </div>
              </div>
              {{ else }}
              <div class="well well-sm">
                <a class="download-file" href="http://www.edshare.soton.ac.uk/id/document/289262" >
                  <!-- <span><span class="glyphicon glyphicon-download"></span> Download</a> &bull;</span> -->
                  <a href="{{ .resource.Keys.GetURL }}"><span class="glyphicon glyphicon-blackboard"></span> See it on {{ .repository.Name }}</a>
                </div>
                <div class="row text-center">
                  <div class="col-md-12 text-center">
                    <div id="resource-content text-center" >
                      <!--<img class="text-center img-responsive" src="http://www.edshare.soton.ac.uk/15322/6/page.jpg" />-->
                    </div>
                  </div>
                </div>
                {{ end }}
              </div>
              <br />
              <div class="panel panel-default">
                <div class="panel-heading">Suggestions</div>
                <div class="panel-body">
                  {{ if not .module }}
                  <span class="text-muted">There are no suggestions for this resource</span>
                  {{ else }}
                  {{ range $index, $resource := .resourcesSuggested }}
                  <a class="resource" href="{{ $resource.GetWisplyURL }}">
                    <h4>
                      {{ if not $resource.IsVisible }}
                      <small><span data-toggle='tooltip' title='This content is not visible to Wisply.' class='glyphicon glyphicon-lock'></span></small>
                      {{ end }}
                      {{range $index, $title := $resource.Keys.Get "title" }}
                      {{ $title }}
                      {{ end }}
                    </h4>
                  </a>
                  {{ end }}
                  {{ end }}
                </div>
              </div>
              <div class="panel panel-default">
                <div class="panel-heading">Resources from the same repository</div>
                <div class="panel-body" id="div-same-collection">

                </div>
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
