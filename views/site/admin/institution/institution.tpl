<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/institutions">Institutions</a></li>
      <li class="active">{{ .institution.Name }}</li>
    </ul></div>
    <div class="panel-body">
      <div class="panel-body">
        <div class="row">
          <div class="col-lg-2 col-md-2 col-sm-2 text-center" >
            <div class="institution-profile">
              <div class="insider">
                {{ if eq .institution.LogoURL "" }}
                <span class="glyphicon glyphicon-education institution-logo-default"></span>
                {{ else }}
                <img alt="{{ .institution.Name }}" src="{{ .institution.LogoURL }}" class="inlogo" />
                {{ end }}
              </div>
            </div>
            <hr />
            <div class="text-left">
              <span class="text-muted">Address:</span> <a href="{{ .institution.URL }}">Web page</a>
            </div>
          </div>
          <div class="col-lg-6 col-md-6 col-sm-6" >
            <div>
              <h1>{{ .institution.Name }}</h1>
                <span class="text-muted">Institution</span>
              </div>
              <div>
                {{ .institution.Description}} <a href="/admin/institutions/modify/{{ .institution.ID}}">Modify</a>
              </div>
            </div>
            <div class="col-lg-4 col-md-4 col-sm-4" >
              <!-- Repositories -->
              <a href="/institutions/{{ .institution.ID}}">Public page</a>
              <br />
              <br />
              <a href="/admin/repositories/add/?institution={{ .institution.ID}}">
                <span class="glyphicon glyphicon-plus-sign text-success"> </span>
                Add repository
              </a>
              <br />
              <br />
              {{ if eq (not .repositories) true }}
              <div class="text-muted">
                :( it does not have repositories
              </div>
              {{ else }}
              <h4>Repositories ({{ .repositories | len }})</h4>
              <div class="list-group" id="repositories">
                {{range $index, $repository := .repositories}}
                <a href="/admin/repositories/repository/{{ $repository.ID }}" class="list-group-item">
                  <h4 class="list-group-item-heading"><span class="glyphicon glyphicon-equalizer"></span> {{ $repository.Name }}</h4>
                  <p class="list-group-item-text">{{ $repository.Description }}</p>
                </a>
                {{ end }}
              </div>
              {{ end }}
            </div>
          </div>
        </div>
        <hr />
        <div>
          <a class="btn btn-primary" href="/admin/institutions/institution/{{ .institution.ID }}/advance-options">Advance options</a>
        </div>
      </div>
    </div>
    <div>
      <style scoped>
      .big-number {
        font-size: 30px;
      }
      </style>
    </div>
    <script>
    $(document).ready(function(){
      $(".institution-status").each(function(){
        var el = $(this),
        status = wisply.institutionsModule.GUI.getStatusColor(el.html())
        el.html(status)
      });
      wisply.activateTooltip()
    });
    </script>
