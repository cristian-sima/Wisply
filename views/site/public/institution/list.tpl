<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li class="active">Institutions</li>
          </ul>
        </div>
        <div class="panel-body">
          <blockquote>
            <p>Wisply is proudly collecting data from these prestigious institutions</p>
          </blockquote>
            {{ if .anything }}
            <div class="row text-center">
              {{range $index, $institution := .institutions}}
              {{$safe := $institution.Name|html}}
              <div class="text-center col-xs-3 col-md-3">
                <a title="{{ $safe }}" href="/institutions/{{ $institution.ID }}" class="thumbnail">
                  {{ if eq $institution.LogoURL "" }}
                  <span class="glyphicon glyphicon-education institution-logo"></span>
                  {{ else }}
                  <img src="{{ $institution.LogoURL }}" />
                  {{ end }}
                  <br />
                  {{ $safe }}
                </a>
              </div>
              {{end }}
            </div>
            {{ else }}
            There are no institution... :(
            {{ end }}
          </div>
        </div>
      </div>
    </div>
  </div>
  <style>
  .institution-logo {
    font-size: 5em;

  }
  </style>
  <script src="/static/js/admin/institution/list.js"></script>
