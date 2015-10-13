
<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li class="active">Repositories</li>
    </ul>
  </div>
  <div class="panel-body">
    <section>
      <h4>Options</h4>
      <div class="btn-group-sm">
        <a href="/admin/repositories/add" class="btn btn-primary">Add repository</a>
        <span id="harvest-history-button" class="btn btn-info hover">Live actions</span>
      </div>
    </section>
    <br />
    <section>
      <div class="row">
        <div class="col-md-1">
          <h4>Repositories</h4>
        </div>
        <div class="col-md-11 text-right">
          <div id="websocket-connection"></div>
        </div>
      </div>
    {{ if ne (len .repositories) 0 }}
    <div class="table-responsive">
      <table class="table table-striped table-hover " id="repositories-list">
        <thead>
          <tr>
            <th>Name</th>
            <th>Current status</th>
            <th>Progress</th>
            <th>Base URL</th>
            <th>Institution</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $element := .repositories}}
          {{ $institution := $element.GetInstitution }}
          {{$safe := $element.Name|html}}
          <tr>
            <td><a href="/admin/repositories/repository/{{ $element.ID }}">{{ $element.Name |html }}</a></td>
            <td> <div  id="rep-status-{{ $element.ID }}">
              {{/* The status can be one of these: unverified, verification-failed, ok, problems, verifying, updating', initializing, verified */}}
              {{ if eq  $element.Status "unverified" }}

              {{ else if eq  $element.Status "ok" }}
              <span class="label label-success">ok</span>

              {{ else if eq  $element.Status "verified" }}
              <span class="label label-success">verified</span>


              {{ else if eq  $element.Status "verifying" }}
              <span class="label label-warning">verifing</span>


              {{ else if eq  $element.Status "updating" }}
              <span class="label label-warning">updating</span>

              {{ else if eq  $element.Status "initializing" }}
              <span class="label label-warning">initializing</span>


              {{ else if eq  $element.Status "verification-failed" }}
              <span class="label label-danger">verification failed</span>


              {{ else if eq  $element.Status "problems" }}
              <span class="label label-danger">problems</span>

              {{ end }}
            </div>
            </td>
            <td id="rep-action-{{ $element.ID }}">
              {{ if eq  $element.Status "unverified" }}
              <a href=""> <span data-toggle='tooltip' data-ID="{{ $element.ID }}" data-placement='top' title='' data-original-title='Update' class='repositories-init-harvest glyphicon glyphicon-retweet hover' ></span></a>

              {{ else if eq  $element.Status "ok" }}
              <a href=""> <span data-toggle='tooltip' data-ID="{{ $element.ID }}" data-placement='top' title='' data-original-title='Update' class='repositories-init-harvest glyphicon glyphicon-retweet hover' ></span></a>

              {{ else if eq  $element.Status "verified" }}
              <span class="text-muted">Working...</span>


              {{ else if eq  $element.Status "verifying" }}
              <span class="text-muted">Working...</span>


              {{ else if eq  $element.Status "updating" }}
              <span class="text-muted">Working...</span>

              {{ else if eq  $element.Status "initializing" }}
              <span class="text-muted">Working...</span>


              {{ else if eq  $element.Status "verification-failed" }}
    				<a href='/admin/log/'>See log</a>


              {{ else if eq  $element.Status "problems" }}
    				<a href='/admin/log/'>See log</a>

              {{ end }}
            </td>
            <td><a href="{{ $element.URL }}" target="_blank">{{ $element.URL |html }}</a></td>
            <td><a href="/admin/institutions/institution/{{ $institution.ID }}">{{ $institution.Name }}</a></td>
          </tr>
          {{end }}
        </tbody>
      </table>

      <div id="harvest-history-container" class="modal fade">
          <div class="modal-dialog">
              <div class="modal-content">
                  <div class="modal-header">
                      <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
                      <h4 class="modal-title">History</h4>
                  </div>
                  <div class="modal-body" id="harvest-history-element">
                  </div>
                  <div class="modal-footer">
                      <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                  </div>
              </div>
          </div>
      </div>

    </div>
    {{ else }}
    There are no repositories... :(
    {{ end }}
  </div>
</div>
<script>
var server = {};
server.host = {{ .host }};
</script>
<script src="/static/js/ws/websockets.js"></script>
<script src="/static/js/admin/repository/list.js"></script>
<script src="/static/js/admin/harvest/harvest.js"></script>
<script src="/static/js/admin/harvest/list.js"></script>
