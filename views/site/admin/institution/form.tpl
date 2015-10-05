<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/institutions">Institutions</a></li>
      <li class="active">{{.action}}</li>
    </ul></div>
    <div class="panel-body">
        <form method="{{.actionType}}" class="form-horizontal" >
          {{ .xsrf_input }}
          {{ $safeDescription := .institutionDescription|html}}
          <fieldset>
            <div class="form-group">
              <label for="institution-name" class="col-lg-2 control-label">Name</label>
              <div class="col-lg-10">
                <input type="text" value="{{.institutionName}}" class="form-control" name="institution-name" id="institution-name" placeholder="Name" required pattern=".{3,255}" title="The name has 3 up to 255 characters!">
              </div>
            </div>
            {{ if eq .action "Add" }}
            <div class="form-group">
              <label for="institution-URL" class="col-lg-2 control-label">Base URL</label>
              <div class="col-lg-10">
                <input type="url" value="{{.institutionUrl}}" class="form-control" name="institution-URL" id="institution-URL" placeholder="URL address" required pattern=".{3,2083}" title="The URL has 3 up to 2083 characters!">
              </div>
            </div>
            {{ end }}
            <div class="form-group text-center" style="min-height:80px">

              <label for="institution-description" class="col-lg-2 control-label text-center" id="institution-logo" ><span  class="institution-logo glyphicon glyphicon-education institution-logo"></span></label>

              <div class="col-lg-10">
                <textarea class="form-control" rows="3" name="institution-description" id="institution-description" maxlength="1000" >{{ .institutionDescription}}</textarea>
                <span class="help-block"><span class="hideMe" id="description-modified"><span class="text-warning"> <span class="glyphicon glyphicon-warning-sign"></span> Description modified</span> <span class="text-danger"><span id="discard-description-changes" data-toggle="tooltip" data-placement="top" title="Discard changes" class="hover glyphicon glyphicon-remove"></span></span></span><br /> This field may contain notes about the intitution. <a id='show-wiki-source' href="#">Modify source</a></span>
              </div>
            </div>
            <div class="hideMe" id="wiki-source-div">
              <div class="form-group">
                <label for="institution-description" class="col-lg-2 control-label">Wiki URL</label>
                <div class="col-lg-10">
                <input type="url" value="" class="form-control" name="institution-wikiURL" id="institution-wikiURL" placeholder="Wiki address" required pattern=".{3,2083}" title="The wiki URL has 3 up to 2083 characters!">
                </div>
              </div>
            </div>
            <div class="form-group">
              <div class="col-lg-10 col-lg-offset-2">
                <input type="submit" class="btn btn-primary" value="Submit" /> <a href="/admin/institutions" class="btn btn-default">Cancel</a>
              </div>
            </div>
          </fieldset>
        </form>
    </div>
  </div>
  <style>
  .institution-logo {
    font-size: 5em;
  }
  </style>
  <script src="/static/3rd_party/others/js/jquery.elastic.source.js"></script>
  <script src="/static/js/wisply/typer.js"></script>
  <script src="/static/js/wisply/wikier.js"></script>
  <script src="/static/js/admin/institution/functionality.js"></script>
