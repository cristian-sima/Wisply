<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/institutions">Institutions</a></li>
            <li><a href="/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></li>
            <li class="active">{{ .module.GetTitle }}</li>
          </ul>
        </div>
        <script>
        var
        dataPool = {};
        </script>
        <div class="panel-body">
          <div style="margin:0px">
            <div class="row">
              <div class="col-md-9">
                <h1>{{ .module.GetTitle }}</h1>
                <span class="text-muted">Module</span> &bull; <a href="/institutions/{{ .institution.ID }}">{{ .institution.Name}}</a>
                <div class="well">
                  {{ .module.GetContent }}
                </div>
              </div>
              <div class="col-md-3">
                <h2>Information</h2>
                <table class="table">
                  <tbody>
                    <tr>
                      <td>Code</td>
                      <td>{{ .module.GetCode }}</td>
                    </tr>
                    <tr>
                      <td>Year</td>
                      <td>{{ .module.GetYear }}</td>
                    </tr>
                    <tr>
                      <td>CATS</td>
                      <td>{{ .module.GetCredits "CATS" }}</td>
                    </tr>
                    <tr>
                      <td>ECTS</td>
                      <td>{{ .module.GetCredits "ECTS" }}</td>
                    </tr>
                    <tr>
                      <td>US credits</td>
                      <td>{{ .module.GetCredits "US" }}</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
          <br />
          <br />
          <!-- Analyses -->
          <h5>Analyses:</h5>
          {{ $analyses := .moduleAnalyses }}
          {{ if eq ($analyses | len) 0 }}
          <div class="text-muted">There are no analyses for this module.</div>
          {{ else }}
          <div class="list-group">
            {{ $analyses := $analyses }}
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
                                <h5>Words distribution</h5>
                                <hr />
                                <br />
                                <!-- overall - All words -->
                                <canvas style="margin:0 auto" width="300px" height="300px" data-id="chart-list-overall-1-{{ $parent.GetID }}" class="text-center chart img-responsive chart chart-doughnut text-center" id="chart-list-overall-1-{{ $parent.GetID }}" >
                                </canvas>
                                <script>
                                dataPool["chart-list-overall-1-{{ $parent.GetID }}"] = JSON.parse({{$general.GetPlainJSON }});
                                </script>
                              </div>
                              This chart contains all the words which appear
                              <hr />
                            </div>
                          </div>
                        </div>
                        <div class="col-md-6">
                          <div class="panel panel-default">
                            <div class="panel-body">
                              <h5>Most prominent words</h5>
                              <hr />
                              <br />
                              <div class="container-canvas">
                                <!-- overall - Most proeminent -->
                                {{ $proeminent := $general.GetMostProeminent 3 }}
                                {{ if eq ($proeminent.GetData | len ) 0 }}
                                <div class="text-muted" style="height:300px;weight:300px">
                                  <br /><br /><br /> There are no prominent words <span class="glyphicon glyphicon-info-sign" data-toggle="tooltip" title="This happens because the words have a normal distribution"></span>
                                </div>
                                {{ else }}
                                <canvas style="margin:0 auto" width="300px" height="300px" data-id="chart-list-overall-2-{{ $parent.GetID }}" class="chart img-responsive chart chart-radar" id="chart-list-overall-2-{{ $parent.GetID }}" >
                                </canvas>
                                <script>
                                dataPool["chart-list-overall-2-{{ $parent.GetID }}"] = JSON.parse({{$proeminent.GetPlainJSON }});
                                </script>
                                {{ end }}
                              </div>
                              <br /> A word is prominent if it appears at least for:<br />
                              <img src="/static/img/public/equation/proeminent-3.png" alt="Prominent equation" style="margin:0 auto" class="img-responsive" />
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
                              <h5>Most prominent words - Specified</h5>
                              <hr />
                              <br />
                              <div class="container-canvas">
                                <!-- Specified - Most proeminent -->
                                {{ $proeminentDescription := $description.GetMostProeminent 3 }}
                                {{ if eq ($proeminentDescription.GetData | len ) 0 }}
                                <div class="text-muted" style="height:300px;weight:300px">
                                  <br /><br /><br /> There are no prominent words <span class="glyphicon glyphicon-info-sign" data-toggle="tooltip" title="This happens because the words have a normal distribution"></span>
                                </div>
                                {{ else }}
                                <canvas style="margin:0 auto" width="300px" height="300px" data-id="chart-list-description-2-{{ $parent.GetID }}" class="chart img-responsive chart chart-radar" id="chart-list-description-2-{{ $parent.GetID }}" >
                                </canvas>
                                <script>
                                dataPool["chart-list-description-2-{{ $parent.GetID }}"] = JSON.parse({{$proeminentDescription.GetPlainJSON }});
                                </script>
                                {{ end }}
                              </div>
                              <br /> A word is prominent if it appears at least for:<br />
                              <img src="/static/img/public/equation/proeminent-3.png" alt="Prominent equation" style="margin:0 auto" class="img-responsive" />
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
                              <h5>Most prominent words - Taught</h5>
                              <hr />
                              <br />
                              <div class="container-canvas">
                                <!-- taught - Most proeminent -->
                                {{ $proeminentKeywords := $keywords.GetMostProeminent 3 }}
                                {{ if eq ($proeminentKeywords.GetData | len ) 0 }}
                                <div class="text-muted" style="height:300px;weight:300px">
                                  <br /><br /><br /> There are no prominent words <span class="glyphicon glyphicon-info-sign" data-toggle="tooltip" title="This happens because the words have a normal distribution"></span>
                                </div>
                                {{ else }}
                                <canvas style="margin:0 auto" width="300px" height="300px" data-id="chart-list-taught-2-{{ $parent.GetID }}" class="chart img-responsive chart chart-radar" id="chart-list-taught-2-{{ $parent.GetID }}" >
                                </canvas>
                                <script>
                                dataPool["chart-list-taught-2-{{ $parent.GetID }}"] = JSON.parse({{$proeminentKeywords.GetPlainJSON }});
                                </script>
                                {{ end }}
                              </div>
                              <br /> A word is prominent if it appears at least for:<br />
                              <img src="/static/img/public/equation/proeminent-3.png" alt="Prominent equation" style="margin:0 auto" class="img-responsive" />
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

                      <h4>Specified</h4>
                      <p>
                        {{range $index, $occurence := $description.GetData }}
                        <span data-word="{{ $occurence.GetWord }}" data-count="{{ $occurence.GetCounter }}" class="word-occurence label label-info">
                          {{ $occurence.GetWord }}
                          &nbsp;
                          <span class="badge">
                            {{ $occurence.GetCounter }}
                          </span>
                        </span>&nbsp;
                        {{ end }}
                      </p>
                      <h4>Formats</h4>
                      {{ if eq ($formats.GetData | len) 0 }}
                      No information available
                      {{ else }}
                      <p>
                        {{range $index, $occurence := $formats.GetData }}
                        <span data-word="{{ $occurence.GetWord }}" data-count="{{ $occurence.GetCounter }}" class="word-occurence label label-info">
                          {{ $occurence.GetWord }}
                          &nbsp;
                          <span class="badge">
                            {{ $occurence.GetCounter }}
                          </span>
                        </span>&nbsp;
                        {{ end }}
                      </p>
                      {{ end }}
                      <h4>Keywords used for teaching</h4>
                      {{ if eq ($keywords.GetData | len) 0 }}
                      No information available
                      {{ else }}
                      <p>
                        {{range $index, $occurence := $keywords.GetData }}
                        <span data-word="{{ $occurence.GetWord }}" data-count="{{ $occurence.GetCounter }}" class="word-occurence label label-info">
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
          <!-- Programs -->
          <h5>Programs of study which include this module:</h5>
          {{ $programs := .module.GetPrograms }}
          {{ if eq ($programs | len) 0 }}
          <div class="text-muted">There are no programs of study which include this module.</div>
          {{ else }}
          <div class="list-group">
            {{ $institution := .institution }}
            {{ range $index, $program := $programs }}
            <a href="/institutions/{{ $institution.ID }}/program/{{ $program.GetID }}" class="list-group-item">
              {{ $program.GetCode }} - {{ $program.GetTitle }}
            </a>
            {{ end }}
          </div>
          {{ end }}
          <h5>Resources identified:</h5>
          <!-- Similar resources -->
          {{ if not .resourcesSuggested }}
          <span class="text-muted">There are no resources identified for this module</span>
          {{ else }}
            {{ range $index, $resource := .resourcesSuggested }}
              <a class="resource" href="{{ $resource.GetWisplyURL }}">
                  <h6>
                    {{ if not $resource.IsVisible }}
                    <small><span data-toggle='tooltip' title='This content is not visible to Wisply.' class='glyphicon glyphicon-lock'></span></small>
                    {{ end }}
                {{range $index, $title := $resource.Keys.Get "title" }}
                {{ $title }}
                {{ end }}
              </h6>
              </a>
            {{ end }}
            {{ end }}
        </div>
      </div>
    </div>
  </div>
</div>
<div>
<style scoped>
.canvas {
  margin: 0 auto;
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
