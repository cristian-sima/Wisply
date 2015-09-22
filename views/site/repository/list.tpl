
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
      <h4>Repositories</h4>
    {{ if .anything }}
    <div class="table-responsive">
      <table class="table table-striped table-hover ">
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
            <td>
              {{ if eq  $element.Status "unverified" }}
              <span class="label label-warning">Unverified</span> <span data-toggle='tooltip' data-ID="{{ $element.ID }}" data-placement='top' title='' data-original-title='Validate now!' class='repositories-init-harvest glyphicon glyphicon-circle-arrow-down hover' ></span>

              {{ else if eq  $element.Status "ok" }}
              <span class="label label-success">Ok</span>


              {{ else if eq  $element.Status "verifying" }}
              <span class="label label-info">Verifing</span>


              {{ else if eq  $element.Status "updating" }}
              <span class="label label-warning">Updating</span>


              {{ else if eq  $element.Status "verification-failed" }}
              <span class="label label-danger">Verification failed</span>
              <a href='' data-toggle='tooltip' data-placement='top' title='' data-original-title='Try again'><span class='glyphicon glyphicon-refresh'  ></span></a>


              {{ else if eq  $element.Status "problems-harvesting" }}
              <span class="label label-danger">Problems harvesting</span>

              {{ end }}
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
<script src="/static/js/admin/repository/list.js"></script>
