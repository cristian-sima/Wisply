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

              <script>
              var
              dataPool = {};
              </script>


              <!-- Analyses -->
              <h5>Analyses:</h5>
              {{ $analyses := .subjectAnalyses }}
              {{ if eq ($analyses | len) 0 }}
              <div class="text-muted">There are no analyses for this subject.</div>
              {{ else }}
              <div class="list-group">
                {{ $analyses := $analyses }}
                {{ $modules := .modules }}
                {{ range $index, $analyse := $analyses }}
                {{ $parent := $analyse.GetParent }}
                {{ $general := $analyse.GetGeneral }}
                {{ $description := $analyse.GetDescriptionDigest }}
                {{ $keywords := $analyse.GetKeywordsDigest }}
                {{ $formats := $analyse.GetFormatsDigest }}
                <div class="panel panel-default">
                  <div class="panel-heading">{{ $parent.GetStartDate }}</div>
                  <div class="panel-body">
                    <br />
                    <a href="/about#how" class="btn btn-xs btn-default">How is it working?</a>


                    <br />
                    <!-- Overview -->
                    <div class="well">
                      <ul class="nav nav-tabs">
                        <li class="active">
                          <a href="#overall-overview-{{ $parent.GetID }}" data-toggle="tab">
                            Overview
                          </a>
                        </li>
                        <li>
                          <a href="#description-words-{{ $parent.GetID }}" data-toggle="tab">
                            Words
                          </a>
                        </li>
                        <li>
                          <a href="#description-json-{{ $parent.GetID }}" data-toggle="tab">
                            Data(JSON)
                          </a>
                        </li>
                      </ul>
                      <div id="tab-{{ $parent.GetID }}" class="tab-content">
                        <div class="tab-pane fade active in" id="overall-overview-{{ $parent.GetID }}">
                          <div class="row text-center">
                            <div class="col-md-6 text-center">
                              <div class="container-canvas text-center">
                                <div class="panel panel-default">
                                  <div class="panel-body">
                                    <h5>General aspects</h5>
                                    <hr />
                                    <br />
                                    <div style="height:360px;weight:300px">
                                      <div class="row">
                                        <div class="col-md-offset-4 col-md-4">
                                          <span class="big-big label-as-a-badge text-info">
                                            {{ $analyse.GetSubject.GetName }}
                                          </span>
                                          <hr />
                                          <span class="big-big label-as-a-badge">
                                            {{ $general.GetData | len }} words
                                          </span>
                                          <hr />
                                          <span class="big-big label-as-a-badge">
                                            {{ $general.GetCounter }} occurences
                                          </span>
                                          <hr />
                                        </div>
                                      </div>
                                    </div>
                                  </div>
                                  <hr />
                                </div>
                              </div>
                            </div>
                            <div class="col-md-6">
                              <div class="panel panel-default">
                                <div class="panel-body">
                                  <h5>Prominent words</h5>
                                  <hr />
                                  <br />
                                  <div class="container-canvas">
                                    <!-- overall - Most prominent -->
                                    {{ $proeminent := $general.GetMostProeminent 8 }}
                                    {{ if eq ($proeminent.GetData | len ) 0 }}
                                    <div class="text-muted" style="height:300px;weight:300px">
                                      <br /><br /><br /> There are no prominent words <span class="glyphicon glyphicon-info-sign" data-toggle="tooltip" title="This happens because the words have a constant distribution"></span>
                                    </div>
                                    {{ else }}
                                    <canvas style="margin:0 auto" width="300px" height="300px" data-id="chart-list-overall-2-{{ $parent.GetID }}" class="chart img-responsive chart chart-radar" id="chart-list-overall-2-{{ $parent.GetID }}" >
                                    </canvas>
                                    <script>
                                    dataPool["chart-list-overall-2-{{ $parent.GetID }}"] = JSON.parse({{$proeminent.GetPlainJSON }});
                                    </script>
                                    {{ end }}
                                  </div>
                                  <br /> A word is prominent if it occures for at least:<br />
                                  <img src="/static/img/public/equation/proeminent-8.png" alt="Prominent equation" style="margin:0 auto" class="img-responsive" />
                                </div>
                              </div>
                              <hr />
                            </div>
                          </div>
                          <hr />
                          <div class="row text-center">
                            <div class="col-md-6 text-center">
                              <div class="container-canvas text-center">
                                <div class="panel panel-default">
                                  <div class="panel-body">
                                    <h5>Most relevant</h5>
                                    <hr />
                                    <br />
                                    <!-- overall - Most relevant -->
                                    <canvas style="margin:0 auto" width="300px" height="300px" data-id="chart-list-overall-3-{{ $parent.GetID }}" class="text-center chart img-responsive chart chart-doughnut text-center" id="chart-list-overall-3-{{ $parent.GetID }}" >
                                    </canvas>
                                    <script>
                                    dataPool["chart-list-overall-3-{{ $parent.GetID }}"] = JSON.parse({{$general.GetMostRelevant.GetPlainJSON }});
                                    </script>
                                  </div>
                                  <br /> A word is relevant if it appears at least for:<br />
                                  <img src="/static/img/public/equation/relevant.png" alt="Relevant equation" style="margin:0 auto" class="img-responsive" />
                                  <hr />
                                </div>
                              </div>
                            </div>
                            <div class="col-md-6">
                              <div class="panel panel-default">
                                <div class="panel-body">
                                  <h5>Top 10</h5>
                                  <hr />
                                  <br />
                                  <div class="container-canvas">
                                    <!-- Overall - Top 10 -->
                                    <canvas style="margin:0 auto" width="300px" height="300px" data-id="chart-list-overall-4-{{ $parent.GetID }}" class="chart img-responsive chart chart-radar" id="chart-list-overall-4-{{ $parent.GetID }}" >
                                    </canvas>
                                    <script>
                                    {{ $top := $general.GetTop 10 }}
                                    dataPool["chart-list-overall-4-{{ $parent.GetID }}"] = JSON.parse({{$top.GetPlainJSON }});
                                    </script>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                          <hr />
                          <h4>How is it taught</h4>
                          {{ $formats := $analyse.GetFormatsDigest }}
                          {{ if eq ($formats.GetData | len) 0 }}
                          No information available
                          {{ else }}
                          <div class="panel panel-default">
                            <div class="panel-body">
                              <div class="container-canvas">
                                <canvas style="margin:0 auto" width="300px" height="300px" data-id="chart-{{ $parent.GetID }}" class="chart img-responsive chart chart-pie" id="chart-{{ $parent.GetID }}" >
                                </canvas>
                                <script>
                                dataPool["chart-{{ $parent.GetID }}"] = JSON.parse({{$formats.GetPlainJSON }});
                                </script>
                              </div>
                            </div>
                          </div>

                          <!--        Comparison     --->
                          <h4>Specified curriculum <span class="text-warning">vs</span> what it is actually taught by teachers</h4>
                          <div class="row text-center">
                            <div class="col-md-6 text-center">
                              <div class="container-canvas text-center">
                                <div class="panel panel-default">
                                  <div class="panel-body">
                                    <h5>Most relevant - Specified</h5>
                                    <hr />
                                    <br />
                                    <!-- Specified - Most relevant -->
                                    <canvas style="margin:0 auto" width="300px" height="300px" data-id="chart-list-description-1-{{ $parent.GetID }}" class="text-center chart img-responsive chart chart-doughnut text-center" id="chart-list-description-1-{{ $parent.GetID }}" >
                                    </canvas>
                                    <script>
                                    dataPool["chart-list-description-1-{{ $parent.GetID }}"] = JSON.parse({{$description.GetMostRelevant.GetPlainJSON }});
                                    </script>
                                  </div>
                                  <br /> A word is relevant if it appears at least for:<br />
                                  <img src="/static/img/public/equation/relevant.png" alt="Relevant equation" style="margin:0 auto" class="img-responsive" />
                                  <hr />
                                </div>
                              </div>
                            </div>
                            <div class="col-md-6">
                              <div class="panel panel-default">
                                <div class="panel-body">
                                  <h5>Prominent words - Specified</h5>
                                  <hr />
                                  <br />
                                  <div class="container-canvas">
                                    <!-- Specified - Most proeminent -->
                                    {{ $proeminentDescription := $description.GetMostProeminent 8 }}
                                    {{ if eq ($proeminentDescription.GetData | len ) 0 }}
                                    <div class="text-muted" style="height:300px;weight:300px">
                                      <br /><br /><br /> There are no prominent words <span class="glyphicon glyphicon-info-sign" data-toggle="tooltip" title="This happens because the words have a constant distribution"></span>
                                    </div>
                                    {{ else }}
                                    <canvas style="margin:0 auto" width="300px" height="300px" data-id="chart-list-description-2-{{ $parent.GetID }}" class="chart img-responsive chart chart-radar" id="chart-list-description-2-{{ $parent.GetID }}" >
                                    </canvas>
                                    <script>
                                    dataPool["chart-list-description-2-{{ $parent.GetID }}"] = JSON.parse({{$proeminentDescription.GetPlainJSON }});
                                    </script>
                                    {{ end }}
                                  </div>
                                  <br /> A word is prominent if it appears at least for:
                                  <img src="/static/img/public/equation/proeminent-8.png" alt="Prominent equation" style="margin:0 auto" class="img-responsive" />
                                </div>
                              </div>
                              <hr />
                            </div>
                          </div>
                          <!-- keywords -->
                          <div class="row text-center">
                            <div class="col-md-6 text-center">
                              <div class="container-canvas text-center">
                                <div class="panel panel-default">
                                  <div class="panel-body">
                                    <h5>Most relevant - Taught</h5>
                                    <hr />
                                    <br />
                                    <!-- Taught - Most relevant -->
                                    <canvas style="margin:0 auto" width="300px" height="300px" data-id="chart-list-taught-1-{{ $parent.GetID }}" class="text-center chart img-responsive chart chart-doughnut text-center" id="chart-list-taught-1-{{ $parent.GetID }}" >
                                    </canvas>
                                    <script>
                                    dataPool["chart-list-taught-1-{{ $parent.GetID }}"] = JSON.parse({{$keywords.GetMostRelevant.GetPlainJSON }});
                                    </script>
                                  </div>
                                  <br /> A word is relevant if it appears at least for:<br />
                                  <img src="/static/img/public/equation/relevant.png" alt="Relevant equation" style="margin:0 auto" class="img-responsive" />
                                  <hr />
                                </div>
                              </div>
                            </div>
                            <div class="col-md-6">
                              <div class="panel panel-default">
                                <div class="panel-body">
                                  <h5>Prominent words - Taught</h5>
                                  <hr />
                                  <br />
                                  <div class="container-canvas">
                                    <!-- taught - Most proeminent -->
                                    {{ $proeminentT := $keywords.GetMostProeminent 8 }}
                                    {{ if eq ($proeminentT.GetData | len ) 0 }}
                                    <div class="text-muted" style="height:300px;weight:300px">
                                      <br /><br /><br /> There are no prominent words <span class="glyphicon glyphicon-info-sign" data-toggle="tooltip" title="This happens because the words have a constant distribution"></span>
                                    </div>
                                    {{ else }}
                                    <canvas style="margin:0 auto" width="300px" height="300px" data-id="chart-list-taught-2-{{ $parent.GetID }}" class="chart img-responsive chart chart-radar" id="chart-list-taught-2-{{ $parent.GetID }}" >
                                    </canvas>
                                    <script>
                                    dataPool["chart-list-taught-2-{{ $parent.GetID }}"] = JSON.parse({{$proeminentT.GetPlainJSON }});
                                    </script>
                                    {{ end }}
                                  </div>
                                  <br /> A word is prominent if it occures for at least:<br />
                                  <img src="/static/img/public/equation/proeminent-8.png" alt="Prominent equation" style="margin:0 auto" class="img-responsive" />
                                </div>
                              </div>
                              <hr />
                            </div>
                          </div>
                          {{ end }}
                        </div>
                        <div  class="tab-pane fade in" id="description-words-{{ $parent.GetID }}">
                          <a href="#" id="showColors" class="btn btn-xs btn-default">
                            <span class="glyphicon glyphicon-text-background text-primary"></span> Remove colors
                          </a> &nbsp;

                          <!-- List Words - Top 25 - Specified -->
                          <h4>Specified curriculum (top 25)</h4>
                          {{ if eq ($keywords.GetData | len) 0 }}
                          No information available
                          {{ else }}
                          <p>
                            {{ $topDes := $description.GetTop 25 }}
                            {{range $index, $occurence := $topDes.GetData }}
                            <span data-group="keywords" data-word="{{ $occurence.GetWord }}" data-count="{{ $occurence.GetCounter }}" class="word-occurence label label-info">
                              {{ $occurence.GetWord }}
                              &nbsp;
                              <span class="badge">
                                {{ $occurence.GetCounter }}
                              </span>
                            </span>&nbsp;
                            {{ end }}
                          </p>
                          <hr />
                          {{ end }}

                          <!-- List Words - Top 25 - Taught -->
                          <h4>Keywords used for teaching (top 25)</h4>
                          {{ if eq ($keywords.GetData | len) 0 }}
                          No information available
                          {{ else }}
                          <p>
                            {{ $topKey := $keywords.GetTop 25 }}
                            {{range $index, $occurence := $topKey.GetData }}
                            <span data-group="specified" data-word="{{ $occurence.GetWord }}" data-count="{{ $occurence.GetCounter }}" class="word-occurence label label-info">
                              {{ $occurence.GetWord }}
                              &nbsp;
                              <span class="badge">
                                {{ $occurence.GetCounter }}
                              </span>
                            </span>&nbsp;
                            {{ end }}
                          </p>
                          <hr />
                          {{ end }}

                          <!-- List Words - All - Specified -->
                          <h4>Specified (All)</h4>
                          <p>
                            {{range $index, $occurence := $description.GetData }}
                            <span data-group="specified" data-word="{{ $occurence.GetWord }}" data-count="{{ $occurence.GetCounter }}" class="word-occurence label label-info">
                              {{ $occurence.GetWord }}
                              &nbsp;
                              <span class="badge">
                                {{ $occurence.GetCounter }}
                              </span>
                            </span>&nbsp;
                            {{ end }}
                          </p>
                          <hr />

                          <!-- List Words - All - Formats -->
                          <h4>Formats</h4>
                          {{ if eq ($formats.GetData | len) 0 }}
                          No information available
                          {{ else }}
                          <p>
                            {{range $index, $occurence := $formats.GetData }}
                            <span data-group="formats" data-word="{{ $occurence.GetWord }}" data-count="{{ $occurence.GetCounter }}" class="word-occurence label label-info">
                              {{ $occurence.GetWord }}
                              &nbsp;
                              <span class="badge">
                                {{ $occurence.GetCounter }}
                              </span>
                            </span>&nbsp;
                            {{ end }}
                            <hr />
                          </p>
                          {{ end }}

                          <!-- List Words - All- Keywords for teaching -->
                          <h4>Keywords used for teaching (All)</h4>
                          {{ if eq ($keywords.GetData | len) 0 }}
                          No information available
                          {{ else }}
                          <p>
                            {{range $index, $occurence := $keywords.GetData }}
                            <span data-group="keywords" data-word="{{ $occurence.GetWord }}" data-count="{{ $occurence.GetCounter }}" class="word-occurence label label-info">
                              {{ $occurence.GetWord }}
                              &nbsp;
                              <span class="badge">
                                {{ $occurence.GetCounter }}
                              </span>
                            </span>&nbsp;
                            {{ end }}
                          </p>
                          {{ end }}

                        </div>
                        <div class="tab-pane fade" id="description-json-{{ $parent.GetID }}">
                          <p>
                            <textarea style="background:white;height:500px" class="big-textarea form-control" placeholder="Result">{{ $general.GetJSON }}
                            </textarea>
                          </p>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                {{ end }}
              </div>
              {{ end }}
              <style scoped>
              .canvas {
                margin: 0 auto;
              }
              .big-big {
                font-size: 22px;
              }
              </style>
            </div>
            <script src="/static/3rd_party/others/js/Chart.min.js"></script>

            <script src="/static/js/wisply/chart.js"></script>
            <script>
            $(document).ready(function(){
              var module = wisply.getModule("chart");
              module.init();
            });
            </script>




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
                <br />
                {{ $programs := $institution.GetProgramsBySubjectID $subject.GetID }}
                {{ if ne ($programs | len ) 0 }}
                <h4><a href="/institutions/{{ $institution.ID }}">{{ $institution.Name }}</a></h4>
                {{ range $index2, $program := $programs }}
                <div>
                  <div class="well">
                    <h5><a href="/institutions/{{ $institution.ID}}/program/{{ $program.GetID }}">{{ $program.GetTitle }}</a></h5>
                    <span class="text-muted">Program of study &bull; <span class="capitalize">{{ $program.GetLevel }}</span></span>
                    <br />
                    <br />
                    {{ $program.GetContent }}
                  </div>
                </div>
                {{ end }}
                {{ end }}
              </div>
              {{ end }}
              <hr />
            </div>
          </div>

          <!-- <span class="text-warning glyphicon glyphicon-warning-sign"></span> Wisply was not able to generate data about <strong>{{ .subject.GetName }}</strong>'s curriculum from the institutions. -->

        </div>
      </div>
    </div>
  </div>
</div>
<div>
<style scoped>
.capitalize {
  text-transform: capitalize;
}
</style>
</div>
