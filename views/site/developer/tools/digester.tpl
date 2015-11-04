<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/developer">API &amp; Developers</a></li>
            <li><a href="/developer/#tools">Tools</a></li>
            <li class="active">Digester</li>
          </ul>
        </div>
        <article class="panel-body">
          <div class="well h5">
            <span class="text-muted glyphicon glyphicon glyphicon-tasks"></span> Digester is a simple tool which takes a text and produces a list of occurences for its words.
            <br />
            <br />
          </div>
          <form method="POST" action="#results" class="form-horizontal" id="digester">
            {{ .xsrf_input }}
            <fieldset>
              <div class="form-group">
                <label for="digester-text" class="col-lg-2 control-label">Text</label>
                <div class="col-lg-10 well">
                  <textarea rows="10" class="form-control" name="digester-text" id="digester-text" placeholder="Please paste your text here..." title="Up to ">{{ .originalText }}</textarea>
                  <br />
                  The text must not exceed 50000 characters (otherwise it is trimmed) <br />
                  The tool will remove the articles, pronouns, conjuctions, and prepositions. Also, some redundant words (such as "yes", "is") will be removed
                </div>
              </div>
              <div class="form-group">
                <div class="col-lg-10 col-lg-offset-2">
                  <input type="submit" class="runButton btn btn-primary" value="Run" />
                </div>
              </div>
            </fieldset>
          </form>
          {{ if .processed }}
          {{ if eq (.processed.GetData | len ) 0 }}
          No data provided
          {{ else }}
          <a class="scroll" href="#digester">Original Text</a>
          <h3 id="results">Results:</h3>
          <div class="well">
            <ul class="nav nav-tabs">
              <li class="active"><a href="#visual" data-toggle="tab">Visual</a></li>
              <li><a href="#json" data-toggle="tab">Raw data(JSON)</a></li>
            </ul>
            <div id="myTabContent" class="tab-content">
              <div class="tab-pane fade active in" id="visual">
                <h4>Word occurence</h4>
                <p>
                  {{range $index, $occurence := .processed.GetData }}
                  <span data-word="{{ $occurence.GetWord }}" data-count="{{ $occurence.GetCounter }}" class="occurence label label-info">{{ $occurence.GetWord }}&nbsp; <span class="badge"> {{ $occurence.GetCounter }}</span></span>
                  {{ end }}
                </p>
              </div>
              <div class="tab-pane fade" id="json">
                <p>
                  <textarea style="height:500px" class="big-textarea form-control" placeholder="Result">{{ .processed.GetJSON }}
                  </textarea>
                </p>
              </div>
            </div>
          </div>
          {{ end }}
          {{ end }}
        </article>
      </div>
    </div>
  </div>
</div> <!--
  <div>
  <style scoped>
  .big-textarea {
  height: 500px;
}
</style>
</div> -->
<script>
$(document).ready(function(){
  $(".runButton").click(function(){
    wisply.message.tellToWait("Running...");
  })
});
</script>
