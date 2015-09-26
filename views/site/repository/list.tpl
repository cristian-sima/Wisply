
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
      <div class="btn-group">
        <a href="/admin/repositories/add" class="btn btn-primary">Add repository</a>
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
    {{ if .anything }}
    <div class="table-responsive">
      <table class="table table-striped table-hover " id="repositories-list">
        <thead>
          <tr>
            <th>Name</th>
            <th>Current status</th>
            <th>Base URL</th>
            <th>Description</th>
            <th>Modify</th>
            <th>Delete</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $element := .repositories}}
          {{$safe := $element.Name|html}}
          <tr>
            <td>{{ $element.Name |html }}</td>
            <td> <div  id="rep-status-{{ $element.ID }}">
              {{/* The status can be one of these: unverified, verification-failed, ok, problems, verifying, updating', initializing, verified */}}
              {{ if eq  $element.Status "unverified" }}
              <span class="label label-info">Unverified</span><a href=""> <span data-toggle='tooltip' data-ID="{{ $element.ID }}" data-placement='top' title='' data-original-title='Start now!' class='repositories-init-harvest glyphicon glyphicon-sort-by-attributes hover' ></span></a>

              {{ else if eq  $element.Status "ok" }}
              <span class="label label-success">Ok</span>

              {{ else if eq  $element.Status "verified" }}
              <span class="label label-success">Verified</span>


              {{ else if eq  $element.Status "verifying" }}
              <span class="label label-warning">Verifing</span>


              {{ else if eq  $element.Status "updating" }}
              <span class="label label-warning">Updating</span>

              {{ else if eq  $element.Status "initializing" }}
              <span class="label label-warning">Initializing</span>
              <span data-toggle='tooltip' data-ID="{{ $element.ID }}" data-placement='top' title='' data-original-title='See process' class='repositories-init-harvest glyphicon glyphicon-th-list hover' ></span>


              {{ else if eq  $element.Status "verification-failed" }}
              <span class="label label-danger">Verification failed</span>
              <a href=""> <span data-toggle='tooltip' data-ID="{{ $element.ID }}" data-placement='top' title='' data-original-title='Try again' class='repositories-init-harvest glyphicon glyphicon glyphicon-refresh hover' ></span></a>


              {{ else if eq  $element.Status "problems" }}
              <span class="label label-danger">Problems</span>

              {{ end }}
            </div>
            </td>
            <td><a href="{{ $element.URL }}" target="_blank">{{ $element.URL |html }}</a></td>
            <td>{{ $element.Description |html }}</td>
            <td>
              <a href="/admin/repositories/modify/{{$element.ID}}">Modify</a>
            </td>
            <td>
              <a class="deleteRepositoryButton" data-id="{{$element.ID}}" data-name="{{$safe}}" href="/">Delete</a>
            </td>
          </tr>
          {{end }}
        </tbody>
      </table>
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
