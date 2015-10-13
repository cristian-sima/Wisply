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
            <div class="col-lg-2 col-md-2 col-sm-2 text-center" >

            </div>
            <!-- Resources -->
            <div class="col-lg-6 col-md-6 col-sm-6" >
            <table class="table">
                <tbody>
                    {{range $index, $record := .records}}







                    <tr>
                        <td>
                          <h4>
                            {{range $index, $title := $record.Keys.Get "title" }}
                            {{ $title }}
                            {{ end }}
                          </h4>
                          {{range $index, $description := $record.Keys.Get "description" }}
                          {{ $description }}
                          {{ end }}

                          <div class="creators"> By
                            <span class="text-muted">
                              {{range $index, $creator := $record.Keys.Get "creator" }}
                              {{ $creator }}
                              {{ end }}
                            </span>
                          </div>
                        </td>
                    </tr>
                    {{ end }}
                </tbody>
            </table>
            </div>
            <!-- Collections -->
            <div class="col-lg-4 col-md-4 col-sm-4" >

            </div>
          </div>

          {{ end }}
        </div>

      </div>
    </div>
  </div>
</div>
<link href="/static/css/public/institution.css" type="text/css" rel="stylesheet" property='stylesheet' />
<script src="/static/js/admin/institution/list.js"></script>
