<div class="page-header">
  <div class="row">
    <div class="col-lg-12 col-md-12 col-sm-12">
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/education">Programs of study</a></li>
            <li class="active">{{ .program.GetName }}</li>
          </ul>
        </div>
        <div class="panel-body">
          <h1>{{ .program.GetName }}</h1>
          <br />
          <span class="text-warning glyphicon glyphicon-warning-sign"></span> Wisply was not able to generate data about <strong>{{ .program.GetName }}</strong>'s curriculum from the institutions.
          <br />
          <br />
          {{ .program.GetHTMLDescription }}
        </div>
      </div>
    </div>
  </div>
</div>
