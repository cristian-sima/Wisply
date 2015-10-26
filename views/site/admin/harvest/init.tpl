<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/repositories">Repositories</a></li>
      <li class="active">{{ .repository.Name }}</li>
    </ul>
  </div>
  <div class="panel-body">
    <div class="row">
      <div class="col-md-8">
        <div class="input-group" style="height:28px;" id="URL-input">
          <input type="input" id="Source-URL" placeholder="http://..." disabled value="{{ .repository.URL }}"/><span class="input-group-btn"><input disabled class="btn btn-primary btn-sm" value="Modifty" id="modifyButton"></span>
        </span>
      </div>
      <div id="Name-Repository" style="display:none">
        <strong>{{ .repository.Name }}</strong>
      </div>
    </div>
    <div class="col-md-4 text-right">
      <div id="websocket-connection"></div>
    </div>
  </div>
  <br />
  <div id="general-indicator" class="progress progress-striped active">
    <div class="progress-bar" style="width: 0%"></div>
  </div>
  <div class="row" >
    <div class="col-lg-3 col-md-3 col-sm-3" id="stages" style="display:none">
      <div id="repository-status"></div>
      <div id="process-status"></div>
      <div id="stage-list" class="list-group">
      </div>
    </div>
    <div class="col-lg-9 col-md-9 col-sm-9" >
      <ul class="nav nav-tabs">
        <li class="active"  id="currentButton"><a href="#current" data-toggle="tab" aria-expanded="true">Current</a></li>
        <li class="" id="historyButton"><a href="#history" data-toggle="tab" aria-expanded="false">History</a>  </li>
      </ul>
      <div id="myTabContent" class="tab-content" >
        <div class="tab-pane fade active in text-center" id="current">
        </div>
        <div class="tab-pane fade" id="history"></div>
      </div>
    </div>
  </div>
</div>
<div>
  <style scoped>
  #Source-URL {
    width: 100%;
    border:none;
  }
  .big-number {
    font-size: 22px;
    color:#dfdfdf;
  }
  .counter-name {
    font-size: 16px;

  }
  </style>
</div>
<script>
var server = {},
repository;
server.repository = {};
repository = server.repository;

repository.id =   {{ .repository.ID }}
repository.name = {{ .repository.Name}}
repository.status = {{ .repository.Status}}
repository.url =  {{ .repositoru.URL }}

server.host = {{ .host }};

server.hasProcess = {{ .hasProcess }};
server.process = {{ .currentProcesses }};

</script>
<script src="/static/js/ws/websockets.js"></script>
<script src="/static/js/admin/repository/list.js"></script>
<script src="/static/js/admin/harvest/harvest.js"></script>
<script src="/static/js/admin/harvest/process.js"></script>
<script src="/static/3rd_party/others/js/countUp.min.js"></script>
