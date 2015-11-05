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
        {{ $definitions := .subject.GetDefinitions }}
        {{ $kas := .subject.GetKAs  }}
        <div class="panel-body">
          <h1>{{ .subject.GetName }}</h1>
          <br />
          <ul class="nav nav-tabs">
            <li class="active">
              <a href="#overview" data-toggle="tab">Overview</a>
            </li>

            {{ if ne ( $definitions | len ) 0 }}
            <li>
              <a href="#definitions" data-toggle="tab">Definitions</a>
            </li>
            {{ end }}

            {{ if ne ( $kas | len ) 0 }}
            <li>
              <a href="#ka" data-toggle="tab">Knowledge areas</a>
            </li>
            {{ end }}

            <li>
              <a href="#courses" data-toggle="tab">Courses</a>
            </li>
          </ul>

          <div id="myTabContent" class="tab-content">
            <div class="tab-pane fade active in" id="overview">
              {{ .subject.GetHTMLDescription }}
            </div>

            <div class="tab-pane fade in" id="definitions">
              <div>
                {{ range $index, $definition := $definitions }}
                <blockquote>
                  <p>{{ $definition.GetContent }}</p>
                  <small>Source <cite title="Source Title">{{ $definition.GetSource }}</cite></small>
                </blockquote>
                {{ end }}
              </div>
            </div>

            <div class="tab-pane fade in" id="ka">
              {{ range $index, $ka := $kas }}
              <div class="panel panel-default">
                <div class="panel-body">
                  <br />
                  <div class="row">
                    <div class="col-md-12">
                      <img  style="margin:10px; float:left" alt="{{ $ka.GetTitle }}" class="thumbnail" src="/static/img/education/cs/ka/{{ $ka.GetCode }}.png"/>

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

            <div class="tab-pane fade in" id="courses">
              {{ $subject := .subject }}
              {{ range $index, $institution := .institutions }}
              <div>
                {{ $programs := $institution.GetProgramsBySubjectID $subject.GetID }}
                <h4>{{ $institution.Name }}</h4>
                {{ range $index2, $program := $programs }}
                <div>
                  <div class="well">
                      <h5><a href="/institutions/{{ $institution.ID}}/program/{{ $program.GetID }}">{{ $program.GetTitle }}</a></h5>
                      {{ $program.GetContent }}
                    </div>
                </div>
                {{ end }}
              </div>
              <hr />
            </div>
          </div>

          <!-- <span class="text-warning glyphicon glyphicon-warning-sign"></span> Wisply was not able to generate data about <strong>{{ .subject.GetName }}</strong>'s curriculum from the institutions. -->

        </div>
      </div>
    </div>
  </div>
</div>
