<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li class="active">Institutions
              {{ if .currentAccount.IsAdministrator  }}
              <a href="/admin/institutions"><span class="label label-default">Admin this</span></a>
              {{ end }}
            </li>
          </ul>
        </div>
        <div class="panel-body">
          <h1>Institutions</h1>
          <div class="well">
            Wisply is proudly collecting data from these prestigious institutions
          </div>
          {{ if .anything }}
          <div class="row text-center">
            {{range $index, $institution := .institutions}}
            {{$safe := $institution.Name|html}}
            <div class="text-center col-xs-12 col-sm-6 col-md-4  col-lg-3">
              <a title="{{ $safe }}" href="/institutions/{{ $institution.ID }}" class="">
                <div class="institution-profile thumbnail">
                  <div class="insider">
                    {{ if eq $institution.LogoURL "" }}
                    <span class="glyphicon glyphicon-education institution-logo-default"></span>
                    {{ else }}
                    <img alt="{{ $institution.Name }}" src="{{ $institution.LogoURL }}" class="inlogo text-center" />
                    {{ end }}
                  </div>
                  <div class="caption">
                    {{ $safe}}
                  </div>
                </div>
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
<link href="/static/css/public/institution.css" property='stylesheet' type="text/css" rel="stylesheet" />
<script src="/static/js/admin/institution/list.js"></script>
