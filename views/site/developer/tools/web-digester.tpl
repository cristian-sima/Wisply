<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/developer">API &amp; Developers</a></li>
            <li><a href="/developer/#tools">Tools</a></li>
            <li class="active">Web Digester</li>
          </ul>
        </div>
        <article class="panel-body">
          <div class="well h5">
            <span class="text-muted glyphicon glyphicon glyphicon-tasks"></span> The web digester finds the occurrences of the words inside the text. Then, it finds all the links and gets the content of these links. The content is again digested.
            <br />
            <br />
          </div>
          <form method="POST" class="form-horizontal" id="digester">
            {{ .xsrf_input }}
            <fieldset>
              <div class="form-group">
                <label for="digester-text" class="col-lg-2 control-label">Text</label>
                <div class="col-lg-10">
                  <textarea class="form-control" name="digester-text" id="digester-text" placeholder="Text" title="Up to ">{{ .originalText }}</textarea>
                  <br />
                  The text must not exceed 10000 characters (otherwise it is trimmed) <br />
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
          <h3>Results:</h3>
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
