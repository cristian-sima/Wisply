
<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/repositories">Repositories</a></li>
      <li class="active">Harvest</li>
      <li class="active">Init</li>
    </ul>
  </div>
  <div class="panel-body">
    <div class="row">
        <div class="col-md-8">
            <div class="input-group" style="height:28px;" id="URL-input">
              <input type="input" id="Source-URL" disabled value="{{ .repository.URL }}"/><span class="input-group-btn"><input disabled class="btn btn-primary btn-sm" value="Modifty" id="modifyButton"></span>
            </span>
          </div>
          <div id="Name-Repository" style="display:none">
              <strong>{{ .repository.Name }}</strong>
          </div>
          </div>
        <div class="col-md-4">
          <div id="connectionStatus">Please wait...</div>
        </div>
    </div>
    <br />
    <div id="generalIndicator" class="progress progress-striped active">
      <div class="progress-bar" style="width: 0%"></div>
    </div>
    <div class="row" >
      <div class="col-lg-3 col-md-3 col-sm-3" >
        <div id="process-status"></div>
          <div id="stages" class="list-group">
          </div>
      </div>
      <div class="col-lg-7 col-md-7 col-sm-7" >
        <ul class="nav nav-tabs">
          <li class="active"  id="currentButton"><a href="#current" data-toggle="tab" aria-expanded="true">Current</a></li>
          <li class="" id="historyButton"><a href="#history" data-toggle="tab" aria-expanded="false">History <span id="historyBadge" class="badge">0</span></a>  </li>
        </ul>
        <div id="myTabContent" class="tab-content" >
          <div class="tab-pane fade active in text-center" id="current">
          </div>
          <div class="tab-pane fade" id="history">

          </div>
        </div>
      </div>

    </div>
  </div>
  <style>
    #Source-URL {
      width: 100%;
      border:none;
    }
  </style>
  <script src="/static/js/admin/repository/list.js"></script>
  <script>
  var data = {};
  data.id = {{ .repository.ID }}
  data.name = {{ .repository.Name}}
  data.host = {{ .host }};
  </script>
  <script src="/static/js/admin/harvest/init.js"></script>
