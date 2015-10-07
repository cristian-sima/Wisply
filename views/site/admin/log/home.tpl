<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/repositories">Repositories</a></li>
      <li class="active">Event log</li>
    </ul>
  </div>
  <div class="panel-body">
    <div class="table-responsive">
      <span class="text-warning">
				<span class="glyphicon glyphicon-warning-sign"></span>
        This page is not live updated.
			</span>
      <table id="list-accounts" class="table table-hover table-bordered table-condensed">
        <thead>
          <tr>
            <th class="hidden-xs">#</th>
            <th>Type</th>
            <th>Repository</th>
            <th>State</th>
            <th>Start</th>
            <th>End</th>
            <th>Duration</th>
            <th>Current Operation</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $element := .processes}}
          <tr>
            <td class="col-md-1">{{ $element.ID }}</td>
            <td class="col-md-1">{{ $element.Action.Content }}</td>
            <td class="col-md-2"><a href="/admin/repositories/repository/{{ $element.Repository.ID }}">{{ $element.Repository.Name }}</a></td>
            <!-- start state -->
            <td class="col-md-1">
            {{ if $element.Action.IsRunning }}
            <span class="text-warning">Working</span>
            {{ else }}
            Finished
            {{ end }}
            </td>
            <!-- end state -->
            <td class="col-md-3">{{ $element.GetStartDate }}</td>
            <td class="col-md-3">{{ $element.GetEndDate }}</td>
            <td class="col-md-1">{{ $element.GetDuration }}</td>
            <td class="col-md-1">
              {{ if $element.GetCurrentOperation }}
              {{ $element.GetCurrentOperation.Content }}
              {{ else }}
              -
              {{ end }}
            </td>
          </tr>
          {{end }}
        </tbody>
      </table>
    </div>
  </div>
</div>
