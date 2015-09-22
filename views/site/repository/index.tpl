
<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li class="active">Repositories</li>
    </ul>
  </div>
  <div class="panel-body">
    <div class="row">
        <div class="col-md-8">
            <div class="input-group">
              <input type="input" id="Source-URL" disabled value="http://www.edshare.soton.ac.uk/cgi/oai2"/>  <span class="input-group-btn"><input disabled class="btn btn-primary btn-sm" value="Modifty" id="modifyButton" />
            </span>
          </div>
          </div>
        <div class="col-md-4"><div id="connectionStatus">Please wait...</div></div>
    </div>
    <br />
    <div id="generalIndicator" class="progress progress-striped active">
      <div class="progress-bar" style="width: 0%"></div>
    </div>
    <div class="row" >
      <div class="col-lg-3 col-md-3 col-sm-3" >
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
border: 1px solid gainsboro;
    }
  </style>
  <script>
  var host = {{ .host }};
  </script>
  <script src="/static/js/admin/repository/index.js"></script>
