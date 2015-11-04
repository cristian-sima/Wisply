<div class="page-header">
  <div class="row">
    <div class="col-lg-12 col-md-12 col-sm-12">
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/education">Subjects</a></li>
            <li class="active">{{ .subject.GetName }}</li>
          </ul>
        </div>
        <div class="panel-body">
          <h1>{{ .subject.GetName }}</h1>
          <br />
          <span class="text-warning glyphicon glyphicon-warning-sign"></span> Wisply was not able to generate data about <strong>{{ .subject.GetName }}</strong>'s curriculum from the institutions.
          <br />
          <br />
          {{ .subject.GetHTMLDescription }}
          <br />
          <div>
            {{ if ne ( .subject.GetDefinitions | len ) 0 }}
            <h2>Formal definitions of {{ .subject.GetName }}</h2>
            <div>
              {{ range $index, $definition := .subject.GetDefinitions }}
              <blockquote>
                <p>{{ $definition.GetContent }}</p>
                <small>Source <cite title="Source Title">{{ $definition.GetSource }}</cite></small>
              </blockquote>
              {{ end }}
            </div>
            {{ end }}
          </div>
          <div>
            {{ if ne ( .subject.GetKAs | len ) 0 }}
            <h2>Knowledge areas for {{ .subject.GetName }}</h2>
            <div>
              {{ range $index, $ka := .subject.GetKAs }}
              <div class="panel panel-default">
                <div class="panel-body">
                  <br />
                  <div class="row">
                    <div class="col-md-12">
                      <img align="left" style="margin:10px" class="thumbnail" src="/static/img/education/cs/ka/{{ $ka.GetCode }}.png" class="img-responsive"/>

                      <span class="h5"><strong>&nbsp;{{ $ka.GetTitle }}</strong></span>
                      <br />
                      {{ $ka.GetContent }}
                      <br />
                      <p class="text-muted text-right">{{ $ka.GetSource }}</p>
                    </div>
                  </div>
                </div>
              </div>
              {{ end }}
            </div>
            {{ end }}
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
