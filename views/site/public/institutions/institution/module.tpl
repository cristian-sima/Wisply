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
        dataPool = {},
        colors = {};
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
            <div class="panel panel-default">
              <div class="panel-heading">{{ $parent.GetStartDate }}</div>
              <div class="panel-body">
                <br />
                <a href="#" id="showColors" class="btn btn-xs btn-primary">Remove colors</a>
                <br />
                <!-- Description -->
                {{ $digester := $analyse.GetDescriptionDigest }}
                <h4>Description</h4>
                <div class="well">
                  <ul class="nav nav-tabs">
                    <li class="active">
                      <a href="#description-overview-{{ $parent.GetID }}" data-toggle="tab">
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
                    <div class="tab-pane fade active in" id="description-overview-{{ $parent.GetID }}">
                      <div class="row text-center">
                        <div class="col-md-6 text-center">
                          <div class="container-canvas text-center">

                            <div class="panel panel-default">
                              <div class="panel-body">
                                <h5>Words distribution</h5>
                                <hr />
                                <br />
                                <!-- Description - All words -->
                                <canvas style="margin:0 auto" width="300px" height="300px" data-id="chart-list-description-1-{{ $parent.GetID }}" class="text-center chart img-responsive chart chart-doughnut text-center" id="chart-list-description-1-{{ $parent.GetID }}" >
                                </canvas>
                                <script>
                                dataPool["chart-list-description-1-{{ $parent.GetID }}"] = JSON.parse({{$digester.GetPlainJSON }});
                                </script>
                              </div>
                              This chart contains all the words which appear in the description
                              <hr />
                            </div>
                          </div>
                        </div>
                        <div class="col-md-6">
                          <div class="panel panel-default">
                            <div class="panel-body">
                              <h5>Most proeminent words</h5>
                              <hr />
                              <br />
                              <div class="container-canvas">
                                <!-- Description - Most proeminent -->
                                <canvas style="margin:0 auto" width="300px" height="300px" data-id="chart-list-description-2-{{ $parent.GetID }}" class="chart img-responsive chart chart-radar" id="chart-list-description-2-{{ $parent.GetID }}" >
                                </canvas>
                                <script>
                                {{ $proeminent := $digester.GetMostProeminent }}
                                dataPool["chart-list-description-2-{{ $parent.GetID }}"] = JSON.parse({{$digester.GetMostProeminent.GetPlainJSON }});
                                </script>
                              </div>
                              <br /> A word is proeminent if it appear at least for:<br />
                              <math>(Total number of occurences)/(Distinct number of words)</math>
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
                                <!-- Description - Most relevant -->
                                <canvas style="margin:0 auto" width="300px" height="300px" data-id="chart-list-description-3-{{ $parent.GetID }}" class="text-center chart img-responsive chart chart-doughnut text-center" id="chart-list-description-3-{{ $parent.GetID }}" >
                                </canvas>
                                <script>
                                dataPool["chart-list-description-3-{{ $parent.GetID }}"] = JSON.parse({{$digester.GetMostRelevant.GetPlainJSON }});
                                </script>
                              </div>
                              <br /> A word is relevant if it appear at least for:<br />
                              <math>(Maximum number of occurences)/(2)</math>
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
                                <!-- Description - Top 10 -->
                                <canvas style="margin:0 auto" width="300px" height="300px" data-id="chart-list-description-4-{{ $parent.GetID }}" class="chart img-responsive chart chart-radar" id="chart-list-description-4-{{ $parent.GetID }}" >
                                </canvas>
                                <script>
                                {{ $top := $digester.GetTop 10 }}
                                dataPool["chart-list-description-4-{{ $parent.GetID }}"] = JSON.parse({{$digester.GetMostProeminent.GetPlainJSON }});
                                </script>
                              </div>

                            </div>
                          </div>
                          <hr />
                        </div>
                      </div>
                      <hr />
                    </div>
                    <div  class="tab-pane fade in" id="description-words-{{ $parent.GetID }}">
                      <h4>All words</h4>
                      <p>
                        {{range $index, $occurence := $digester.GetData }}
                        <span data-word="{{ $occurence.GetWord }}" data-count="{{ $occurence.GetCounter }}" class="word-occurence label label-info">
                          {{ $occurence.GetWord }}
                          &nbsp;
                          <span class="badge">
                            {{ $occurence.GetCounter }}
                          </span>
                        </span>&nbsp;
                        {{ end }}
                      </p>
                    </div>
                    <div class="tab-pane fade" id="description-json-{{ $parent.GetID }}">
                      <p>
                        <textarea style="background:white;height:500px" class="big-textarea form-control" placeholder="Result">{{ $digester.GetJSON }}
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
<script>
var showColors = true;
$(document).ready(function() {
  "use strict";
  $("#showColors").click(function(event){
    event.preventDefault();
    var element = $(this);
    if(!showColors) {
      colorWords();
      element.html("Remove colors");
    } else {
      $(".word-occurence").css({
        "background-color" : "",
        "color": "",
      });
      element.html("Show colors");
    }
    showColors = !showColors;
  });
  $(".chart-doughnut").each(function(){
    var element = $(this),
    id = element.data("id"),
    ctx = element[0].getContext("2d"),
    data = getData(id),
    options =  {
      //String - A legend template
      legendTemplate : "<ul class=\"<%=name.toLowerCase()%>-legend\"><% for (var i=0; i<segments.length; i++){%><li><span style=\"background-color:<%=segments[i].fillColor%>\"></span><%if(segments[i].label){%><%=segments[i].label%><%}%></li><%}%></ul>"
    };
    new Chart(ctx).Doughnut(data, options);
    colorWords();
  });
  $(".chart-radar").each(function(){
    var element = $(this),
    id = element.data("id"),
    ctx = element[0].getContext("2d"),
    data = getData(id),
    options =  {
      //String - A legend template
      legendTemplate : "<ul class=\"<%=name.toLowerCase()%>-legend\"><% for (var i=0; i<segments.length; i++){%><li><span style=\"background-color:<%=segments[i].fillColor%>\"></span><%if(segments[i].label){%><%=segments[i].label%><%}%></li><%}%></ul>"
    };
    new Chart(ctx).PolarArea(data, options);
    colorWords();
  });
  $(".chart-pie").each(function(){
    var element = $(this),
    id = element.data("id"),
    ctx = element[0].getContext("2d"),
    data = getData(id),
    options =  {
      //String - A legend template
      legendTemplate : "<ul class=\"<%=name.toLowerCase()%>-legend\"><% for (var i=0; i<segments.length; i++){%><li><span style=\"background-color:<%=segments[i].fillColor%>\"></span><%if(segments[i].label){%><%=segments[i].label%><%}%></li><%}%></ul>"
    };
    new Chart(ctx).Pie(data, options);
    colorWords();
  });
  function colorWords() {
    $(".word-occurence").each(function(){
      var element = $(this),
      word = element.data("word"),
      color = getColorForWord(word);
      element.css({"background-color": color.background,
      "color": color.font });
    });
  }
  /**
  * It gets the data for the chart. Also, it processes it
  * @param  {string} id The ID of the analyse
  * @return {object} The data
  */
  function getData(id) {
    var data = dataPool[id],
    i, occurence, newSet = [], newItem;
    // transform it
    console.log(id)
    for(i=0; i<data.length;i++) {
      occurence = data[i];
      newItem = {
        "label": occurence.Word,
        "value": occurence.Counter,
        color : getColorForWord(occurence.Word).background,
      };
      newSet.push(newItem);
    }
    return newSet;
  }
  /**
  * It checks to see if the color for the word is already stored. If not, it generates a new one
  * @param  {string} word The word
  * @return {object} The color in RGB format
  */
  function getColorForWord(word) {
    if(!colors[word]) {
      colors[word] = getRandomColor();
    }
    return colors[word];
  }
  /**
  * It returns a random color
  * @return {object} A random color for bg and font
  */
  function getRandomColor() {
    function getContrastYIQ(hexcolor){
      var r = parseInt(hexcolor.substr(0,2),16);
      var g = parseInt(hexcolor.substr(2,2),16);
      var b = parseInt(hexcolor.substr(4,2),16);
      var yiq = ((r*299)+(g*587)+(b*114))/1000;
      return (yiq >= 128) ? 'black' : 'white';
    }
    var obj = {},
    background = Math.floor(Math.random()*16777215).toString(16);
    // credits http://stackoverflow.com/questions/11070007/style-each-div-with-a-different-color-using-jquery-or-javascript
    obj.background = '#' + background;
    obj.font = getContrastYIQ(background);
    return obj;
  }
});
</script>
